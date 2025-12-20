package consumer

import (
	"context"
	"errors"
	"flip/internal/entity"
	"flip/internal/repository"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gofiber/fiber/v2/log"
)

type ReconciliationConsumer struct {
	Event               chan ReconciliationConsumerMessage
	Worker              int
	Buffer              int
	Processed           sync.Map
	MaxRetry            int
	StatementRepository repository.StatementRepositoryMethod
	BusyWorker          atomic.Int64
	RetryInflight       atomic.Int64
}

type ReconciliationConsumerMessage struct {
	ID      string
	Attempt int
	Data    entity.Statement
}

type ReconciliationConsumerMethod interface {
	Listen(ctx context.Context)
	Publish(ctx context.Context, message ReconciliationConsumerMessage)
	GetStatus(ctx context.Context) *entity.Health
}

func NewReconciliationConsumer(worker int, maxRetry int, statementRepository repository.StatementRepositoryMethod) ReconciliationConsumerMethod {
	return &ReconciliationConsumer{
		Worker:              worker,
		MaxRetry:            maxRetry,
		Event:               make(chan ReconciliationConsumerMessage),
		Processed:           sync.Map{},
		StatementRepository: statementRepository,
	}
}

func (i *ReconciliationConsumer) GetStatus(ctx context.Context) *entity.Health {
	health := new(entity.Health)

	health.Worker.Total = i.Worker
	health.Worker.Busy = int(i.BusyWorker.Load())
	health.Worker.Available = i.Worker - int(i.BusyWorker.Load())

	health.Retry.Inflight = int(i.RetryInflight.Load())

	var memstat runtime.MemStats
	runtime.ReadMemStats(&memstat)

	const mb = 1024 * 1024

	health.Usage.CurrentAllocation = memstat.Alloc / mb
	health.Usage.TotalAllocation = memstat.TotalAlloc / mb

	return health
}

func (i *ReconciliationConsumer) Listen(ctx context.Context) {
	for n := 0; n < i.Worker; n++ {
		go i.runWorker(ctx, n)
	}
}

func (i *ReconciliationConsumer) runWorker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			log.Info("cancel called")
			return
		case message := <-i.Event:
			log.Infof("incoming message %s consumed by worker %d", message.ID, id)
			i.BusyWorker.Add(1)

			if _, ok := i.Processed.Load(message.ID); ok {
				i.BusyWorker.Add(-1)
				continue
			}

			//do reconciliate here
			//lets make a bet of true and false, if true status will be updated to success, if false then generate error (to replicate retry mechanism)

			randomInt := rand.Intn(2)
			time.Sleep(time.Second * time.Duration(randomInt))

			randomBool := randomInt == 0
			var err error
			if randomBool {
				i.StatementRepository.UpdateToSuccess(ctx, message.Data.UploadID, message.Data.ID)
			} else {
				err = errors.New("test retry")
			}

			if err != nil {
				message.Attempt++
				if message.Attempt > i.MaxRetry {
					log.Warnf("statement id %s is failed permanently", message.ID)
					i.BusyWorker.Add(-1)
					continue
				}

				delay := time.Duration(1<<message.Attempt) * time.Second
				i.retryLater(ctx, message, delay)
				i.BusyWorker.Add(-1)
				continue
			}

			log.Infof("success reconcile message %s", message.ID)
			i.Processed.Store(message.ID, struct{}{})
			i.BusyWorker.Add(-1)
		}
	}
}

func (i *ReconciliationConsumer) retryLater(ctx context.Context, message ReconciliationConsumerMessage, delay time.Duration) {
	log.Infof("retry attempt %d for message id %s", message.Attempt, message.ID)
	i.RetryInflight.Add(1)
	go func() {
		timer := time.NewTimer(delay)
		defer timer.Stop()

		select {
		case <-ctx.Done():
			return

		case <-timer.C:
			i.Event <- message
			i.RetryInflight.Add(-1)
		}
	}()
}

func (i *ReconciliationConsumer) Publish(ctx context.Context, message ReconciliationConsumerMessage) {
	i.Event <- message
}

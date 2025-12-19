package consumer

import (
	"context"
	"errors"
	"flip/internal/entity"
	"flip/internal/repository"
	"math/rand"
	"sync"
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
}

type ReconciliationConsumerMessage struct {
	ID      string
	Attempt int
	Data    *entity.Statement
}

type ReconciliationConsumerMethod interface {
	Listen(ctx context.Context)
	Publish(ctx context.Context, message ReconciliationConsumerMessage)
}

func NewReconciliationConsumer(worker int, buffer int, maxRetry int, statementRepository repository.StatementRepositoryMethod) ReconciliationConsumerMethod {
	return &ReconciliationConsumer{
		Worker:              worker,
		MaxRetry:            maxRetry,
		Event:               make(chan ReconciliationConsumerMessage, buffer),
		Processed:           sync.Map{},
		StatementRepository: statementRepository,
	}
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
			log.Infof("incoming message %s", message.ID)
			if _, ok := i.Processed.Load(message.ID); ok {
				continue
			}

			//do reconciliate here
			//lets make a bet of true and false, if true status will be updated to success, if false then generate error (to replicate retry mechanism)

			randomBool := rand.Intn(2) == 0
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
					continue
				}

				delay := time.Duration(1<<message.Attempt) * time.Second
				i.retryLater(ctx, message, delay)

				continue
			}

			log.Infof("success reconcile message %s", message.ID)
			i.Processed.Store(message.ID, struct{}{})
		}
	}
}

func (i *ReconciliationConsumer) retryLater(ctx context.Context, message ReconciliationConsumerMessage, delay time.Duration) {
	log.Infof("retry attempt for message id %s on %d times", message.ID, message.Attempt)
	go func() {
		timer := time.NewTimer(delay)
		defer timer.Stop()

		select {
		case <-ctx.Done():
			return

		case <-timer.C:
			i.Event <- message
		}
	}()
}

func (i *ReconciliationConsumer) Publish(ctx context.Context, message ReconciliationConsumerMessage) {
	i.Event <- message
}

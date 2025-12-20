package controller

import (
	"context"
	"encoding/csv"

	"flip/internal/consumer"
	"flip/internal/entity"
	"flip/internal/helper/response"

	"io"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

func (i *Controller) ActionStatementPost(c *fiber.Ctx) error {
	uploadID := uuid.NewString()
	csvFile, err := c.FormFile("file")
	if err != nil {
		return response.CreateFail(c, err.Error())
	}

	file, err := csvFile.Open()
	if err != nil {
		return response.CreateError(c, err.Error())
	}
	defer file.Close()

	go func() error {
		reader := csv.NewReader(file)

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			}

			if err != nil {
				return err
			}

			timestamp, err := strconv.Atoi(record[0])
			if err != nil {
				timestamp = 0
			}

			amount, err := strconv.Atoi(record[3])
			if err != nil {
				amount = 0
			}

			txID := uuid.NewString()
			data := entity.Statement{
				UploadID:     uploadID,
				ID:           txID,
				Timestamp:    timestamp,
				Counterparty: record[1],
				Type:         record[2],
				Amount:       amount,
				Status:       record[4],
				Description:  record[5],
			}

			err = i.Usecase.CreateStatement(context.Background(), uploadID, data)
			if err != nil {
				log.Errorf("failed to add statement for %s", uploadID)
			}

			if data.IsFailed() {
				i.Usecase.Publish(context.Background(), consumer.ReconciliationConsumerMessage{
					ID:      txID,
					Attempt: 0,
					Data:    data,
				})
			}
		}
		return nil
	}()

	return response.CreateSuccess(c, map[string]interface{}{
		"upload_id": uploadID,
	})
}

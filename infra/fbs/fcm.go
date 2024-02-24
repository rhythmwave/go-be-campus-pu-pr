package fbs

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sccicitb/pupr-backend/config"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/data/repositories/notification"
	"github.com/sccicitb/pupr-backend/infra/db"
	"github.com/sccicitb/pupr-backend/utils"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"google.golang.org/grpc/codes"

	models "github.com/sccicitb/pupr-backend/data/models/notification"
	objects "github.com/sccicitb/pupr-backend/objects/notification"
)

type fcm struct {
	*db.DB
	*config.Firebase
}

// NewFCM connects fcm to FCMInterface
// Params:
// db: database
// cfg: FCM config
// Returns FCMInterface
func NewFCM(db *db.DB, cfg *config.Firebase) utils.NotificationInterface {
	return &fcm{
		db,
		cfg,
	}
}

func send(client *messaging.Client, data objects.SendNotification) {

	b, _ := json.Marshal(data.Data.PayloadData)
	messageData := map[string]string{
		"status":             data.Data.Status,
		"body":               data.Body,
		"title":              data.Title,
		"android_channel_id": data.Data.AndroidChannelID,
		"priority":           data.Data.Priority,
		"click_action":       data.Data.ClickAction,
		"type":               data.Data.Type,
		"id":                 data.Data.ID,
		"ref_url":            data.Data.RefURL,
		"payload_data":       string(b),
	}

	targetTokens := []string{}
	for _, v := range data.TargetTokens {
		if v != "" {
			targetTokens = append(targetTokens, v)
		}
	}

	for i := 0; i < len(targetTokens); i += constants.DefaultSingleSendNotification {
		lastIndex := i + constants.DefaultSingleSendNotification
		if lastIndex > len(targetTokens) {
			lastIndex = len(targetTokens)
		}

		message := messaging.MulticastMessage{
			Tokens: targetTokens[i:lastIndex],
			Notification: &messaging.Notification{
				Title: data.Title,
				Body:  data.Body,
			},
			Android: &messaging.AndroidConfig{
				Notification: &messaging.AndroidNotification{
					Title: data.Title,
					Body:  data.Body,
				},
				Data: messageData,
			},
			Data: messageData,
		}

		go func() {
			if len(message.Tokens) > 0 {
				br, err := client.SendMulticast(context.Background(), &message)
				if err != nil {
					log.Errorln(err)
				}

				if br.FailureCount > 0 {
					log.Error(fmt.Sprintf("failed sending %d message", br.FailureCount))
					for _, response := range br.Responses {
						log.Errorln(response)
					}
				}
			}
		}()
	}
}

func (f fcm) saveToTable(ctx context.Context, tx *sqlx.Tx, data objects.SendNotification) *constants.ErrorResponse {
	notifRepo := notification.NewNotificationRepository(f.DB)

	notifData := []models.Create{}
	for _, v := range data.UserIDs {
		notifData = append(notifData, models.Create{
			UserID:    v,
			Title:     data.Title,
			Body:      data.Body,
			LinkUrl:   utils.NewNullString(data.Data.ClickAction),
			ExpiredAt: utils.NewNullTime(data.ExpiryTime),
			IdType:    utils.NewNullString(data.Data.ID),
			Type:      data.Data.Type,
			Data:      utils.StructToJson(data),
		})
	}
	if len(notifData) > 0 {
		errs := notifRepo.BulkCreate(ctx, tx, notifData)
		if errs != nil {
			return errs
		}
	}

	return nil
}

func (f fcm) GetProviderName() string {
	return constants.Fcm
}

// SendNotification function to send mobile notification via FCM
// Params:
// data: send notification data
func (f fcm) SendNotification(data objects.SendNotification) *constants.ErrorResponse {
	ctx := context.Background()

	tx, err := f.DB.Begin(ctx)
	if err != nil {
		return constants.ErrUnknown
	}

	if data.SaveToTable {
		errs := f.saveToTable(ctx, tx, data)
		if errs != nil {
			_ = tx.Rollback()
			return errs
		}
	}
	err = tx.Commit()

	if err != nil {
		_ = tx.Rollback()
		return constants.ErrUnknown
	}

	opt := option.WithCredentialsFile(f.KeyFileDir)
	config := &firebase.Config{
		ProjectID: f.ProjectID,
	}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		// _ = tx.Rollback()
		return constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		// _ = tx.Rollback()
		return constants.Error(http.StatusInternalServerError, codes.Internal, constants.DefaultCustomErrorCode, err.Error())
	}

	send(client, data)

	return nil
}

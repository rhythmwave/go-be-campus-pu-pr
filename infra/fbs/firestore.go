package fbs

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/sccicitb/pupr-backend/config"
	"google.golang.org/api/option"
)

// FireStore FireStore
type FireStore struct {
	Context context.Context
	Client  *firestore.Client
}

// NewFirestore init new Firestore instance
func NewFirestore(ctx context.Context, cfg *config.Firebase) (*FireStore, error) {
	opt := option.WithCredentialsFile(cfg.KeyFileDir)

	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	// close client
	defer client.Close().Error()

	return &FireStore{ctx, client}, nil
}

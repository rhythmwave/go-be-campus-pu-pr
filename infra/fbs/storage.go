package fbs

import (
	"context"
	"fmt"
	"io"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go"
	fsStorage "firebase.google.com/go/storage"
	"github.com/sccicitb/pupr-backend/config"
	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

// FirebaseStorage implementation to Storage interface
type FirebaseStorage struct {
	Context             context.Context
	Client              *fsStorage.Client
	Bucket              *storage.BucketHandle
	PublicAccess        bool
	PublicAccessBaseURL string
}

// NewFirebaseStorage new instance for firebase object
func NewFirebaseStorage(ctx context.Context, cfg *config.Firebase) (utils.StorageInterface, error) {
	var fs FirebaseStorage

	configStorage := &firebase.Config{
		StorageBucket: cfg.Bucket,
	}
	opt := option.WithCredentialsFile(cfg.KeyFileDir)
	app, err := firebase.NewApp(ctx, configStorage, opt)
	if err != nil {
		log.Error(err)
		return &fs, err
	}

	client, err := app.Storage(ctx)
	if err != nil {
		log.Error(err)
		return &fs, err
	}
	// defer client.Close()

	bucket, err := client.DefaultBucket()
	if err != nil {
		log.Error(err)
		return &fs, err
	}

	return &FirebaseStorage{
		Context:             ctx,
		Bucket:              bucket,
		Client:              client,
		PublicAccess:        cfg.StoragePublicAccess,
		PublicAccessBaseURL: fmt.Sprintf("%s/%s", "https://storage.googleapis.com", strings.ReplaceAll(cfg.Bucket, "gs://", "")),
	}, nil
}

// PutFile to firebase storage
func (fs *FirebaseStorage) PutFile(data io.Reader, path, mimeType string) error {

	// upload it to google cloud store
	wc := fs.Bucket.Object(path).NewWriter(fs.Context)
	if _, err := io.Copy(wc, data); err != nil {
		log.Error(err)
		return err
	}
	err := wc.Close()
	if err != nil {
		log.Error(err)
		return err
	}
	if fs.PublicAccess {
		if err := fs.Bucket.Object(path).ACL().Set(fs.Context, storage.AllUsers, storage.RoleReader); err != nil {
			return err
		}
	}

	return nil
}

// DeleteFile from firebase storage
func (fs *FirebaseStorage) DeleteFile(path string) error {
	return fs.Bucket.Object(path).Delete(fs.Context)
}

// GetBytes from firebase storage
func (fs *FirebaseStorage) GetBytes(path string) ([]byte, error) {
	rc, err := fs.Bucket.Object(path).NewReader(fs.Context)
	if err != nil {
		return []byte{}, err
	}
	defer rc.Close()
	slurp, err := io.ReadAll(rc)
	if err != nil {
		return []byte{}, err
	}
	return slurp, nil
}

// GetURL get public url of firebase storage object
func (fs *FirebaseStorage) GetURL(path string, expTime *time.Time) (string, error) {
	if fs.PublicAccess {
		// if err := fs.Bucket.Object(path).ACL().Set(fs.Context, storage.AllUsers, storage.RoleReader); err != nil {
		// 	return "", err
		// }
		return fmt.Sprintf("%s/%s", fs.PublicAccessBaseURL, path), nil
	}
	exp := time.Now().Add(24 * time.Hour)
	if expTime != nil {
		exp = *expTime
	}
	opts := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "GET",
		Expires: exp,
	}

	return fs.Bucket.SignedURL(path, opts)
}

// GetProvider GetProvider
func (fs *FirebaseStorage) GetProvider() string {
	return constants.FirebaseStorageProvider
}

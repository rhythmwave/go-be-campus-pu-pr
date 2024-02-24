package file

import (
	"context"

	"github.com/sccicitb/pupr-backend/constants"
	"github.com/sccicitb/pupr-backend/infra/context/infra"
	fileObject "github.com/sccicitb/pupr-backend/objects/file"
)

// FileServiceInterface FileServiceInterface
type FileServiceInterface interface {
	UploadBlobTemp(ctx context.Context, data []byte, filePath string) (fileObject.FileSavedResponse, *constants.ErrorResponse)
	UploadBase64Temp(ctx context.Context, data string, filePath string) (fileObject.FileSavedResponse, *constants.ErrorResponse)
	GetURL(ctx context.Context, path string) (string, *constants.ErrorResponse)
}

// NewFileService NewFileService
func NewFileService(infraCtx *infra.InfraCtx) FileServiceInterface {
	return &fileService{
		infraCtx,
	}
}

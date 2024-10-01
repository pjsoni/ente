package filedata

import (
	"fmt"

	"github.com/ente-io/museum/ente"
)

type Entity struct {
	FileID           int64           `json:"fileID"`
	Type             ente.ObjectType `json:"type"`
	EncryptedData    string          `json:"encryptedData"`
	DecryptionHeader string          `json:"decryptionHeader"`
}

// GetFilesData should only be used for getting the preview video playlist and derived metadata.
type GetFilesData struct {
	FileIDs []int64         `json:"fileIDs" binding:"required"`
	Type    ente.ObjectType `json:"type" binding:"required"`
}

func (g *GetFilesData) Validate() error {
	if g.Type != ente.PreviewVideo && g.Type != ente.MlData {
		return ente.NewBadRequestWithMessage(fmt.Sprintf("unsupported object type %s", g.Type))
	}
	if len(g.FileIDs) == 0 {
		return ente.NewBadRequestWithMessage("fileIDs are required")
	}
	if len(g.FileIDs) > 200 {
		return ente.NewBadRequestWithMessage("fileIDs should be less than or equal to 200")
	}
	return nil
}

type GetFileData struct {
	FileID int64           `form:"fileID" binding:"required"`
	Type   ente.ObjectType `form:"type" binding:"required"`
}

func (g *GetFileData) Validate() error {
	if g.Type != ente.PreviewVideo && g.Type != ente.MlData {
		return ente.NewBadRequestWithMessage(fmt.Sprintf("unsupported object type %s", g.Type))
	}
	return nil
}

type GetFilesDataResponse struct {
	Data                []Entity `json:"data"`
	PendingIndexFileIDs []int64  `json:"pendingIndexFileIDs"`
	ErrFileIDs          []int64  `json:"errFileIDs"`
}

// S3FileMetadata stuck represents the metadata that is stored in the S3 bucket for non-file type metadata
// that is stored in the S3 bucket.
type S3FileMetadata struct {
	Version          int    `json:"v"`
	EncryptedData    string `json:"encryptedData"`
	DecryptionHeader string `json:"header"`
	Client           string `json:"client"`
}

type GetPreviewURLRequest struct {
	FileID int64           `form:"fileID" binding:"required"`
	Type   ente.ObjectType `form:"type" binding:"required"`
	Suffix *string         `form:"suffix"`
}

func (g *GetPreviewURLRequest) Validate() error {
	if g.Type != ente.PreviewVideo && g.Type != ente.PreviewImage {
		return ente.NewBadRequestWithMessage(fmt.Sprintf("unsupported object type %s", g.Type))
	}
	return nil
}

type PreviewUploadUrlRequest struct {
	FileID int64           `form:"fileID" binding:"required"`
	Type   ente.ObjectType `form:"type" binding:"required"`
}

type PreviewUploadUrl struct {
	ObjectID string `json:"objectID" binding:"required"`
	Url      string `json:"url" binding:"required"`
}

func (g *PreviewUploadUrlRequest) Validate() error {
	if g.Type != ente.PreviewVideo && g.Type != ente.PreviewImage {
		return ente.NewBadRequestWithMessage(fmt.Sprintf("unsupported object type %s", g.Type))
	}
	return nil
}

// Row represents the data that is stored in the file_data table.
type Row struct {
	FileID int64
	UserID int64
	Type   ente.ObjectType
	// If a file type has multiple objects, then the size is the sum of all the objects.
	Size         int64
	LatestBucket string
	ObjectID     *string
	// For HLS video object, there's no object nonce, all relevant data
	// is stored in the metadata object that primarily contains the playlist.
	ObjectNonce       *string
	ReplicatedBuckets []string
	DeleteFromBuckets []string
	InflightReplicas  []string
	PendingSync       bool
	IsDeleted         bool
	SyncLockedTill    int64
	CreatedAt         int64
	UpdatedAt         int64
}

// S3FileMetadataObjectKey returns the object key for the metadata stored in the S3 bucket.
func (r *Row) S3FileMetadataObjectKey() string {
	if r.Type == ente.MlData || r.Type == ente.PreviewVideo {
		return ObjectMedata(r.FileID, r.UserID, r.Type, r.ObjectID)
	}
	panic(fmt.Sprintf("S3FileMetadata should not be written for %s type", r.Type))
}

// GetS3FileObjectKey returns the object key for the file data stored in the S3 bucket.
func (r *Row) GetS3FileObjectKey() string {
	if r.Type == ente.PreviewVideo || r.Type == ente.PreviewImage {
		return ObjectKey(r.FileID, r.UserID, r.Type, r.ObjectID)
	}
	panic(fmt.Sprintf("unsupported object type %s", r.Type))
}

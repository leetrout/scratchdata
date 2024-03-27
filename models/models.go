package models

import (
	"github.com/scratchdata/scratchdata/pkg/storage/blobstore"
	"github.com/scratchdata/scratchdata/pkg/storage/cache"
	"github.com/scratchdata/scratchdata/pkg/storage/database"
	"github.com/scratchdata/scratchdata/pkg/storage/queue"
)

type StorageServices struct {
	Database  database.Database
	Cache     cache.Cache
	Queue     queue.Queue
	BlobStore blobstore.BlobStore
}

type Column struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	JSONType string `json:"-"`
}

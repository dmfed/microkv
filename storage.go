package microkv

import (
	"github.com/dmfed/microkv/internal/filesystem"
	"github.com/dmfed/microkv/internal/natsobj"
)

var (
	_ Storage = (natsobj.Storage)(nil)
	_ Storage = (filesystem.Storage)(nil)
)

// Storage is a very basic object store.
type Storage interface {
	// Save puts file with name 'key' into the store. If a file with such name
	// already exists it it gets overwritten.
	Save(key string, data []byte) error
	// Load returns contents of file named 'key'.
	Load(key string) ([]byte, error)
	// Delete removes file named 'key' from the store.
	// IF such file doe not exist delete returns nil.
	Delete(key string) error
	// Close must be called when you're done working with Storage.
	Close() error
}

// NewNats connects to NATS messaging system and tries to create
// a new object storage with name 'bucket'. The returned Storage
// uses the created bucket as underlying physical store.
func NewNats(bucket string, url string) (Storage, error) {
	return natsobj.Open(bucket, url)
}

// NewFS established a key/value within the directory 'path'
// and uses is as underlying physical store.
func NewFS(path string) (Storage, error) {
	return filesystem.Open(path)
}

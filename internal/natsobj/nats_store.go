package natsobj

import (
	"github.com/nats-io/nats.go"
)

// Storage
type Storage interface {
	Save(key string, data []byte) error
	Load(key string) ([]byte, error)
	Delete(key string) error
	Close() error
}

func Open(bucket string, url string) (Storage, error) {
	return newNatsStorage(bucket, url)
}

type natsStorage struct {
	store nats.ObjectStore
	conn  *nats.Conn
}

func newNatsStorage(bucket, url string) (Storage, error) {
	nc, err := nats.Connect(url)
	if err != nil {
		return nil, err
	}

	js, err := nc.JetStream()
	if err != nil {
		return nil, err
	}

	cfg := &nats.ObjectStoreConfig{
		Bucket:      bucket,
		Description: "microkv bucket",
		MaxBytes:    -1,
		Storage:     nats.FileStorage,
		Compression: true,
	}

	store, err := js.CreateObjectStore(cfg)
	if err != nil {
		return nil, err
	}

	st := &natsStorage{store: store, conn: nc}

	return st, nil
}

func (n *natsStorage) Save(key string, data []byte) error {
	if _, err := n.store.PutBytes(key, data); err != nil {
		return err
	}
	return nil
}

func (n *natsStorage) Load(key string) ([]byte, error) {
	b, err := n.store.GetBytes(key)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (n *natsStorage) Delete(key string) error {
	if err := n.store.Delete(key); err != nil {
		return err
	}

	return nil
}

func (n *natsStorage) Close() error {
	n.conn.Close()

	return nil
}

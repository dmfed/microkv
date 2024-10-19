package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

// Storage
type Storage interface {
	Save(key string, data []byte) error
	Load(key string) ([]byte, error)
	Delete(key string) error
	Close() error
}

func Open(path string) (Storage, error) {
	info, err := os.Stat(path)
	if err != nil {
		return nil, err
	}

	if !info.IsDir() {
		return nil, fmt.Errorf("provided path %s is not a directory", path)
	}

	abs, err := filepath.Abs(path)
	if err != nil {
		return nil, fmt.Errorf("could not tarnslate %s to absolute path", path)
	}

	return &fsStorage{prefix: abs}, nil
}

type fsStorage struct {
	prefix string
}

func (f *fsStorage) Save(key string, data []byte) error {
	path := f.toAbs(key)

	if err := os.WriteFile(path, data, 0664); err != nil {
		return err
	}

	return nil
}
func (f *fsStorage) Load(key string) ([]byte, error) {
	path := f.toAbs(key)

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return b, nil
}
func (f *fsStorage) Delete(key string) error {
	path := f.toAbs(key)

	err := os.Remove(path)
	if err != nil && os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	return nil
}
func (f *fsStorage) Close() error {
	return nil
}

func (f *fsStorage) toAbs(path string) string {
	return filepath.Join(f.prefix, path)
}

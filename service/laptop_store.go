package service

import (
	"errors"
	"fmt"
	"sync"

	message "github.com/aniketh3014/simple-protobuf/pb"
	"github.com/jinzhu/copier"
)

type LaptopStore interface {
	Save(laptop *message.Laptop) error
	Find(id string) (*message.Laptop, error)
}

type InmemoryLaptopStore struct {
	mutex sync.RWMutex
	data  map[string]*message.Laptop
}

func NewInmemoryLaptopStore() *InmemoryLaptopStore {
	return &InmemoryLaptopStore{
		data: make(map[string]*message.Laptop),
	}
}

var ErrAlreadyExists = errors.New("record already exists")

func (store *InmemoryLaptopStore) Find(id string) (*message.Laptop, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	laptop := store.data[id]
	if laptop == nil {
		return nil, fmt.Errorf("cannot find laptop with ID %s", id)
	}

	copiedLaptop, err := deepCopy(laptop)
	if err != nil {
		return nil, fmt.Errorf("failed to copy laptop: %v", err)
	}

	return copiedLaptop, nil
}

func (store *InmemoryLaptopStore) Save(laptop *message.Laptop) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[laptop.Id] != nil {
		return ErrAlreadyExists
	}

	other, err := deepCopy(laptop)
	if err != nil {
		return err
	}

	store.data[other.Id] = other
	return nil
}

func deepCopy(laptop *message.Laptop) (*message.Laptop, error) {
	other := &message.Laptop{}

	err := copier.Copy(other, laptop)
	if err != nil {
		return nil, fmt.Errorf("cannot copy laptop data: %w", err)
	}
	return other, nil
}

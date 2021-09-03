package storage

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
)

type StorageContextConfig struct {
	AccountId  string
	Region     string
	Date       time.Time
	Cloud      string
	Service    string
	DataSource string
}

type StorageBackendContext interface {
	Store(ctx context.Context, obj interface{}) error
	Close(ctx context.Context) error
	GetConfig() StorageContextConfig
}

func LogContextError(context StorageBackendContext, message string) {
	config := context.GetConfig()
	logrus.WithFields(logrus.Fields{
		"account_id": config.AccountId,
		"region":     config.Region,
		"date":       config.Date,
		"cloud":      config.Cloud,
		"service":    config.Service,
		"datasource": config.DataSource,
	}).Error(message)
}

type StorageBackend interface {
	GetStorageContext(config StorageContextConfig, sampleObject interface{}) (StorageBackendContext, error)
}

type StorageBackendContextSet struct {
	Contexts []StorageBackendContext
}

func (set *StorageBackendContextSet) Store(ctx context.Context, obj interface{}) map[StorageBackendContext]error {
	errors := map[StorageBackendContext]error{}
	for _, storageContext := range set.Contexts {
		err := storageContext.Store(ctx, obj)
		if err != nil {
			//QUESTIONN is this storing a pointer to storageContext in the map, and hence will always point to the last context at the end?
			errors[storageContext] = err
		}
	}

	if len(errors) == 0 {
		return nil
	} else {
		return errors
	}
}

func (set *StorageBackendContextSet) Close(ctx context.Context) map[StorageBackendContext]error {
	errors := map[StorageBackendContext]error{}
	for _, storageContext := range set.Contexts {
		err := storageContext.Close(ctx)
		if err != nil {
			//QUESTIONN is this storing a pointer to storageContext in the map, and hence will always point to the last context at the end?
			errors[storageContext] = err
		}
	}

	if len(errors) == 0 {
		return nil
	} else {
		return errors
	}
}

type StorageManager struct {
	StorageBackends []StorageBackend
}

func (sm *StorageManager) GetStorageContextSet(config StorageContextConfig, sampleObject interface{}) (*StorageBackendContextSet, error) {
	contexts := make([]StorageBackendContext, len(sm.StorageBackends))
	for i, backend := range sm.StorageBackends {
		storageContext, err := backend.GetStorageContext(config, sampleObject)
		if err != nil {
			return nil, err
		}
		contexts[i] = storageContext
	}

	return &StorageBackendContextSet{
		Contexts: contexts,
	}, nil
}

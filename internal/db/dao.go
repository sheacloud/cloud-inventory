package db

import (
	"context"

	"github.com/sheacloud/cloud-inventory/pkg/meta"
)

type DAO interface {
	AWS() AWSDAO
	WriteInventoryResults(ctx context.Context, metadata *meta.InventoryResults) error
	WriteIngestionTimestamp(ctx context.Context, timestamp *meta.IngestionTimestamp) error
}

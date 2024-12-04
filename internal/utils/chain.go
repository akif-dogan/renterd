package utils

import (
	"time"

	"go.thebigfile.com/core/types"
)

func IsSynced(b types.Block) bool {
	return time.Since(b.Timestamp) <= 3*time.Hour
}

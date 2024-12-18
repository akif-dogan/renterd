package stores

import (
	"context"

	"go.thebigfile.com/core/types"
	"go.thebigfile.com/renterd/stores/sql"
)

// ChainIndex returns the last stored chain index.
func (s *SQLStore) ChainIndex(ctx context.Context) (ci types.ChainIndex, err error) {
	err = s.db.Transaction(ctx, func(tx sql.DatabaseTx) error {
		ci, err = tx.Tip(ctx)
		return err
	})
	return
}

// ProcessChainUpdate returns a callback function that process a chain update
// inside a transaction.
func (s *SQLStore) ProcessChainUpdate(ctx context.Context, applyFn func(sql.ChainUpdateTx) error) error {
	return s.db.Transaction(ctx, func(tx sql.DatabaseTx) error {
		return tx.ProcessChainUpdate(ctx, applyFn)
	})
}

// ResetChainState deletes all chain data in the database.
func (s *SQLStore) ResetChainState(ctx context.Context) error {
	return s.db.Transaction(ctx, func(tx sql.DatabaseTx) error {
		return tx.ResetChainState(ctx)
	})
}

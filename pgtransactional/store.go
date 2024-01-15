package pgtransactional

import (
	"context"

	"github.com/thefabric-io/pgspecification"
	"github.com/thefabric-io/transactional"
)

type Store[A any, R any] interface {
	Save(ctx context.Context, transaction transactional.Transaction, aa ...*A) error
	Load(ctx context.Context, transaction transactional.Transaction, specs ...pgspecification.Specification) (*R, error)
	Exists(ctx context.Context, transaction transactional.Transaction, specs ...pgspecification.Specification) (bool, error)
	List(ctx context.Context, transaction transactional.Transaction, specs ...pgspecification.Specification) ([]*R, error)
}

type SearchableStore[A any, R any] interface {
	Store[A, R]
	Search(ctx context.Context, transaction transactional.Transaction, query []float64, specifications ...pgspecification.Specification) ([]*R, error)
}

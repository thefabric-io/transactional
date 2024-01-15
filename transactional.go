package transactional

import (
	"context"
)

type Transaction interface {
	Commit() error
	Rollback() error
}

type Transactional interface {
	BeginTransaction(ctx context.Context, opts BeginTransactionOptions) (Transaction, error)
	DefaultLogFields() map[string]any
}

type TxAccessMode string

func (t TxAccessMode) String() string {
	return string(t)
}

const (
	ReadWrite TxAccessMode = "read write"
	ReadOnly  TxAccessMode = "read only"
)

type TxIsoLevel string

func (t TxIsoLevel) String() string {
	return string(t)
}

const (
	Serializable    TxIsoLevel = "serializable"
	RepeatableRead  TxIsoLevel = "repeatable read"
	ReadCommitted   TxIsoLevel = "read committed"
	ReadUncommitted TxIsoLevel = "read uncommitted"
)

type TxDeferrableMode string

func (t TxDeferrableMode) String() string {
	return string(t)
}

const (
	Deferrable    TxDeferrableMode = "deferrable"
	NotDeferrable TxDeferrableMode = "not deferrable"
)

func DefaultWriteTransactionOptions() BeginTransactionOptions {
	return BeginTransactionOptions{
		AccessMode:     ReadWrite,
		IsolationLevel: Serializable,
		DeferrableMode: NotDeferrable,
	}
}

func DefaultReadOnlyTransactionOptions() BeginTransactionOptions {
	return BeginTransactionOptions{
		AccessMode:     ReadOnly,
		IsolationLevel: Serializable,
		DeferrableMode: NotDeferrable,
	}
}

type BeginTransactionOptions struct {
	AccessMode     TxAccessMode
	IsolationLevel TxIsoLevel
	DeferrableMode TxDeferrableMode
}

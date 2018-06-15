package scope

import (
	"context"
	"io"
)

type CloserContext interface {
	io.Closer
	context.Context
}

type closerContext struct {
	context.Context
	parent io.Closer
	cancel context.CancelFunc
}

func NewCloserContext(parent io.Closer) CloserContext {
	ctx, cancel := context.WithCancel(context.Background())

	return &closerContext{
		Context: ctx,
		parent:  parent,
		cancel:  cancel,
	}
}

func (c *closerContext) Close() error {
	c.cancel()
	return c.parent.Close()
}

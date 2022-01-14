package contexts

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/mdalbrid/utils/logger"
)

const (
	CtxCanceledErr = "context canceled"
)

type InternalModuleMainContext interface {
	context.Context
	Add(delta int)
	Wait()
	WgDone()

	GetCtx() context.Context
	Cancel()

	GetName() string
	Logger() logger.Logger
}

type InternalModuleMainContextStruct struct {
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
	name   string
	logger logger.Logger
}

func NewInternalModuleMainContext(name string) InternalModuleMainContext {
	ctx, cancel := context.WithCancel(context.Background())
	mc := &InternalModuleMainContextStruct{
		ctx:    ctx,
		cancel: cancel,
		name:   name,
		logger: logger.NewLogger(fmt.Sprintf(`[%s]`, name)),
	}
	go func() {
		sigCh := make(chan os.Signal, 2)
		signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
		sig := <-sigCh
		mc.logger.Warn(" interrupt signal ", sig)
		mc.Cancel()
	}()
	return mc
}

func (c *InternalModuleMainContextStruct) Add(delta int) {
	c.wg.Add(delta)
}

func (c *InternalModuleMainContextStruct) Wait() {
	c.wg.Wait()
}

func (c *InternalModuleMainContextStruct) WgDone() {
	c.wg.Done()
}

func (c *InternalModuleMainContextStruct) GetCtx() context.Context {
	return c.ctx
}
func (c *InternalModuleMainContextStruct) Cancel() {
	c.Logger().Info(" Cancel main context")
	c.cancel()
}

func (c *InternalModuleMainContextStruct) GetName() string {
	return c.name
}
func (c *InternalModuleMainContextStruct) Logger() logger.Logger {
	return c.logger
}

// Deadline proxy context.Context.Deadline
func (c *InternalModuleMainContextStruct) Deadline() (deadline time.Time, ok bool) {
	return c.ctx.Deadline()
}

// Done proxy context.Context.Done
func (c *InternalModuleMainContextStruct) Done() <-chan struct{} {
	return c.ctx.Done()
}

// Err proxy context.Context.Err
func (c *InternalModuleMainContextStruct) Err() error {
	return c.ctx.Err()
}

// Value proxy context.Context.Value
func (c *InternalModuleMainContextStruct) Value(key interface{}) interface{} {
	return c.ctx.Value(key)
}

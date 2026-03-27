package server

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	gracefulShutdown *GracefulShutdown
	once             sync.Once
)

type GracefulShutdown struct {
	shutdownFuncs     []func()
	waitGroup         *sync.WaitGroup
	isShutdownProcess bool
	lock              sync.RWMutex
}

func GetGracefulShutdown() *GracefulShutdown {
	once.Do(func() {
		gracefulShutdown = &GracefulShutdown{
			shutdownFuncs:     []func(){},
			isShutdownProcess: false,
			lock:              sync.RWMutex{},
			waitGroup:         new(sync.WaitGroup),
		}
	})
	return gracefulShutdown
}
func (g *GracefulShutdown) SetShutdownFunc(shutdownFuncs []func()) *GracefulShutdown {
	g.shutdownFuncs = shutdownFuncs
	return g
}
func (g *GracefulShutdown) AddShutdownFunc(shutdownFunc ...func()) *GracefulShutdown {
	g.shutdownFuncs = append(g.shutdownFuncs, shutdownFunc...)
	return g
}
func (g *GracefulShutdown) Listen() (chan os.Signal, func()) {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	return done, func() {
		close(done)
	}
}
func (g *GracefulShutdown) Shutdown() {
	defer os.Exit(0)
	g.enableShutdownProcess()
	for _, fn := range g.shutdownFuncs {
		g.waitGroup.Add(1)

		go func(fn func()) {
			defer g.waitGroup.Done()
			fn()
		}(fn)
	}

	g.waitGroup.Wait()
}

func (g *GracefulShutdown) getShutdownProcess() bool {
	g.lock.RLock()
	defer g.lock.RUnlock()
	return g.isShutdownProcess
}
func (g *GracefulShutdown) enableShutdownProcess() {
	g.lock.Lock()
	g.isShutdownProcess = true
	g.lock.Unlock()
}

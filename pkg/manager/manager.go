package manager

import (
	"context"
	"errors"
	"fmt"
	//"github.com/cossim/coss-server/pkg/cluster"
	"github.com/cossim/coss-server/pkg/config"
	"github.com/cossim/coss-server/pkg/discovery"
	"github.com/cossim/coss-server/pkg/healthz"
	plog "github.com/cossim/coss-server/pkg/log"
	"github.com/cossim/coss-server/pkg/server"
	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"go.uber.org/zap/zapcore"
	"net"
	"net/http"
	"time"
)

// Manager initializes shared dependencies such as Caches and Clients, and provides them to Runnables.
type Manager interface {
	// Cluster holds a variety of methods to interact with a cluster.
	//cluster.Cluster

	// Add will set requested dependencies on the component, and cause the component to be
	// started when Start is called.  Add will inject any dependencies for which the argument
	// implements the inject interface - e.g. inject.Client.
	Add(Runnable) error

	// AddMetricsExtraHandler adds an extra handler served on path to the http httpServer that serves metrics.
	// Might be useful to register some diagnostic endpoints e.g. pprof. Note that these endpoints meant to be
	// sensitive and shouldn't be exposed publicly.
	// If the simple path -> handler mapping offered here is not enough, a new http httpServer/listener should be added as
	// Runnable to the manager via Add method.
	AddMetricsExtraHandler(path string, handler http.Handler) error

	// AddHealthzCheck allows you to add Healthz checker
	AddHealthzCheck(name string, check healthz.Checker) error

	// AddReadyzCheck allows you to add Readyz checker
	AddReadyzCheck(name string, check healthz.Checker) error

	// Start starts all registered Controllers and blocks until the context is cancelled.
	Start(ctx context.Context) error

	// GetLogger returns this manager's logger.
	GetLogger() logr.Logger
}

type GrpcServer struct {
	MetricsBindAddress  string
	HealthzCheckAddress string
	server.GRPCService
}

type HttpServer struct {
	MetricsBindAddress string
	HealthCheckAddress string
	server.HTTPService
}

type Registry struct {
	// 启用服务注册
	Discover bool

	// 启用服务发现
	Register bool

	// 发现的服务名称
	Servers []string
}

type Config struct {
	// 通过配置中心获取配置
	LoadFromConfigCenter bool

	Registry Registry

	// 热更新配置
	Hot bool

	// 加载自身服务配置文件的路径
	Key string

	// 监听的其他文件 如common/mysql、common/redis
	Keys []string

	// 从本地路径加载配置
	LocalDir string

	// 配置中心地址
	RemoteConfigAddr string

	// 访问token
	RemoteConfigToken string
}

type Options struct {
	Grpc   GrpcServer
	Http   HttpServer
	Logger logr.Logger
	Config Config

	// Readiness probe endpoint name, defaults to "ready"
	ReadinessEndpointName string

	// Liveness probe endpoint name, defaults to "health"
	LivenessEndpointName string

	// MetricsBindAddress is the TCP address that the controller should bind to
	// for serving prometheus metrics.
	// It can be set to "0" to disable the metrics serving.
	MetricsBindAddress string

	// GracefulShutdownTimeout is the duration given to runnable to stop before the manager actually returns on stop.
	// To disable graceful shutdown, set to time.Duration(0)
	// To use graceful shutdown without timeout, set to a negative duration, e.G. time.Duration(-1)
	GracefulShutdownTimeout *time.Duration

	newHealthProbeListener func(addr string) (net.Listener, error)
}

// New returns a new Manager for creating Controllers.
// Note that if ContentType in the given config is not set, "application/vnd.kubernetes.protobuf"
// will be used for all built-in resources of Kubernetes, and "application/json" is for other types
// including all CRD resources.
func New(cfg *config.AppConfig, opts Options) (Manager, error) {
	if cfg == nil && !opts.Config.LoadFromConfigCenter {
		return nil, errors.New("must specify Config")
	}

	// Set default values for options fields
	opts = setOptionsDefaults(opts)

	upCh := make(chan discovery.ConfigUpdate, 1)
	if opts.Config.LoadFromConfigCenter {
		newCfg, err := loadConfigFromRemote(opts, upCh)
		if err != nil {
			return nil, err
		}
		if cfg == nil {
			cfg = &config.AppConfig{}
		}
		cfg = newCfg
	}

	if opts.Config.Registry.Register {
		cfg.Register.Register = true
	}

	if opts.Config.Registry.Discover {
		cfg.Register.Discover = true
	}

	if opts.Http.HTTPService == nil && opts.Grpc.GRPCService == nil {
		return nil, errors.New("must specify at least one of Http or Grpc")
	}

	hs := server.NewHttpService(cfg, opts.Http, opts.Http.HealthCheckAddress+opts.LivenessEndpointName, opts.Logger)

	// Create health probes listener. This will throw an error if the bind
	// address is invalid or already in use.
	httpHealthProbeListener, err := opts.newHealthProbeListener(opts.Http.HealthCheckAddress)
	if err != nil {
		return nil, err
	}

	errChan := make(chan error, 1)
	runnables := newRunnables(context.Background, errChan)

	return &controllerManager{
		stopProcedureEngaged:    new(int64),
		runnables:               runnables,
		errChan:                 errChan,
		httpServer:              hs,
		optsHttpServer:          opts.Http,
		readinessEndpointName:   opts.ReadinessEndpointName,
		livenessEndpointName:    opts.LivenessEndpointName,
		healthCheckAddress:      opts.Http.HealthCheckAddress,
		httpHealthProbeListener: httpHealthProbeListener,
		internalProceduresStop:  make(chan struct{}),
		configUpdateCh:          upCh,
		logger:                  opts.Logger,
		gracefulShutdownTimeout: *opts.GracefulShutdownTimeout,
		config:                  cfg,
	}, nil
}

func setOptionsDefaults(opts Options) Options {
	if opts.ReadinessEndpointName == "" {
		opts.ReadinessEndpointName = defaultReadinessEndpoint
	}

	if opts.LivenessEndpointName == "" {
		opts.LivenessEndpointName = defaultLivenessEndpoint
	}

	if opts.Http.HealthCheckAddress == "" {
		opts.Http.HealthCheckAddress = defaultHealthProbeBindAddress
	}

	if opts.newHealthProbeListener == nil {
		opts.newHealthProbeListener = defaultHealthProbeListener
	}

	if opts.GracefulShutdownTimeout == nil {
		gracefulShutdownTimeout := defaultGracefulShutdownPeriod
		opts.GracefulShutdownTimeout = &gracefulShutdownTimeout
	}

	if opts.Logger.GetSink() == nil {
		opts.Logger = zapr.NewLogger(plog.NewLogger("console", int8(zapcore.DebugLevel)))
	}

	return opts
}

// defaultHealthProbeListener creates the default health probes listener bound to the given address.
func defaultHealthProbeListener(addr string) (net.Listener, error) {
	if addr == "" || addr == "0" {
		return nil, nil
	}

	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("error listening on %s: %w", addr, err)
	}
	return ln, nil
}

func loadConfigFromRemote(opts Options, upCh chan<- discovery.ConfigUpdate) (*config.AppConfig, error) {
	cc, err := discovery.NewDefaultRemoteConfigManager(opts.Config.RemoteConfigAddr, discovery.WithToken(opts.Config.RemoteConfigToken))
	if err != nil {
		return nil, err
	}

	c, err := cc.Get(opts.Config.Key)
	if err != nil {
		return nil, err
	}

	if opts.Config.Hot {
		if err := cc.Watch(c, upCh, opts.Config.Key, opts.Config.Keys...); err != nil {
			return nil, err
		}
	}

	return c, nil
}

// BaseContextFunc is a function used to provide a base Context to Runnables
// managed by a Manager.
type BaseContextFunc func() context.Context

// Runnable allows a component to be started.
// It's very important that Start blocks until
// it's done running.
type Runnable interface {
	// Start starts running the component.  The component will stop running
	// when the context is closed. Start blocks until the context is closed or
	// an error occurs.
	Start(context.Context) error
}

// runnables handles all the runnables for a manager by grouping them accordingly to their
// type (webhooks, caches etc.).
type runnables struct {
	HTTPServers *runnableGroup
	GRPCServers *runnableGroup
}

//func (r *runnableGroup) Hot(ctx context.Context) {
//	fmt.Println("r.stopQueue => ", r.stopQueue)
//	for _, rn := range r.stopQueue {
//		rn.signalReady = true
//		rn.Stop(r.ctx)
//	}
//}

// StopAndWait waits for all the runnables to finish before returning.
func (r *runnableGroup) StopAndWait(ctx context.Context) {
	r.stopOnce.Do(func() {
		// Close the reconciler channel once we're done.
		defer close(r.ch)

		_ = r.Start(ctx)
		r.stop.Lock()
		// Store the stopped variable so we don't accept any new
		// runnables for the time being.
		r.stopped = true
		r.stop.Unlock()

		// Cancel the internal channel.
		r.cancel()

		done := make(chan struct{})
		go func() {
			defer close(done)
			// Wait for all the runnables to finish.
			r.wg.Wait()
		}()

		select {
		case <-done:
			// We're done, exit.
		case <-ctx.Done():
			// Calling context has expired, exit.
		}
	})
}

// Start starts the group and waits for all
// initially registered runnables to start.
// It can only be called once, subsequent calls have no effect.
func (r *runnableGroup) Start(ctx context.Context) error {
	var retErr error

	r.startOnce.Do(func() {
		defer close(r.startReadyCh)

		// Start the internal reconciler.
		go r.reconcile()

		// Start the group and queue up all
		// the runnables that were added prior.
		r.start.Lock()
		r.started = true
		for _, rn := range r.startQueue {
			rn.signalReady = true
			r.ch <- rn
		}
		r.start.Unlock()

		// If we don't have any queue, return.
		if len(r.startQueue) == 0 {
			return
		}

		// Wait for all runnables to signal.
		for {
			select {
			case <-ctx.Done():
				fmt.Println("11111")
				if err := ctx.Err(); !errors.Is(err, context.Canceled) {
					retErr = err
				}
			case rn := <-r.startReadyCh:
				for i, existing := range r.startQueue {
					if existing == rn {
						// Remove the item from the start queue.
						r.startQueue = append(r.startQueue[:i], r.startQueue[i+1:]...)
						break
					}
				}
				// We're done waiting if the queue is empty, return.
				if len(r.startQueue) == 0 {
					return
				}
			}
		}
	})

	return retErr
}

// reconcile is our main entrypoint for every runnable added
// to this group. Its primary job is to read off the internal channel
// and schedule runnables while tracking their state.
func (r *runnableGroup) reconcile() {
	for runnable := range r.ch {
		// Handle stop.
		// If the shutdown has been called we want to avoid
		// adding new goroutines to the WaitGroup because Wait()
		// panics if Add() is called after it.
		{
			r.stop.RLock()
			if r.stopped {
				// Drop any runnables if we're stopped.
				r.errChan <- errRunnableGroupStopped
				r.stop.RUnlock()
				continue
			}

			// Why is this here?
			// When StopAndWait is called, if a runnable is in the process
			// of being added, we could end up in a situation where
			// the WaitGroup is incremented while StopAndWait has called Wait(),
			// which would result in a panic.
			r.wg.Add(1)
			r.stop.RUnlock()
		}

		// Start the runnable.
		go func(rn *readyRunnable) {
			go func() {
				if rn.Check(r.ctx) {
					if rn.signalReady {
						r.startReadyCh <- rn
					}
				}
			}()

			// If we return, the runnable ended cleanly
			// or returned an error to the channel.
			//
			// We should always decrement the WaitGroup here.
			defer r.wg.Done()

			// Start the runnable.
			if err := rn.Start(r.ctx); err != nil {
				r.errChan <- err
			}
		}(runnable)
	}
}

var (
	errRunnableGroupStopped = errors.New("can't accept new runnable as stop procedure is already engaged")
)
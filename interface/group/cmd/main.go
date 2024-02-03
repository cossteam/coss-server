package main

import (
	"flag"
	_ "github.com/cossim/coss-server/docs"
	"github.com/cossim/coss-server/interface/group/config"
	"github.com/cossim/coss-server/interface/group/server/http"
	"github.com/cossim/coss-server/interface/group/service"
	"github.com/cossim/coss-server/pkg/discovery"
	"github.com/cossim/coss-server/pkg/version"
	"os"
	"os/signal"
	"syscall"
)

var (
	file              string
	discover          bool
	remoteConfig      bool
	remoteConfigAddr  string
	remoteConfigToken string
)

func init() {
	flag.StringVar(&file, "config", "/config/config.yaml", "Path to configuration file")
	flag.BoolVar(&discover, "discover", false, "Enable service discovery")
	flag.BoolVar(&remoteConfig, "remote-config", false, "Load configuration from remote source")
	flag.StringVar(&remoteConfigAddr, "config-center-addr", "", "Address of the configuration center")
	flag.StringVar(&remoteConfigToken, "config-center-token", "", "Token for accessing the configuration center")

	flag.Parse()
}

func main() {
	version.FullVersion()

	ch := make(chan discovery.ConfigUpdate)
	var err error
	if !remoteConfig {
		if err = config.LoadConfigFromFile(file); err != nil {
			panic(err)
		}
	} else {
		ch, err = discovery.LoadDefaultRemoteConfig(remoteConfigAddr, discovery.InterfaceConfigPrefix+"group_relation", remoteConfigToken, config.Conf)
		if err != nil {
			panic(err)
		}
	}

	if config.Conf == nil {
		panic("Config not initialized")
	}

	svc := service.New()
	http.Start(svc)
	svc.Start(discover)

	go func() {
		for {
			select {
			case _ = <-ch:
				http.Stop()
				svc.Stop(discover)
				svc = service.Restart(discover)
				http.Restart(svc)
			}
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			http.Stop()
			svc.Stop(discover)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

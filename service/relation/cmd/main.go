package main

import (
	"flag"
	"fmt"
	"github.com/cossim/coss-server/pkg/db"
	"github.com/cossim/coss-server/pkg/discovery"
	dialoggrpcv1 "github.com/cossim/coss-server/service/relation/api/v1/dialog"
	relationgrpcv1 "github.com/cossim/coss-server/service/relation/api/v1/group_join_request"
	grouprelationgrpcv1 "github.com/cossim/coss-server/service/relation/api/v1/group_relation"
	userfriendgrpcv1 "github.com/cossim/coss-server/service/relation/api/v1/user_friend_request"
	userrelationgrpcv1 "github.com/cossim/coss-server/service/relation/api/v1/user_relation"
	"github.com/cossim/coss-server/service/relation/config"
	"github.com/cossim/coss-server/service/relation/infrastructure/persistence"
	"github.com/cossim/coss-server/service/relation/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
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

	grpcServer *grpc.Server
	svc        *service.Service
	lis        net.Listener
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
	ch := make(chan discovery.ConfigUpdate)
	var err error
	if !remoteConfig {
		if err = config.LoadConfigFromFile(file); err != nil {
			panic(err)
		}
	} else {
		ch, err = discovery.LoadDefaultRemoteConfig(remoteConfigAddr, discovery.ServiceConfigPrefix+"relation", remoteConfigToken, config.Conf)
		if err != nil {
			panic(err)
		}
	}

	if config.Conf == nil {
		panic("Config not initialized")
	}

	// 启动 gRPC 服务器
	startGRPCServer()

	go func() {
		for {
			select {
			case _ = <-ch:
				log.Printf("Config updated, restarting service")
				grpcServer.Stop()
				svc.Stop(discover)
				startGRPCServer()
				svc.Start(discover)
			}
		}
	}()

	svc.Start(discover)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			svc.Stop(discover)
			grpcServer.Stop()
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}

func startGRPCServer() {
	lisAddr := fmt.Sprintf("%s", config.Conf.GRPC.Addr())
	var err error
	lis, err = net.Listen("tcp", lisAddr)
	if err != nil {
		panic(err)
	}

	dbConn, err := db.NewMySQLFromDSN(config.Conf.MySQL.DSN).GetConnection()
	if err != nil {
		panic(err)
	}

	infra := persistence.NewRepositories(dbConn)
	if err = infra.Automigrate(); err != nil {
		panic(err)
	}

	grpcServer = grpc.NewServer()
	svc = service.NewService(infra, dbConn, config.Conf)
	userrelationgrpcv1.RegisterUserRelationServiceServer(grpcServer, svc)
	grouprelationgrpcv1.RegisterGroupRelationServiceServer(grpcServer, svc)
	dialoggrpcv1.RegisterDialogServiceServer(grpcServer, svc)
	relationgrpcv1.RegisterGroupJoinRequestServiceServer(grpcServer, svc)
	userfriendgrpcv1.RegisterUserFriendRequestServiceServer(grpcServer, svc)
	// 注册服务开启健康检查
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())

	fmt.Printf("gRPC server is running on addr: %s\n", lisAddr)

	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()
}

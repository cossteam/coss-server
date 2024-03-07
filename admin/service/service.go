package service

import (
	"context"
	"fmt"
	"github.com/cossim/coss-server/admin/infrastructure/persistence"
	pkgconfig "github.com/cossim/coss-server/pkg/config"
	"github.com/cossim/coss-server/pkg/db"
	plog "github.com/cossim/coss-server/pkg/log"
	"github.com/cossim/coss-server/pkg/msg_queue"
	"github.com/cossim/coss-server/pkg/storage"
	"github.com/cossim/coss-server/pkg/storage/minio"
	"github.com/cossim/coss-server/pkg/utils/os"
	msggrpcv1 "github.com/cossim/coss-server/service/msg/api/v1"
	relationgrpcv1 "github.com/cossim/coss-server/service/relation/api/v1"
	user "github.com/cossim/coss-server/service/user/api/v1"
	"github.com/rs/xid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type Service struct {
	ac             *pkgconfig.AppConfig
	repo           *persistence.Repositories
	logger         *zap.Logger
	userClient     user.UserServiceClient
	relationClient relationgrpcv1.UserRelationServiceClient
	msgClient      msggrpcv1.MsgServiceClient
	rabbitMQClient *msg_queue.RabbitMQ
	sp             storage.StorageProvider

	dtmGrpcServer  string
	userGrpcServer string

	gatewayAddress string
	gatewayPort    string

	sid         string
	appPath     string
	downloadURL string
}

func New(ac *pkgconfig.AppConfig) (s *Service) {
	s = &Service{
		ac:             ac,
		logger:         plog.NewDefaultLogger("admin_bff", int8(ac.Log.Level)),
		sid:            xid.New().String(),
		rabbitMQClient: setRabbitMQProvider(ac),
		dtmGrpcServer:  ac.Dtm.Addr(),
		sp:             setMinIOProvider(ac),
		downloadURL:    "/api/v1/storage/files/download",
	}
	s.setLoadSystem()
	s.setupDBConn()
	return s
}

func (s *Service) HandlerGrpcClient(serviceName string, conn *grpc.ClientConn) error {
	switch serviceName {
	case "user_service":
		s.userClient = user.NewUserServiceClient(conn)
		s.userGrpcServer = conn.Target()
		err := s.InitAdmin()
		if err != nil {
			return err
		}
		s.logger.Info("gRPC client for user service initialized", zap.String("addr", conn.Target()))
	case "relation_service":
		s.relationClient = relationgrpcv1.NewUserRelationServiceClient(conn)
		s.logger.Info("gRPC client for relation service initialized", zap.String("addr", conn.Target()))
	case "msg_service":
		s.msgClient = msggrpcv1.NewMsgServiceClient(conn)
		s.logger.Info("gRPC client for msg service initialized", zap.String("addr", conn.Target()))
	}
	err := s.InitAdmin()
	if err != nil {
		panic(err)
	}
	return nil
}

func setRabbitMQProvider(ac *pkgconfig.AppConfig) *msg_queue.RabbitMQ {
	rmq, err := msg_queue.NewRabbitMQ(fmt.Sprintf("amqp://%s:%s@%s", ac.MessageQueue.Username, ac.MessageQueue.Password, ac.MessageQueue.Addr()))
	if err != nil {
		panic(err)
	}
	return rmq
}

func (s *Service) Stop(ctx context.Context) error {
	return nil
}

func (s *Service) setupDBConn() {
	dbConn, err := db.NewMySQLFromDSN(s.ac.MySQL.DSN).GetConnection()
	if err != nil {
		panic(err)
	}
	infra := persistence.NewRepositories(dbConn)
	if err = infra.Automigrate(); err != nil {
		panic(err)
	}
	s.repo = infra
}

func (s *Service) setLoadSystem() {

	env := s.ac.SystemConfig.Environment
	if env == "" {
		env = "dev"
	}

	switch env {
	case "prod":
		path := s.ac.SystemConfig.AvatarFilePath
		if path == "" {
			path = "/.catch/"
		}
		s.appPath = path

		gatewayAdd := s.ac.SystemConfig.GatewayAddress
		if gatewayAdd == "" {
			gatewayAdd = "43.229.28.107"
		}

		s.gatewayAddress = gatewayAdd

		gatewayPo := s.ac.SystemConfig.GatewayPort
		if gatewayPo == "" {
			gatewayPo = "8080"
		}
		s.gatewayPort = gatewayPo
	default:
		path := s.ac.SystemConfig.AvatarFilePathDev
		if path == "" {
			npath, err := os.GetPackagePath()
			if err != nil {
				panic(err)
			}
			path = npath + "deploy/docker/config/common/"
		}
		s.appPath = path

		gatewayAdd := s.ac.SystemConfig.GatewayAddressDev
		if gatewayAdd == "" {
			gatewayAdd = "127.0.0.1"
		}

		s.gatewayAddress = gatewayAdd

		gatewayPo := s.ac.SystemConfig.GatewayPortDev
		if gatewayPo == "" {
			gatewayPo = "8080"
		}
		s.gatewayPort = gatewayPo
	}
	s.gatewayAddress = s.gatewayAddress + ":" + s.gatewayPort
}

func setMinIOProvider(ac *pkgconfig.AppConfig) storage.StorageProvider {
	var err error
	sp, err := minio.NewMinIOStorage(ac.OSS["minio"].Addr(), ac.OSS["minio"].AccessKey, ac.OSS["minio"].SecretKey, ac.OSS["minio"].SSL)
	if err != nil {
		panic(err)
	}

	return sp
}

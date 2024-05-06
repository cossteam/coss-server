package grpc

import (
	"context"
	"errors"
	api "github.com/cossim/coss-server/internal/user/api/grpc/v1"
	"github.com/cossim/coss-server/internal/user/cache"
	"github.com/cossim/coss-server/internal/user/domain/entity"
	"github.com/cossim/coss-server/internal/user/domain/repository"
	"github.com/cossim/coss-server/pkg/code"
	"github.com/cossim/coss-server/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

//var _ api.UserAuthServiceServer = &authServiceServer{}

type authServiceServer struct {
	secret    string
	userCache cache.UserCache
	ur        repository.UserRepository
}

func (s *UserServiceServer) ParseToken(ctx context.Context, request *api.ParseTokenRequest) (*api.ParseTokenResponse, error) {
	_, claims, err := utils.ParseToken(request.Token, s.secret)
	if err != nil {
		return nil, err
	}

	return &api.ParseTokenResponse{
		UserID:    claims.UserId,
		Email:     claims.Email,
		DriverID:  claims.DriverId,
		PublicKey: claims.PublicKey,
	}, nil
}

func (s *UserServiceServer) GenerateUserToken(ctx context.Context, request *api.GenerateUserTokenRequest) (*api.GenerateUserTokenResponse, error) {
	token, err := utils.GenerateToken(request.UserID, request.Email, request.DriverID, request.PublicKey, s.secret)
	if err != nil {
		return nil, err
	}

	return &api.GenerateUserTokenResponse{Token: token}, nil
}

func (s *UserServiceServer) Access(ctx context.Context, request *api.AccessRequest) (*api.AccessResponse, error) {
	resp := &api.AccessResponse{}

	_, claims, err := utils.ParseToken(request.Token, s.secret)
	if err != nil {
		return nil, err
	}

	infos, err := s.userCache.GetUserLoginInfos(ctx, claims.UserId)
	if err == nil {
		var found bool
		for _, v := range infos {
			if v.Token == request.Token {
				found = true
				break
			}
		}

		if !found {
			return nil, code.Unauthorized
		}

		return resp, nil
	}

	info, err := s.ur.GetUserInfoByUid(ctx, claims.UserId)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return resp, status.Error(codes.Code(code.UserErrNotExistOrPassword.Code()), err.Error())
		}
		return resp, status.Error(codes.Code(code.UserErrGetUserInfoFailed.Code()), err.Error())
	}

	if info.Status != entity.UserStatusNormal {
		return nil, code.UserErrStatusException
	}

	return resp, nil
}

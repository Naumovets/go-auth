package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/Naumovets/go-auth/internal/entities"
	"github.com/Naumovets/go-auth/internal/utils"
	desc "github.com/Naumovets/go-auth/pkg/auth_v1"
)

func (s *serverAuth) Register(ctx context.Context, req *desc.RegisterRequest) (*desc.RegisterResponse, error) {

	exists, err := s.rep.ExistsUser(req.GetUsername())

	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errors.New("error register: user already exists")
	}

	user := &entities.User{
		Username: req.GetUsername(),
		Password: req.GetPassword(),
	}

	err = s.rep.AddUser(user)

	if err != nil {
		return nil, err
	}

	refresh_token, err := utils.GenerateToken(
		*user,
		[]byte(s.cfg.refreshTokenSecretKey),
		refreshTokenExpiration,
	)

	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	return &desc.RegisterResponse{RefreshToken: refresh_token}, nil
}
func (s *serverAuth) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {

	user, err := s.rep.GetUserByUsername(req.GetUsername())

	if err != nil {
		return nil, err
	}

	if !utils.VerifyPassword(user.Password, req.GetPassword()) {
		return nil, fmt.Errorf("error login: password no verify")
	}

	refresh_token, err := utils.GenerateToken(
		*user,
		[]byte(s.cfg.refreshTokenSecretKey),
		refreshTokenExpiration,
	)

	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	return &desc.LoginResponse{RefreshToken: refresh_token}, nil
}
func (s *serverAuth) GetRefreshToken(ctx context.Context, req *desc.GetRefreshTokenRequest) (*desc.GetRefreshTokenResponse, error) {
	claims, err := utils.VerifyToken(req.GetRefreshToken(), []byte(s.cfg.refreshTokenSecretKey))
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	refreshToken, err := utils.GenerateToken(entities.User{
		Username: claims.Username,
		Id:       claims.Id,
	},
		[]byte(s.cfg.refreshTokenSecretKey),
		refreshTokenExpiration,
	)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}
	return &desc.GetRefreshTokenResponse{RefreshToken: refreshToken}, nil
}
func (s *serverAuth) GetAccessToken(ctx context.Context, req *desc.GetAccessTokenRequest) (*desc.GetAccessTokenResponse, error) {
	claims, err := utils.VerifyToken(req.GetRefreshToken(), []byte(s.cfg.refreshTokenSecretKey))
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	accessToken, err := utils.GenerateToken(entities.User{
		Username: claims.Username,
		Id:       claims.Id,
	},
		[]byte(s.cfg.accessTokenSecretKey),
		accessTokenExpiration,
	)
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}
	return &desc.GetAccessTokenResponse{AccessToken: accessToken}, nil
}
func (s *serverAuth) GetUsersById(ctx context.Context, req *desc.GetUsersByIdRequest) (*desc.GetUsersByIdResponse, error) {
	users, err := s.rep.GetUsersByIds(req.Ids)
	if err != nil {
		return nil, err
	}

	resUser := make([]*desc.User, 0)

	for _, user := range users {
		resUser = append(resUser, &desc.User{
			Username: user.Username,
			Id:       user.Id,
		})
	}

	return &desc.GetUsersByIdResponse{
		Users: resUser,
	}, nil
}
func (s *serverAuth) GetUserInfo(ctx context.Context, req *desc.GetUserInfoRequest) (*desc.GetUserInfoResponse, error) {
	claims, err := utils.VerifyToken(req.GetAccessToken(), []byte(s.cfg.refreshTokenSecretKey))
	if err != nil {
		return nil, fmt.Errorf("error: %s", err)
	}

	return &desc.GetUserInfoResponse{
		User: &desc.User{
			Id:       claims.Id,
			Username: claims.Username,
		},
	}, nil
}

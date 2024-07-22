package auth

import (
	"time"

	"github.com/Naumovets/go-auth/internal/repositories"
	desc "github.com/Naumovets/go-auth/pkg/auth_v1"
)

const (
	authPrefix             = "Bearer "
	refreshTokenExpiration = 30 * 25 * time.Hour
	accessTokenExpiration  = 60 * time.Minute
)

type serverAuth struct {
	desc.AuthV1Server
	rep *repositories.Repository
	cfg *Config
}

func NewServerAuth(rep *repositories.Repository, cfg *Config) *serverAuth {
	return &serverAuth{
		rep: rep,
		cfg: cfg,
	}
}

package service

import (
	"github.com/arakhimiy/edu-connect/internal/repository"
	"github.com/go-telegram/bot"
)

type AuthorizationI interface {
}

type I interface {
	Authorization() AuthorizationI
}

type service struct {
	AuthorizationI
}

func (s *service) Authorization() AuthorizationI {
	return s.AuthorizationI
}

func NewService(repo repository.I, bot *bot.Bot) I {
	return &service{
		AuthorizationI: NewAuthorizationS(repo.Authorization(), bot),
	}
}

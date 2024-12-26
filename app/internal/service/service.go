package service

import (
	"gitlab.saidoff.uz/company/muslim-administration/mosque/back/internal/repository"
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

func NewService(repo repository.I) I {
	return &service{
		AuthorizationI: NewAuthorizationS(repo.Authorization()),
	}
}

package service

import "gitlab.saidoff.uz/company/muslim-administration/mosque/back/internal/repository"

type AuthorizationS struct {
	repo repository.AuthorizationI
}

func NewAuthorizationS(repo repository.AuthorizationI) *AuthorizationS {
	return &AuthorizationS{repo: repo}
}

package service

import (
	"github.com/arakhimiy/edu-connect/internal/repository"
	"github.com/go-telegram/bot"
)

type AuthorizationS struct {
	repo repository.AuthorizationI
	bot  *bot.Bot
}

func NewAuthorizationS(repo repository.AuthorizationI, bot *bot.Bot) *AuthorizationS {
	return &AuthorizationS{repo: repo, bot: bot}
}

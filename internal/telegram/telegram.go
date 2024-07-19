package telegram

import (
	"context"
	"log"

	"gopkg.in/telebot.v3"

	"github.com/Anton-Kraev/wb-stats-bot/internal/domain"
)

const (
	startCommand = "/start"
	tokenCommand = "/token"
)

type tokenRepository interface {
	Upsert(ctx context.Context, chatID int64, token domain.Token) error
	Get(ctx context.Context, chatID int64) (domain.Token, error)
}

type Bot struct {
	tg        *telebot.Bot
	tokenRepo tokenRepository
}

func NewBot(tg *telebot.Bot, tokenRepo tokenRepository) *Bot {
	return &Bot{tg: tg, tokenRepo: tokenRepo}
}

func (b *Bot) Start() {
	log.Printf("Authorized on account %s", b.tg.Me.Username)

	b.SetHandlers()
	b.tg.Start()
}

func (b *Bot) SetHandlers() {
	b.tg.Handle(telebot.OnText, b.handleUnknownCommand)
	b.tg.Handle(startCommand, b.handleStartCommand)
	b.tg.Handle(tokenCommand, b.handleTokenCommand)
}

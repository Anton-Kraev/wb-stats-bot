package telegram

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/Anton-Kraev/wb-stats-bot/internal/domain"
)

type tokenRepository interface {
	Upsert(ctx context.Context, chatID int64, token domain.Token) error
	Get(ctx context.Context, chatID int64) (domain.Token, error)
}

type Bot struct {
	api       *tgbotapi.BotAPI
	tokenRepo tokenRepository
}

func NewBot(api *tgbotapi.BotAPI, tokenRepo tokenRepository) *Bot {
	return &Bot{api: api, tokenRepo: tokenRepo}
}

func (b *Bot) Start() {
	log.Printf("Authorized on account %s", b.api.Self.UserName)

	updates := b.initUpdatesChannel()
	b.handleUpdates(updates)
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	updConfig := tgbotapi.NewUpdate(0)
	updConfig.Timeout = 60

	return b.api.GetUpdatesChan(updConfig)
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			b.handleCommand(update.Message)
		} else {
			b.handleMessage(update.Message)
		}
	}
}

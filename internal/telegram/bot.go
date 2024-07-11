package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	api *tgbotapi.BotAPI
}

func NewBot(api *tgbotapi.BotAPI) *Bot {
	return &Bot{api: api}
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

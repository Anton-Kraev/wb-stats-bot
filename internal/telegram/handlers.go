package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	wbAuthURL         = "https://openapi.wb.ru/general/authorization/ru/"
	startCommand      = "start"
	replyStartCommand = "Привет! Для начала работы с ботом вам необходимо добавить токен для доступа к Wildberries API. \n\nИнструкцию по созданию токена можно найти по ссылке: " + wbAuthURL
)

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, replyStartCommand)
	_, err := b.api.Send(msg)

	return err
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Неизвестная команда")
	_, err := b.api.Send(msg)

	return err
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case startCommand:
		return b.handleStartCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}
}

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	msg.ReplyToMessageID = message.MessageID

	_, err := b.api.Send(msg)

	return err
}

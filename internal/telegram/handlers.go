package telegram

import (
	"context"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/Anton-Kraev/wb-stats-bot/internal/domain"
)

const (
	wbAuthURL         = "https://openapi.wb.ru/general/authorization/ru/"
	startCommand      = "start"
	replyStartCommand = "Привет! Для начала работы с ботом вам необходимо добавить токен для доступа к Wildberries API. \n\nИнструкцию по созданию токена можно найти по ссылке: " + wbAuthURL + "\n\nДля добавления токена в бота используйте команду /token: /token <ваш-токен>"

	tokenCommand      = "token"
	replyTokenCommand = "Токен успешно сохранен!"
	emptyToken        = "После команды /token через пробел необходимо передать ваш API токен для сохранения"
)

func (b *Bot) handleStartCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, replyStartCommand)
	_, err := b.api.Send(msg)

	return err
}

func (b *Bot) handleTokenCommand(message *tgbotapi.Message) error {
	var (
		ctx    = context.Background()
		chatID = message.Chat.ID
		token  = domain.Token(message.CommandArguments())
		msg    = tgbotapi.NewMessage(chatID, replyTokenCommand)
	)

	if len(token) == 0 {
		msg.Text = emptyToken
		_, err := b.api.Send(msg)

		return err
	}

	if err := b.tokenRepo.Upsert(ctx, chatID, token); err != nil {
		msg.Text = err.Error()
		_, err = b.api.Send(msg)

		return err
	}

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
	case tokenCommand:
		return b.handleTokenCommand(message)
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

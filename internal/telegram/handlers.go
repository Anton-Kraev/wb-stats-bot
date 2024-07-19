package telegram

import (
	"context"

	"gopkg.in/telebot.v3"

	"github.com/Anton-Kraev/wb-stats-bot/internal/domain"
)

const (
	wbAuthURL         = "https://openapi.wb.ru/general/authorization/ru/"
	replyStartCommand = "Привет! Для начала работы с ботом вам необходимо добавить токен для доступа к Wildberries API. \n\nИнструкцию по созданию токена можно найти по ссылке: " + wbAuthURL + "\n\nДля добавления токена в бота используйте команду /token: /token <ваш-токен>"

	replyTokenCommand = "Токен успешно сохранен!"
	emptyToken        = "После команды /token через пробел необходимо передать ваш API токен для сохранения"
)

func (b *Bot) handleStartCommand(ctx telebot.Context) error {
	return ctx.Send(replyStartCommand)
}

func (b *Bot) handleTokenCommand(ctx telebot.Context) error {
	if len(ctx.Args()) == 0 {
		return ctx.Send(emptyToken)
	}

	var (
		repoContext = context.Background()
		chatID      = ctx.Chat().ID
		token       = domain.Token(ctx.Args()[0])
	)

	if err := b.tokenRepo.Upsert(repoContext, chatID, token); err != nil {
		return ctx.Send(err.Error())
	}

	return ctx.Send(replyTokenCommand)
}

func (b *Bot) handleUnknownCommand(ctx telebot.Context) error {
	return ctx.Send("Неизвестная команда")
}

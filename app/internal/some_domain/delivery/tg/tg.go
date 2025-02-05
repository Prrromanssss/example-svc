package tg

import (
	"context"
	"example-svc/internal/some_domain/delivery"
	"example-svc/internal/some_domain/usecases/models"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TgBot struct {
	BotAPI *tgbotapi.BotAPI
	tgUc   delivery.SomeDomain
}

func NewTgBot(
	usecase delivery.SomeDomain,
	botAPI *tgbotapi.BotAPI,
) *TgBot {
	return &TgBot{
		BotAPI: botAPI,
		tgUc:   usecase,
	}
}

func (t *TgBot) Run() error {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = UpdateConfigTime

	for update := range t.BotAPI.GetUpdatesChan(updateConfig) {
		if msg := update.Message; msg != nil {
			if msg.IsCommand() {
				//
				t.SomeCommandHandler(models.CreateCommand{
					Param1: "param1",
					Param2: "param2",
				})
			} else {
				//
				t.SomeNonCommandHandler()
			}
		}
		if cq := update.CallbackQuery; cq != nil {
			//
			t.SomeCallbackQueryHandler()
		}
	}

	return nil
}

func (t *TgBot) SomeCommandHandler(command models.CreateCommand) {
	id, err := t.tgUc.SomeMethodCreate(context.Background(), command)
	if err != nil {
		return
	}

	chatID, err := t.tgUc.GetChatIDByUserID(context.Background(), 1)
	if err != nil {
		return
	}

	_, err = t.BotAPI.Send(tgbotapi.NewMessage(chatID, fmt.Sprintf("Created: %d", id)))
	if err != nil {
		return
	}
}

func (t *TgBot) SomeNonCommandHandler() {

}

func (t *TgBot) SomeCallbackQueryHandler() {

}

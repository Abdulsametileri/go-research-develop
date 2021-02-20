package services

import (
	"log"

	"github.com/Abdulsametileri/ingilizce-kelime-go/config"
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct{}

func (t *Telegram) setup() {
	telegramConfig := config.Get().Telegram
	botToken := telegramConfig.BotToken
	chatId := telegramConfig.AbdulsametTelegramId

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = false

	poolConfig := tgbotapi.SendPollConfig{
		BaseChat: tgbotapi.BaseChat{
			ChatID:           chatId,
			ReplyToMessageID: 0,
		},
		Question:        "Soru?",
		Type:            "quiz",
		Options:         []string{"Cevap 1", "Cevap 2", "Cevap 3", "Cevap 4"},
		IsAnonymous:     true,
		CorrectOptionID: 1,
		Explanation:     "huhuhucan",
	}

	_, err = bot.Send(poolConfig)
}

package services

// v5.0.0-rc1.0.20200424181826-774f1e72e764
import (
	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"

	"github.com/Abdulsametileri/ingilizce-kelime-go/config"
)

type Telegram struct{}

func (t *Telegram) setup() {
	telegramConfig := config.Get().TelegramConfig
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

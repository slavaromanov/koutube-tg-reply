package tg

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type Token string
type Converter interface {
	ConvertVideoURL(s string) (bool, string)
}

type Bot struct {
	token  Token
	conv   Converter
	logger *zap.Logger
}

func New(token Token, conv Converter, logger *zap.Logger) *Bot {
	return &Bot{
		token:  token,
		conv:   conv,
		logger: logger,
	}
}

func (b *Bot) Run(ctx context.Context) error {
	bot, err := tgbotapi.NewBotAPI(string(b.token))
	if err != nil {
		return err
	}
	msgCh := bot.GetUpdatesChan(tgbotapi.NewUpdate(0))
	for {
		select {
		case <-ctx.Done():
			return nil
		case update := <-msgCh:
			msg := update.Message
			if msg == nil {
				continue
			}
			ok, url := b.conv.ConvertVideoURL(msg.Text)
			if !ok {
				continue
			}
			outMsg := tgbotapi.NewMessage(msg.Chat.ID, url)
			outMsg.ReplyToMessageID = msg.MessageID
			if _, err = bot.Send(outMsg); err != nil {
				b.logger.Error("failed to send message", zap.Error(err))
			}
		}
	}
}

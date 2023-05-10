package handlers

import (
	"errors"

	"github.com/Totus-Floreo/tg-pwd-manager/pkg/application"
	"github.com/mymmrac/telego"
	"github.com/mymmrac/telego/telegohandler"
	"go.uber.org/zap"
)

type Handler struct {
	PassManService application.PassManService
	Logger         *zap.SugaredLogger
}

func (h *Handler) Start() telegohandler.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		exists, err := h.PassManService.GetUserExists(update.Message.From.ID)
		if err != nil {
			h.Logger.Error(err)
		}
		h.Logger.Error(errors.New("ErrUserExists"))
		if !exists {
			if err := h.PassManService.CreateUser(update.Message.From.ID); err != nil {
				h.Logger.Error(err)
			}
		}
	}
}

func (h *Handler) Set() telegohandler.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		if user, err := h.PassManService.GetUserByID(update.Message.From.ID); err != nil {
			h.Logger.Error(err)
		} else {
			err = h.PassManService.CreateCredentials(user, update.Message.Text)
			if err != nil {
				h.Logger.Error(err)
			}
		}
	}
}

func (h *Handler) Get() telegohandler.Handler {
	return func(bot *telego.Bot, update telego.Update) {
		if user, err := h.PassManService.GetUserByID(update.Message.From.ID); err != nil {
			h.Logger.Error(err)
		} else {
			cred, err := h.PassManService.GetUserCredentials(user, update.Message.Text)
			if err != nil {
				h.Logger.Error(err)
			}
			msg := &telego.SendMessageParams{}
			chatID := telego.ChatID{update.Message.From.ID, "@" + update.Message.From.Username}
			msg = msg.WithChatID(chatID)
			msg = msg.WithText(cred.Login + cred.Password)
			bot.SendMessage(msg)
		}
	}
}

func (h *Handler) Delete() telegohandler.Handler {}

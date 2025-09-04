package handler

import (
	"fmt"
	"log"
	"math"
	"telegram-weather-bot/clients/openWeather"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Handler struct {
	bot      *tgbotapi.BotAPI
	owClient *openWeather.OpenWeatherClient
}

func New(bot *tgbotapi.BotAPI, owClient *openWeather.OpenWeatherClient) *Handler {
	return &Handler{
		bot:      bot,
		owClient: owClient,
	}
}

func (h *Handler) Start() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := h.bot.GetUpdatesChan(u)

	for update := range updates {
		h.handleUpdate(update)
	}
}

func (h *Handler) handleUpdate(update tgbotapi.Update) {
	if update.Message == nil {
		return
	}

	if update.Message.Text == "/start" {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf(
			"Приветствую %s %s!\nЧтобы узнать погоду, напишите город в формате: \"Москва\"", update.Message.From.FirstName, update.Message.From.LastName),
		)
		h.bot.Send(msg)
	}

	if update.Message.Text[0:1] == "/" && update.Message.Text != "start" {
		return
	}

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
	coordinates, err := h.owClient.Coordinates(update.Message.Text)
	if err != nil {
		log.Println(err)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не смогли получить координаты")
		msg.ReplyToMessageID = update.Message.MessageID
		h.bot.Send(msg)
		return
	}

	weather, err := h.owClient.Weather(coordinates.Lat, coordinates.Lon)
	if err != nil {
		log.Println(err)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Не смогли получить погоду в этой местности")
		msg.ReplyToMessageID = update.Message.MessageID
		h.bot.Send(msg)
		return
	}

	msg := tgbotapi.NewMessage(
		update.Message.Chat.ID,
		fmt.Sprintf("Температура в %s: %d°C", update.Message.Text, int(math.Round(weather.Temp))),
	)
	msg.ReplyToMessageID = update.Message.MessageID

	h.bot.Send(msg)
}

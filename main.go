package main

import (
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

var (
	bot *tgbotapi.BotAPI
	err error
)

func init() {
	config := LoadConfig()
	bot, err = tgbotapi.NewBotAPI(config.Token)
	if err != nil {
		log.Panic(err)
	}
}

func main() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		lat, lon := FindCoordinates(update.Message.Text)

		if lat != "" && lon != "" {
			fLat, err := strconv.ParseFloat(lat, 32)
			if err != nil {
				continue
			}
			fLon, err := strconv.ParseFloat(lon, 32)
			if err != nil {
				continue
			}

			locMsg := tgbotapi.NewLocation(update.Message.Chat.ID, fLat, fLon)

			gMapsUrl := fmt.Sprintf("https://www.google.com/maps/search/?api=1&query=%s,%s", lat, lon)
			gMapsButton := tgbotapi.InlineKeyboardButton{
				Text: "GoogleMaps",
				URL:  &gMapsUrl,
			}

			twoGisUrl := "https://2gis.com.cy/cyprus/geo/" + lon + "%2c" + lat
			twoGisButton := tgbotapi.InlineKeyboardButton{
				Text: "2GIS",
				URL:  &twoGisUrl,
			}

			mapsMeUrl := fmt.Sprintf("https://dlink.maps.me/map?v=1&ll=%s,%s", lat, lon)
			mapsMeButton := tgbotapi.InlineKeyboardButton{
				Text: "MAPS.ME",
				URL:  &mapsMeUrl,
			}

			buttons := tgbotapi.NewInlineKeyboardRow(gMapsButton, twoGisButton, mapsMeButton)

			locMsg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(buttons)

			_, err = bot.Send(locMsg)
			if err != nil {
				log.Println(err)
			}
		}
	}
}

package commands

import (
	"fmt"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

// TimeCommand will handle a /time command from a chat, and print back the current server time.
type TimeCommand struct {
	bot *tgbotapi.BotAPI
}

// SetBotAPI is used to make the bot api available for the handler.
func (hc *TimeCommand) SetBotAPI(bot *tgbotapi.BotAPI) {
	hc.bot = bot
}

// IsCommandMatch will check if the message string contains an help command.
func (hc *TimeCommand) IsCommandMatch(update *tgbotapi.Update) bool {
	return update.Message.IsCommand() && strings.ToLower(update.Message.Command()) == "time"
}

// PreProcessText does nothing to the message as its not used.
func (hc *TimeCommand) PreProcessText(update *tgbotapi.Update) error {
	return nil
}

// Execute will run the help command towards the chat where the command was posted.
func (hc *TimeCommand) Execute(update *tgbotapi.Update) error {
	hour, min, sec := time.Now().Clock()
	timestring := fmt.Sprintf("%02d:%02d:%02d", hour, min, sec)
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, timestring)
	_, err := hc.bot.Send(msg)

	return err
}

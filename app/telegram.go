package app

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"io"
	"log"
)

// https://api.telegram.org/6847302523:AAGH73xiPrQhKaP5n76IItYpYGG0v_0Dffo/setWebhook?url=https://gogogogogogo.vercel.app/telegram-webhook

var accessToken = "6847302523:AAGH73xiPrQhKaP5n76IItYpYGG0v_0DffoY"
var channelName = "test_usdv111111"

var commandCheckUserJoined = "check_user_joined"

type TelegramClient struct {
	Bot    *tgbotapi.BotAPI
	Update *tgbotapi.Update
}

func NewTelegramClient(r *gin.Context) (*TelegramClient, error) {
	bot, err := tgbotapi.NewBotAPI(accessToken)
	bot.Debug = true
	if err != nil {
		return nil, err
	}

	bytes, _ := io.ReadAll(r.Request.Body)
	var update tgbotapi.Update
	err = json.Unmarshal(bytes, &update)
	if err != nil {
		return nil, err
	}
	return &TelegramClient{
		Bot:    bot,
		Update: &update,
	}, nil
}

func (c *TelegramClient) HandleUpdate() {
	if c.Update.Message == nil && c.Update.CallbackQuery != nil {
		switch c.Update.CallbackQuery.Data {
		case commandCheckUserJoined:
			joined := c.CheckUserInChannel()
			if joined {
				c.SendMessage("Congratulation!!!")
			} else {
				c.SendMessage("You haven't join our channel, click join channel now !")
				c.SendWelcomeMessage()
			}
		default:
			c.SendMessage("I don't know what is that")
		}
	}

	if c.Update.Message != nil {
		if c.Update.Message.IsCommand() {

			switch c.Update.Message.Command() {
			case "start":
				c.SendWelcomeMessage()
			default:
				c.SendMessage("I don't know what is that")
			}
		} else {
			c.SendMessage("I don't know what is that, type /start to begin your journey")
		}
	}
}

func (c *TelegramClient) CheckUserInChannel() bool {
	userID := c.Update.CallbackQuery.From.ID

	chatConfigWithUser := tgbotapi.ChatConfigWithUser{
		SuperGroupUsername: "@" + channelName,
		UserID:             userID,
	}
	chatMember, err := c.Bot.GetChatMember(chatConfigWithUser)
	if err != nil {
		log.Println("Error getting chat member:", err)
		return false
	}

	// Check if the user's status is 'member', 'administrator' or 'creator'
	return chatMember.Status == "member" || chatMember.Status == "administrator" || chatMember.Status == "creator"
}

func (c *TelegramClient) SendWelcomeMessage() {
	msg := tgbotapi.NewMessage(c.GetChatID(), "Welcome to USDV !!!")
	row1 := []tgbotapi.InlineKeyboardButton{
		tgbotapi.NewInlineKeyboardButtonURL("Join Channel", "https://t.me/"+channelName),
		tgbotapi.NewInlineKeyboardButtonData("GO Next", commandCheckUserJoined),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(row1)
	c.Bot.Send(msg)
}

func (c *TelegramClient) SendMessage(message string) {
	msg := tgbotapi.NewMessage(c.GetChatID(), message)
	c.Bot.Send(msg)
}

func (c *TelegramClient) GetChatID() (chatID int64) {
	if c.Update.Message != nil {
		chatID = c.Update.Message.Chat.ID
	} else {
		chatID = c.Update.CallbackQuery.Message.Chat.ID
	}
	return
}

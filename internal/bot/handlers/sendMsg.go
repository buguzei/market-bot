package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func NewMsg(bot *tgbotapi.BotAPI, chatID int64, text string) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(chatID, text)

	msgConfig, err := bot.Send(msg)
	if err != nil {
		return tgbotapi.Message{}, err
	}
	return msgConfig, nil
}

func NewMsgAndMarkup(bot *tgbotapi.BotAPI, chatID int64, text string, kb tgbotapi.InlineKeyboardMarkup) (tgbotapi.Message, error) {
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = kb

	msgConfig, err := bot.Send(msg)
	if err != nil {
		return msgConfig, err
	}

	return msgConfig, nil
}

func NewEditMsgAndMarkup(bot *tgbotapi.BotAPI, chatID int64, msgID int, text string, kb tgbotapi.InlineKeyboardMarkup, needKb bool) error {
	editedMsg := tgbotapi.NewEditMessageTextAndMarkup(chatID, msgID, text, kb)

	if !needKb {
		editedMsg.ReplyMarkup = nil
	}

	_, err := bot.Send(editedMsg)
	if err != nil {
		return err
	}

	return nil
}

func NewDelMsg(bot *tgbotapi.BotAPI, chatID int64, msgID int) error {
	delMsg := tgbotapi.NewDeleteMessage(chatID, msgID)

	_, err := bot.Send(delMsg)
	if err != nil {
		return err
	}

	return nil
}

func NewPhoto(bot *tgbotapi.BotAPI, chatID int64, file tgbotapi.RequestFileData, caption string) error {
	newPh := tgbotapi.NewPhoto(chatID, file)
	newPh.Caption = caption

	_, err := bot.Send(newPh)
	if err != nil {
		return err
	}

	return nil
}

func NewPhotoAndMarkup(bot *tgbotapi.BotAPI, chatID int64, file tgbotapi.RequestFileData, caption string, kb tgbotapi.InlineKeyboardMarkup) error {
	newPh := tgbotapi.NewPhoto(chatID, file)
	newPh.Caption = caption
	newPh.ReplyMarkup = kb

	_, err := bot.Send(newPh)
	if err != nil {
		return err
	}

	return nil
}

package main

import (
	"strings"

	"github.com/yanzay/tbot/v2"
)

var (
	admin_id string = "741994015"
	bot             = tbot.New("5858460508:AAGAAWN638pWBuTBCi6vumUr3hQqo9ksBvI")
	client          = bot.Client()
)

func main() {
	bot.HandleMessage("/start", func(m *tbot.Message) {
		client.SendMessage(m.Chat.ID, "Assalomu alaykum "+m.Chat.FirstName)
	})

	bot.HandleMessage("", func(m *tbot.Message) {
		if m.Chat.ID == admin_id {

			if m.ReplyToMessage != nil {
				sendToClient(m)
			} else {
				client.SendMessage(admin_id, "Xabar aniq manziliga yetishi uchun kerakli xabarga javob (reply) sifatida yozing")
			}
		} else {
			sendToAdmin(m)
			if len(m.Text) > 0 {
				client.SendMessage(admin_id, m.Chat.ID+"\n"+m.Chat.FirstName+" "+m.Chat.LastName+"(@"+m.Chat.Username+")"+" dan xabar:\n\n"+m.Text)
			}
			if m.Sticker != nil {
				client.SendMessage(admin_id, m.Chat.ID+"\n"+m.Chat.FirstName+" "+m.Chat.LastName+"(@"+m.Chat.Username+")"+" sticker jo'natdi:\n\n"+m.Text)
				client.SendSticker(admin_id, m.Sticker.FileID)
			}
			if m.Audio != nil {
				m.Audio.Title = "salom"
				client.SendAudio(
					admin_id,
					m.Audio.FileID,
					tbot.OptCaption(m.Chat.ID+"\n"+m.Chat.FirstName+" "+m.Chat.LastName+"(@"+m.Chat.Username+")"+" audio jo'natdi"),
				)
			}
		}
	})

	bot.Start()
}

func sendToAdmin(m *tbot.Message) {
	id := admin_id
	user := m.Chat.ID + "\n" + m.Chat.FirstName + " " + m.Chat.LastName + "(@" + m.Chat.Username + ")"

	if len(m.Text) > 0 {
		client.SendMessage(id, user+" dan xabar:\n\n"+m.Text)
	}
	if m.Audio != nil {
		client.SendAudio(id, m.Audio.FileID, tbot.OptCaption(user+" dan audio xabar"))
	}
	if m.Video != nil {
		client.SendVideo(id, m.Video.FileID, tbot.OptCaption(user+" dan video xabar"))
	}
	if m.Photo != nil {
		client.SendMessage(id, user+" dan rasmlar:")
		for _, i := range m.Photo {
			client.SendPhoto(id, i.FileID)
		}
	}
	if m.Document != nil {
		client.SendDocument(id, m.Document.FileID, tbot.OptCaption(user+" dan hujjat"))
	}
	if m.Game != nil {
		client.SendGame(id, m.Game.Title, tbot.OptCaption(user+" o'yin jo'natdi"))
	}
	if m.Voice != nil {
		client.SendVoice(id, m.Voice.FileID, tbot.OptCaption(user+" dan ovozli xabar"))
	}
	if m.Venue != nil {
		client.SendVenue(id, m.Venue.Location.Latitude, m.Venue.Location.Longitude, m.Venue.Title, m.Venue.Address, tbot.OptCaption(user+" bino manzilini jo'natdi"))
	}
	if m.Poll != nil {
		opt := make([]string, 0)
		for _, i := range m.Poll.Options {
			opt = append(opt, i.Text)
		}
		client.SendPoll(id, m.Poll.Question, opt, tbot.OptCaption(user+" dan poll"))
	}
	if m.Location != nil {
		client.SendLocation(id, m.Location.Latitude, m.Location.Longitude, tbot.OptCaption(user+" manzil jo'natdi"))
	}
}

func sendToClient(m *tbot.Message) {
	id := getID(m.ReplyToMessage)
	if len(id) == 0 {
		client.SendMessage(admin_id, "Bu turdagi xabar uchun javob yubora olmayman (")
		return
	}
	if len(m.Text) > 0 {
		client.SendMessage(id, m.Text)
		return
	}
	if m.Audio != nil {
		client.SendAudio(id, m.Audio.FileID)
	}
	if m.Video != nil {
		client.SendVideo(id, m.Video.FileID)
	}
	if m.Photo != nil {
		for _, i := range m.Photo {
			client.SendPhoto(id, i.FileID)
		}
	}
	if m.Document != nil {
		client.SendDocument(id, m.Document.FileID)
	}
	if m.Game != nil {
		client.SendGame(id, m.Game.Title)
	}
	if m.Voice != nil {
		client.SendVoice(id, m.Voice.FileID)
	}
	if m.Venue != nil {
		client.SendVenue(id, m.Venue.Location.Latitude, m.Venue.Location.Longitude, m.Venue.Title, m.Venue.Address)
	}
	if m.Poll != nil {
		opt := make([]string, 0)
		for _, i := range m.Poll.Options {
			opt = append(opt, i.Text)
		}
		client.SendPoll(id, m.Poll.Question, opt)
	}
	if m.Location != nil {
		client.SendLocation(id, m.Location.Latitude, m.Location.Longitude)
	}
}

func getID(m *tbot.Message) string {
	if len(m.Text) > 0 {
		return strings.Split(m.Text, " ")[0]
	} else if len(m.Caption) > 0 {
		return strings.Split(m.Caption, " ")[0]
	}
	return ""
}
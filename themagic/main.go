package main

import (
	"context"

	"time"

	"github.com/rs/zerolog/log"
	tb "gopkg.in/tucnak/telebot.v2"
)

// THE AGENDA:
// 1. Interfaces
// 2. Usage
// 3. Inheritance
// 4. A bit about the internal structure
// 5. The interfaces in the telegram bot
// 6. JSON logging and JSON itself
// 7. Data format (JSON, PROTOBUF)

const (
	pollerPeriod = 10 * time.Second
)

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Err(err).Msg("can't load config")
		return
	}

	b, err := tb.NewBot(tb.Settings{
		Token:  config.APIKey,
		Poller: &tb.LongPoller{Timeout: pollerPeriod},
	})
	if err != nil {
		log.Err(err).Msg("can't create new bot")
		return
	}

	handler := newHandler(newProc())

	b.Handle(tb.OnText, func(m *tb.Message) {
		// all messages handler
	})

	b.Handle(helloHandlerName, func(m *tb.Message) {
		text, err := handler.handleHello(context.Background(), m)
		if err != nil {
			log.Print("error while handle hello:", err)
			return
		}
		b.Send(m.Sender, text)
	})

	b.Handle(catHandlerName, func(m *tb.Message) {
		photo, err := handler.handlePhoto(context.Background(), m, "cat")
		if err != nil {
			log.Print("error while handle photo:", err)
			return
		}

		if _, err = b.Send(m.Sender, photo); err != nil {
			log.Print(err)
		}
	})

	log.Print("cat bot started")
	b.Start()
}

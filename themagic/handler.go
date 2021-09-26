package main

import (
	"context"

	"github.com/rs/zerolog/log"
	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	helloHandlerName = "/hello"
	catHandlerName   = "/cat"
)

type (
	handler interface {
		handlePhoto(ctx context.Context, m *tb.Message, kind string) (tb.Sendable, error)
		handleHello(ctx context.Context, m *tb.Message) (string, error)
	}

	handlerImpl struct {
		proc proc
	}
)

func newHandler(proc proc) handler {
	return &handlerImpl{
		proc: proc,
	}
}

func (h *handlerImpl) handleHello(ctx context.Context, m *tb.Message) (string, error) {
	log.Info().Msg("handle hello request")
	// logging
	// metrics
	// validation
	// pre-processing (enrichment)
	return h.proc.replyHello(ctx, m)
}

func (h *handlerImpl) handlePhoto(ctx context.Context, m *tb.Message, kind string) (tb.Sendable, error) {
	log.Info().Msg("handle cat request")
	return h.proc.replyPhoto(ctx, kind)
}

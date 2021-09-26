package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"

	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	catDirPath     = "cats"
	greetingString = "hello, dear catman (^_^)"
)

type (
	proc interface {
		replyPhoto(ctx context.Context, kind string) (tb.Sendable, error)
		replyHello(ctx context.Context, m *tb.Message) (string, error)
	}

	procImpl struct {
	}
)

func newProc() proc {
	return &procImpl{}
}

func (p *procImpl) replyPhoto(ctx context.Context, kind string) (tb.Sendable, error) {
	files, err := ioutil.ReadDir(catDirPath)
	if err != nil {
		return nil, fmt.Errorf("error while read dir: %w", err)
	}

	file := files[rand.Intn(len(files))]
	return &tb.Photo{File: tb.FromDisk(fmt.Sprintf("%s/%s", catDirPath, file.Name()))}, nil
}

func (p *procImpl) replyHello(ctx context.Context, m *tb.Message) (string, error) {
	return greetingString, nil
}

package domain

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type HelloWorldPrinter interface {
	PrintHelloWorld(name string) error
}

type NameAsker interface {
	AskName() (string, error)
}

type App struct {
	printer HelloWorldPrinter
	asker   NameAsker
	log     *zap.Logger
}

func NewApp(printer HelloWorldPrinter, asker NameAsker, logger *zap.Logger) *App {
	return &App{
		printer: printer,
		asker:   asker,
		log:     logger,
	}
}

func (a *App) Start() error {
	name, err := a.asker.AskName()
	if err != nil {
		a.log.Error("ask name", zap.Error(err))
		return errors.Wrap(err, "failed to ask name")
	}

	err = a.printer.PrintHelloWorld(name)
	if err != nil {
		a.log.Error("print hello world", zap.Error(err))
		return errors.Wrap(err, "failed to print hello world")
	}

	return nil
}

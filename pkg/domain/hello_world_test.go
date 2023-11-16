package domain

import (
	mock_domain "clear-hello-world/pkg/mock"
	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"testing"
)

func TestApp_OK(t *testing.T) {
	ctrl := gomock.NewController(t)

	name := "some_name"

	asker := mock_domain.NewMockNameAsker(ctrl)
	printer := mock_domain.NewMockHelloWorldPrinter(ctrl)

	asker.EXPECT().AskName().Return(name, nil)
	printer.EXPECT().PrintHelloWorld(name).Return(nil)

	app := NewApp(printer, asker, zap.NewExample())

	err := app.Start()
	if err != nil {
		t.Fatal("failed start app", zap.Error(err))
	}
}

func TestApp_AskerError(t *testing.T) {
	ctrl := gomock.NewController(t)

	name := "some_name"
	someErr := errors.New("some error")

	asker := mock_domain.NewMockNameAsker(ctrl)
	printer := mock_domain.NewMockHelloWorldPrinter(ctrl)

	asker.EXPECT().AskName().Return(name, someErr)

	app := NewApp(printer, asker, zap.NewExample())

	err := app.Start()
	if !errors.Is(err, someErr) {
		t.Fatal("wrong error", zap.Error(err))
	}
}

func TestApp_PrinterError(t *testing.T) {
	ctrl := gomock.NewController(t)

	name := "some_name"
	someErr := errors.New("some error")

	asker := mock_domain.NewMockNameAsker(ctrl)
	printer := mock_domain.NewMockHelloWorldPrinter(ctrl)

	asker.EXPECT().AskName().Return(name, nil)
	printer.EXPECT().PrintHelloWorld(name).Return(someErr)

	app := NewApp(printer, asker, zap.NewExample())

	err := app.Start()
	if !errors.Is(err, someErr) {
		t.Fatal("wrong error", zap.Error(err))
	}
}

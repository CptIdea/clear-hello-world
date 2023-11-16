package fmt_asker

import (
	"clear-hello-world/pkg/domain"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type asker struct {
	log *zap.Logger
}

func New(log *zap.Logger) domain.NameAsker {
	return &asker{log: log}
}

func (a *asker) AskName() (string, error) {
	var name string

	fmt.Print("Enter your name: ")
	_, err := fmt.Scan(&name)
	if err != nil {
		a.log.Error("fmt.Scanf ask name", zap.Error(err))
		return "", errors.Wrap(err, "ask name")
	}

	return name, nil
}

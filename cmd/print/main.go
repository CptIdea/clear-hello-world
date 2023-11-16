package main

import (
	fmt_asker "clear-hello-world/pkg/askers/fmt"
	"clear-hello-world/pkg/domain"
	fmt_printer "clear-hello-world/pkg/printers/fmt"
	"go.uber.org/zap"
)

func main() {
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	asker := fmt_asker.New(log)

	printer := fmt_printer.New()

	app := domain.NewApp(printer, asker, log)
	err = app.Start()
	if err != nil {
		log.Fatal("failed start app", zap.Error(err))
	}
}

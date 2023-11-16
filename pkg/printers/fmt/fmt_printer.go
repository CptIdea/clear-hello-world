package fmt_printer

import (
	"clear-hello-world/pkg/domain"
	"fmt"
)

type printer struct {
}

func New() domain.HelloWorldPrinter {
	return &printer{}
}

func (p *printer) PrintHelloWorld(name string) error {
	fmt.Printf("Hello, %s!\n", name)

	return nil
}

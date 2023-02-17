package main

import (
	_ "embed"
	"log"

	"github.com/wailsapp/wails/v3/internal/parser/testdata/struct_literal_multiple_other/services"
	"github.com/wailsapp/wails/v3/pkg/application"
)

// GreetService is great
type GreetService struct {
	SomeVariable int
	lowerCase    string
}

// Greet does XYZ
func (*GreetService) Greet(name string) string {
	return "Hello " + name
}

func main() {
	app := application.New(application.Options{
		Bind: []interface{}{
			&GreetService{},
			&services.OtherService{},
		},
	})

	app.NewWebviewWindow()

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}

}

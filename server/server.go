package server

import (
	"fmt"
	// "log"
	"net/http"
	"strings"
)

type App struct {
	mux          *http.ServeMux
	handlerCount int
}

func NewApp() *App {
	return &App{
		mux:          http.NewServeMux(),
		handlerCount: 0,
	}
}

func (app *App) RunServer(port string) error {
	app.printBanner(port)
	server := &http.Server{
		Addr:    port,
		Handler: app.mux,
	}
	return server.ListenAndServe()
}

func (app *App) printBanner(port string) {
	urlBase := fmt.Sprintf("http://localhost%s", port)
	handlerCount := fmt.Sprintf("Handlers......: %d", app.handlerCount)

	fmt.Println("┌───────────────────────────────────────────────────┐")
	fmt.Printf("│%s│\n", textCenter("MyServer v1.0.0", 51))
	fmt.Printf("│%s│\n", textCenter(urlBase, 51))
	fmt.Printf("│%s│\n", strings.Repeat("", 51))
	fmt.Printf("│%s│\n", textCenter(handlerCount, 51))
	fmt.Println("└───────────────────────────────────────────────────┘")
}

func textCenter(text string, width int) string {
	if len(text) >= width {
		return text[:width]
	}

	padding := (width - len(text)) / 2

	return strings.Repeat(" ", padding) + text + strings.Repeat(" ", width-len(text)-padding)
}

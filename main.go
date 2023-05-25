package main

import (
	"context"
	"github.com/Buco7854/http-shutdown/errors"
	"github.com/Buco7854/http-shutdown/router"
	"github.com/Buco7854/http-shutdown/serializers"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"time"
)

const shutdownTimeout = 10 * time.Second

func main() {
	r := new(router.Router)

	r.AddRoute("GET", "/shutdown", func(writer http.ResponseWriter, request *http.Request) {
		if err := exec.Command("cmd", "/C", "shutdown", "/s", "/t", "10").Run(); err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				exitStatus := exitError.ExitCode()
				if exitStatus == 1190 {
					errors.JSONError(
						writer,
						"Sorry, a system shutdown is already scheduled. Please wait until the system shuts down automatically.",
						http.StatusConflict,
					)
					return
				}
			}
			errors.JSONError(
				writer,
				"An error has occurred while trying to initiate the shutdown procedure",
				http.StatusInternalServerError,
			)
			return
		}
		serializers.JsonResponse(
			writer,
			struct {
				Detail string `json:"detail"`
			}{
				Detail: "Shutdown procedure successfully initialized",
			},
		)
	})
	r.AddRoute("GET", "/launch-apex-legends", func(writer http.ResponseWriter, request *http.Request) {
		if err := exec.Command("cmd", "/C", "start", "steam://rungameid/1172470").Run(); err != nil {
			errors.JSONError(
				writer,
				"An error has occurred while trying to launch the game",
				http.StatusInternalServerError,
			)
			return
		}
		serializers.JsonResponse(
			writer,
			struct {
				Detail string `json:"detail"`
			}{
				Detail: "Successfully Launched Apex Legends",
			},
		)
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	log.Printf("Server listening on %v", server.Addr)

	go func() {
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ERROR on ListenAndServe(): %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Println("ERROR on Shutdown()", err)
	}
	log.Println("Server Shutdown Complete")
}

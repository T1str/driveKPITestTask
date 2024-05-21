package main

import (
	"net/http"
	"os"
	"runtime"

	"github.com/rs/zerolog"
	"testtask/internal/buffer"
	"testtask/internal/config"
	"testtask/internal/entities"
	"testtask/internal/handlers"
)

const bufferLimit = 1000

func main() {
	runtime.GOMAXPROCS(1)
	cfg := config.InitConfig()
	b := buffer.NewBuffer(*cfg)

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	defer logger.Info().Msg("application stopped")
	logger.Info().Msg("system initialization started")

	ch, done := make(chan entities.Request, bufferLimit), make(chan struct{})

	go func() {
		for {
			select {
			case req := <-ch:
				ok, err := b.HandleItem(req)
				if err != nil {
					logger.Error().Msgf("Error handling item: %v", err)
					continue
				}
				if !ok {
					if err = b.Retry(req); err != nil {
						logger.Error().Msgf("Error retrying to handle item: %v", err)
						break
					}
				}
			case <-done:
				return
			}
		}
	}()

	http.HandleFunc("/buffer", func(w http.ResponseWriter, r *http.Request) {
		handlers.BufferHandler(w, r, ch)
	})

	logger.Info().Msgf("Server is listening on port: %v", cfg.Port)
	err := http.ListenAndServe(":"+cfg.Port, nil)
	if err != nil {
		logger.Fatal().Msgf("Error starting server: %v", err)
	}
}

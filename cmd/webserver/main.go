package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"

	"lab-quade.de/serverManager/internal/config"
	"lab-quade.de/serverManager/internal/handler"
)

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	errorLogger, err := getErrorLogger(cfg)
	if err != nil {
		log.Fatal(err)
	}

	accessLogger, err := getAccessLogger(cfg)
	if err != nil {
		log.Fatal(err)
	}

	templateCache, err := buildTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	httpHandler := handler.Handler{
		ErrorLogger:   errorLogger,
		AccessLogger:  accessLogger,
		TemplateCache: templateCache,
		Config:        cfg,
	}

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.WebServer.Port),
		Handler:      httpHandler.Routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	errorLogger.Info(fmt.Sprintf("starting server on %s", server.Addr))
	err = server.ListenAndServe()
	errorLogger.Error(err.Error())
	os.Exit(1)
}

func getErrorLogger(cfg *config.Config) (*slog.Logger, error) {
	errorLogFile, err := os.OpenFile(cfg.Logging.ErrorLog, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("Fehler beim öffnen des Logfiles: %w", err)
	}

	opts := &slog.HandlerOptions{}
	switch strings.ToUpper(cfg.Logging.LogLevel) {
	case "DEBUG":
		opts.Level = slog.LevelDebug
	case "INFO":
		opts.Level = slog.LevelInfo
	case "WARN":
		opts.Level = slog.LevelWarn
	case "ERROR":
		opts.Level = slog.LevelError
	default:
		log.Printf("Ungültiges LogLevel: %s, setze auf INFO\n", cfg.Logging.LogLevel)
		opts.Level = slog.LevelInfo
	}

	opts.ReplaceAttr = func(groups []string, a slog.Attr) slog.Attr {
		if a.Key == slog.TimeKey {
			t := time.Now()
			a.Value = slog.StringValue(t.Format("2006-01-02 15:04:05"))
		}

		return a
	}

	jsonHandler := slog.NewJSONHandler(errorLogFile, opts)

	logger := slog.New(jsonHandler)

	return logger, nil
}

func getAccessLogger(cfg *config.Config) (*log.Logger, error) {
	var accessLogger *log.Logger = nil

	if cfg.Logging.EnableAccessLog {
		accessLogFile, err := os.OpenFile(cfg.Logging.AccessLog, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			return nil, fmt.Errorf("Fehler beim Öffnen des Access-Log-Files: %w", err)
		}
		accessLogger = log.New(accessLogFile, "", log.Ldate|log.Ltime)
	}

	return accessLogger, nil
}

package logger

import (
	"log/slog"
	"os"
)

var Log *slog.Logger

func Init() {
	
	file, err := os.OpenFile(
		"game.log",
		os.O_CREATE|os.O_APPEND|os.O_WRONLY,
		0644,
	)

	if err != nil {
		panic(err)
	}

	var handler slog.Handler

	handler = slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	Log = slog.New(handler)
}

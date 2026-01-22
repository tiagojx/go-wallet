package middleware

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

// inicializa o slog, direciona os logs JSON para o terminal, defini o slog como log padrão.
func InitLogger() *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)
	return logger
}

// o middleware para os logs em JSON.
func Logging(next http.Handler) http.Handler {
	// retorna uma função lambda para o http.Handler.
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// registra a hora em que a Request iniciou.
		start := time.Now()

		// serve o Handler especificado nos parâmetros.
		next.ServeHTTP(w, r)
		// saída do Handler...

		// calcula o tempo que a Request durou.
		duration := time.Since(start)

		// exibe o log.
		slog.Info("request have been proccessed",
			"method", r.Method,
			"path", r.URL.Path,
			"remote_addr", r.RemoteAddr,
			"duration_ms", duration.Microseconds(),
		)
	})
}

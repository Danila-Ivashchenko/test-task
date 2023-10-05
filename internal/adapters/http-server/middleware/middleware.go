package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

type middleware struct {
	logger *slog.Logger
}

func New(l *slog.Logger) *middleware {
	return &middleware{
		logger: l,
	}
}

func (m middleware) RequestID(c *gin.Context) {
	requestID := uuid.New().String()
	c.Set("RequestID", requestID)
	c.Writer.Header().Set("X-Request-ID", requestID)

	c.Next()
}

// func (m middleware) Logger(c *gin.Context) {
// 	log := m.logger.With(
// 		slog.String("component", "http-server"),
// 	)

// 	log.Info("logger middleware available")

// 	requestId, ok := c.Get("RequestID")

// 	entry := log.With(
// 		slog.String("method", c.Request.Method),
// 		slog.String("path", c.Request.URL.Path),
// 		slog.String("user_agent", c.Request.UserAgent()),
// 		slog.String("remate_addr", c.Request.RemoteAddr),
// 	)

// 	if ok {
// 		entry = entry.With(
// 			slog.String("request_id", requestId.(string)),
// 		)
// 	}

// 	t := time.Now()

// 	defer func() {
// 		entry.Info(
// 			"request complited",
// 			slog.Int("status", c.Request.Response.StatusCode),
// 			slog.String("duration", time.Since(t).String()),
// 		)
// 	}()

// 	c.Next()

// }

package event

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func NewConnection() string {
	_ = godotenv.Load()
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/",
		os.Getenv("MQ_USER"),
		os.Getenv("MQ_PASSWORD"),
		os.Getenv("MQ_HOST"),
		os.Getenv("MQ_PORT_GO"),
	)

	return connStr
}

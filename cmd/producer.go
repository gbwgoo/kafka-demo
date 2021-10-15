package cmd

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
)

var produceCmd = &cobra.Command{
	Use:   "produce",
	Short: "to produce messages",
	Long:  "to produce messages",
	Run: func(cmd *cobra.Command, args []string) {
		topic := "demo"
		partition := 0

		conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
		if err != nil {
			log.Fatal("failed to dial leader:", err)
		}

		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		_, err = conn.WriteMessages(
			kafka.Message{Value: []byte("one!")},
			kafka.Message{Value: []byte("two!")},
			kafka.Message{Value: []byte("three!")},
		)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}

		if err := conn.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}
	},
}

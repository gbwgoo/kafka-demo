package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
)

var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "to consume messages",
	Long:  "to consume messages",
	Run: func(cmd *cobra.Command, args []string) {
		topic := "demo"
		partition := 0

		conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
		if err != nil {
			log.Fatal("failed to dial leader:", err)
		}

		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

		b := make([]byte, 10e3) // 10KB max per message
		for {
			_, err := batch.Read(b)
			if err != nil {
				break
			}
			fmt.Println(string(b))
		}

		if err := batch.Close(); err != nil {
			log.Fatal("failed to close batch:", err)
		}

		if err := conn.Close(); err != nil {
			log.Fatal("failed to close connection:", err)
		}
	},
}

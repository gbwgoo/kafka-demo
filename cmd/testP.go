package cmd

import (
	"context"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
)

var str string

var testPCmd = &cobra.Command{
	Use:   "testP",
	Short: "test produce messages",
	Long:  "test produce messages",
	Run: func(cmd *cobra.Command, args []string) {
		topic := "demo"
		partition := 0

		conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
		if err != nil {
			log.Fatal("failed to dial leader:", err)
		}

		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
		_, err = conn.WriteMessages(
			kafka.Message{Value: []byte(str)},
		)
		if err != nil {
			log.Fatal("failed to write messages:", err)
		}

		if err := conn.Close(); err != nil {
			log.Fatal("failed to close writer:", err)
		}
	},
}

func init() {
	testPCmd.Flags().StringVarP(&str, "str", "s", "", "请输入要发送的内容")
}

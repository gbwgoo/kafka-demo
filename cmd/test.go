package cmd

// import (
// 	"context"
// 	"log"

// 	"github.com/segmentio/kafka-go"
// 	"github.com/spf13/cobra"
// )

// var testCCmd = &cobra.Command{
// 	Use:   "testC",
// 	Short: "test consume messages",
// 	Long:  "test consume messages",
// 	Run: func(cmd *cobra.Command, args []string) {
// 		topic := "demo"
// 		partition := 0

// 		conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
// 		if err != nil {
// 			log.Fatal("failed to dial leader:", err)
// 		}

// 		// conn.SetReadDeadline(time.Now().Add(5 * time.Second))

// 		// first, last, err := conn.ReadOffsets()
// 		// ps, err := conn.ReadPartitions("test", "none")
// 		// conn.DeleteTopics("test")
// 		// conn.CreateTopics(kafka.TopicConfig{
// 		// 	Topic:             "test",
// 		// 	NumPartitions:     1,
// 		// 	ReplicationFactor: 1,
// 		// })
// 		// conn.Seek(1, kafka.SeekStart)
// 		conn.Seek(2, kafka.SeekEnd)

// 		for {
// 			m, err := conn.ReadMessage(10e3)

// 			if err != nil {
// 				log.Println("read err:", err)
// 				break
// 			}

// 			log.Println("m:", string(m.Value))
// 		}

// 		if err := conn.Close(); err != nil {
// 			log.Fatal("failed to close connection:", err)
// 		}
// 	},
// }

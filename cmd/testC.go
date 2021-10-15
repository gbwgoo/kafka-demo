package cmd

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/spf13/cobra"
)

var testCCmd = &cobra.Command{
	Use:   "testC",
	Short: "test consume messages",
	Long:  "test consume messages",
	Run: func(cmd *cobra.Command, args []string) {
		topic := "demo"
		partition := 0

		conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
		if err != nil {
			log.Fatal("failed to dial leader:", err)
		}

		// conn.SetReadDeadline(time.Now().Add(5 * time.Second))

		// first, last, err := conn.ReadOffsets()
		// ps, err := conn.ReadPartitions("test", "none")
		// conn.DeleteTopics("test")
		// conn.CreateTopics(kafka.TopicConfig{
		// 	Topic:             "test",
		// 	NumPartitions:     1,
		// 	ReplicationFactor: 1,
		// })

		// conn.Seek(5, kafka.SeekEnd) //4条msg
		// offset, whence := conn.Offset()
		// log.Println("offset:", offset) //0
		// log.Println("whence:", whence) //0
		//output: 4条msg

		// conn.Seek(4, kafka.SeekEnd) //4条msg
		// offset, whence := conn.Offset()
		// log.Println("offset:", offset) //0
		// log.Println("whence:", whence) //1
		//output: 3条msg

		// conn.Seek(3, kafka.SeekEnd) //4条msg
		// offset, whence := conn.Offset()
		// log.Println("offset:", offset) //1
		// log.Println("whence:", whence) //1
		//output: 3条msg

		// conn.Seek(2, kafka.SeekEnd) //4条msg
		// offset, whence := conn.Offset()
		// log.Println("offset:", offset) //2
		// log.Println("whence:", whence) //1
		//output: 2条msg

		// conn.Seek(1, kafka.SeekEnd) //4条msg
		// offset, whence := conn.Offset()
		// log.Println("offset:", offset) //3
		// log.Println("whence:", whence) //1
		//output: 1条msg

		// conn.Seek(0, kafka.SeekEnd) //4条msg
		// offset, whence := conn.Offset()
		// log.Println("offset:", offset) //4
		// log.Println("whence:", whence) //1
		// output: 0条msg

		conn.Seek(4, 1)
		for {
			m, err := conn.ReadMessage(10e3) //这个消息大小要改到viper里
			if err != nil {
				log.Println("read err:", err)
				break
			}

			log.Println("m:", string(m.Value))
			offset, whence := conn.Offset()
			log.Println("after offset:", offset)
			log.Println("after whence:", whence)
		}

		if err := conn.Close(); err != nil {
			log.Fatal("failed to close connection:", err)
		}
	},
}

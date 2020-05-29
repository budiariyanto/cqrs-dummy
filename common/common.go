package common

import (
	kafka "github.com/segmentio/kafka-go"
	"github.com/jmoiron/sqlx"
	"time"
	"fmt"
)

func NewKafkaWriter() *kafka.Writer {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:32772"},
		Topic:   "donation-create",
		Balancer: &kafka.LeastBytes{},
	})

	return w
}

func NewKafkaReader(group, topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:32772"},
		// GroupID:   "group2",
		// Topic:     "donation-create",
		GroupID:   group,
		Topic:     topic,
		ReadBackoffMin: 50*time.Millisecond,
		ReadBackoffMax: 500*time.Millisecond,
		MinBytes:  10e1, // 10KB
		MaxBytes:  10e6, // 10MB
	})
}


func NewDB(dbName string) (db *sqlx.DB){
	connString := fmt.Sprintf("root:12345@tcp(localhost:32771)/%s?parseTime=true", dbName)
	db, err := sqlx.Connect("mysql", connString)
	if err != nil {
		panic(err)
	}

	return
}
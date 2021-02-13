package main

import (
	"github.com/streadway/amqp"
	"github.com/joho/godotenv"
	"log"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf(msg)
	}
}

func main() {
	err := godotenv.Load()
	failOnError(err, "Error loading .env file")
	rabbitMQUrl := os.Getenv("RABBITMQ_URL")

	conn, err := amqp.Dial(RABBITMQ_URL)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to create channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		"test",
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to consume queue")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("MESSAGE: %s", d.Body)
			d.Ack(true)
		}
	}()

	log.Printf("[*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

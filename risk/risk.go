package main

import (
	"fmt"
	"log"

	"github.com/jailtonjunior94/bank/risk/infra/environments"
	"github.com/jailtonjunior94/bank/risk/infra/ioc"
)

func main() {
	environments.NewSettings()

	ioc.SetupDependencyInjection()

	channel := ioc.RabbitMQ.GetChannel()
	defer channel.Close()

	channel.QueueDeclare(environments.QueueLoanRisk, true, false, false, false, nil)
	messages, err := channel.Consume(environments.QueueLoanRisk, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to our RabbitMQ instance")
	fmt.Println("Waiting for messages")

	forever := make(chan bool)

	go func() {
		for message := range messages {
			fmt.Printf("Recieved Message: %s\n", message.Body)
			ioc.RiskHandle.AnalyzeRisk(message.Body)
		}
	}()

	<-forever
}

package environments

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	MongoConnectionString = ""
	LoanDatabase          = ""
	LoanCollection        = ""
	CustomerBaseURL       = ""
	CustomerRoute         = ""
	RabbitMQConnection    = ""
	QueueLoanRisk         = ""
)

func NewSettings() {
	var err error

	viper.SetConfigName(fmt.Sprintf("settings.%s", os.Getenv("ENVIRONMENT")))
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}

	MongoConnectionString = viper.GetString("mongo.connectionString")
	LoanDatabase = viper.GetString("mongo.database")
	LoanCollection = viper.GetString("mongo.loanCollection")
	CustomerBaseURL = viper.GetString("customerApi.baseURL")
	CustomerRoute = viper.GetString("customerApi.customers")
	RabbitMQConnection = viper.GetString("rabbitMQ.connection")
	QueueLoanRisk = viper.GetString("rabbitMQ.queueLoanRisk")
}

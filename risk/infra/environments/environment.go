package environments

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	RabbitMQConnection = ""
	QueueLoanRisk      = ""
	LoanBaseURL        = ""
	LoanRoute          = ""
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

	RabbitMQConnection = viper.GetString("rabbitMQ.connection")
	QueueLoanRisk = viper.GetString("rabbitMQ.queueLoanRisk")
	LoanBaseURL = viper.GetString("loanApi.baseURL")
	LoanRoute = viper.GetString("loanApi.updateLoan")
}

package environments

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

var (
	MongoConnectionString = ""
	CustomerDatabase      = ""
	CustomerCollection    = ""
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
	CustomerDatabase = viper.GetString("mongo.database")
	CustomerCollection = viper.GetString("mongo.customersCollection")
}

package repositories

import (
	"context"
	"errors"
	"log"

	"github.com/jailtonjunior94/bank/customer/business/entities"
	"github.com/jailtonjunior94/bank/customer/business/interfaces"
	"github.com/jailtonjunior94/bank/customer/infra/database"
	"github.com/jailtonjunior94/bank/customer/infra/environments"

	"go.mongodb.org/mongo-driver/bson"
)

type CustomerRepository struct {
	Client database.IMongoConnection
}

func NewCustomerRepository(client database.IMongoConnection) interfaces.ICustomerRepository {
	return &CustomerRepository{Client: client}
}

func (r *CustomerRepository) Add(c *entities.Customer) (customer *entities.Customer, err error) {
	collection, err := r.Client.GetCollection(environments.CustomerDatabase, environments.CustomerCollection)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Ocorreu um erro inesperado!")
	}

	_, err = collection.InsertOne(context.Background(), &c)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Ocorreu um erro inesperado!")
	}

	return c, nil
}

func (r *CustomerRepository) GetByDocument(document string) (customer *entities.Customer, err error) {
	collection, err := r.Client.GetCollection(environments.CustomerDatabase, environments.CustomerCollection)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Ocorreu um erro inesperado!")
	}

	if err := collection.FindOne(context.Background(), bson.M{"document": document}).Decode(&customer); err != nil {
		log.Println(err)
		return nil, nil
	}

	return customer, nil
}

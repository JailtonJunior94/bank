package repositories

import (
	"context"
	"errors"
	"log"

	"github.com/jailtonjunior94/bank/loan/business/entities"
	"github.com/jailtonjunior94/bank/loan/business/interfaces"
	"github.com/jailtonjunior94/bank/loan/infra/database"
	"github.com/jailtonjunior94/bank/loan/infra/environments"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoanRepository struct {
	Client database.IMongoConnection
}

func NewLoanRepository(client database.IMongoConnection) interfaces.ILoanRepository {
	return &LoanRepository{Client: client}
}

func (r *LoanRepository) GetById(id string) (loan *entities.Loan, err error) {
	collection, err := r.Client.GetCollection(environments.LoanDatabase, environments.LoanCollection)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Ocorreu um erro inesperado!")
	}

	if err := collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&loan); err != nil {
		log.Println(err)
		return nil, nil
	}

	return loan, nil
}

func (r *LoanRepository) GetByDocument(document string) (loans []entities.Loan, err error) {
	collection, err := r.Client.GetCollection(environments.LoanDatabase, environments.LoanCollection)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Ocorreu um erro inesperado!")
	}

	cursor, err := collection.Find(context.Background(), bson.M{"customer.document": document})
	if err != nil {
		log.Println(err)
		return nil, errors.New("Ocorreu um erro inesperado!")
	}

	if err = cursor.All(context.Background(), &loans); err != nil {
		log.Println(err)
		return nil, errors.New("Ocorreu um erro inesperado!")
	}

	return loans, nil
}

func (r *LoanRepository) Add(l *entities.Loan) (loan *entities.Loan, err error) {
	collection, err := r.Client.GetCollection(environments.LoanDatabase, environments.LoanCollection)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Ocorreu um erro inesperado!")
	}

	_, err = collection.InsertOne(context.Background(), &l)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Ocorreu um erro inesperado!")
	}

	return l, nil
}

func (r *LoanRepository) Update(l *entities.Loan) (loan *entities.Loan, err error) {
	collection, err := r.Client.GetCollection(environments.LoanDatabase, environments.LoanCollection)
	if err != nil {
		log.Println(err)
		return nil, errors.New("Ocorreu um erro inesperado!")
	}

	res, err := collection.UpdateOne(context.Background(), bson.M{"_id": l.ID}, bson.M{"$set": l})
	if err != nil {
		log.Println(err)
		return nil, errors.New("Ocorreu um erro inesperado!")
	}

	if res.ModifiedCount == 0 {
		log.Println(err)
		return nil, errors.New("Ocorreu um erro inesperado!")
	}

	return l, nil
}

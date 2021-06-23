package entities

type Customer struct {
	Document string `bson:"document"`
	Name     string `bson:"name"`
}

func NewCustomer(document, name string) *Customer {
	return &Customer{
		Document: document,
		Name:     name,
	}
}

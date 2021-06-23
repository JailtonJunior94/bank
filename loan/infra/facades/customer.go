package facades

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jailtonjunior94/bank/loan/business/dtos"
	"github.com/jailtonjunior94/bank/loan/business/interfaces"
	"github.com/jailtonjunior94/bank/loan/infra/environments"
)

type CustomerFacade struct {
}

func NewCustomerFacade() interfaces.ICustomerFacade {
	return &CustomerFacade{}
}

func (f *CustomerFacade) CustomerByDocument(document string) (customer *dtos.CustomerResponse, err error) {
	res, err := http.Get(fmt.Sprintf(environments.CustomerBaseURL+environments.CustomerRoute, document))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&customer); err != nil {
		log.Println(err)
		return nil, err
	}

	return customer, nil
}

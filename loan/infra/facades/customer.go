package facades

import (
	"encoding/json"
	"fmt"
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
	url := fmt.Sprintf(environments.CustomerBaseURL+environments.CustomerRoute, document)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&customer); err != nil {
		return nil, err
	}

	return customer, nil
}

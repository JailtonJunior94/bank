package facades

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jailtonjunior94/bank/risk/business/dtos"
	"github.com/jailtonjunior94/bank/risk/business/interfaces"
	"github.com/jailtonjunior94/bank/risk/infra/environments"
)

type LoanFacade struct {
}

func NewLoanFacade() interfaces.ILoanFacade {
	return &LoanFacade{}
}

func (f *LoanFacade) UpdateLoan(id string, r *dtos.Evaluation) error {
	client := &http.Client{}

	json, err := json.Marshal(&r)
	if err != nil {
		return err
	}

	url := fmt.Sprintf(environments.LoanBaseURL+environments.LoanRoute, id)
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(json))
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	fmt.Println(resp.StatusCode)
	return nil
}

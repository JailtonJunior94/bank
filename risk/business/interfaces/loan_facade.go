package interfaces

type ILoanFacade interface {
	GetById(id string) (loan interface{}, err error)
}

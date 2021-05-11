package alphavantage

type Usecase interface {
	GetSymbol(stockParam string) (interface{}, error)
}

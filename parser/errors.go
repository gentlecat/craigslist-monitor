package parser

type NoPriceError struct{}

func (m *NoPriceError) Error() string {
	return "Unable to find the price"
}

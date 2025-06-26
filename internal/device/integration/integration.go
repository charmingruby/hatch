package integration

// BankService is an example contract for a third party bank service.
type BankService interface {
	Call() error
}

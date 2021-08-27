package hawk

// Sender is used to perform network interaction between Catcher and Hawk.
type Sender interface {
	// Send sends error reports to Hawk.
	Send([]ErrorReport) error
}

type Transport interface {
}

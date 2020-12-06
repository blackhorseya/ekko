package health

type impl struct {
}

// NewImpl is a constructor of implement business
func NewImpl() Biz {
	return &impl{}
}

// Readiness to handle application has been ready
func (i *impl) Readiness() (ok bool, err error) {
	return true, nil
}

// Liveness to handle application was alive
func (i *impl) Liveness() (ok bool, err error) {
	return true, nil
}

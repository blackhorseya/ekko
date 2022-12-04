package genx

// Generator declare a genx factory
//go:generate mockery --all --inpackage
type Generator interface {
	Int64() int64
}

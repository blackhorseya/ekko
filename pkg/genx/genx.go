//go:generate mockgen -destination=./mock_genx.go -package=genx -source=genx.go

package genx

// Generator declare a genx factory
type Generator interface {
	Int64() int64
}

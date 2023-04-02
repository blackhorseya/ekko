//go:generate mockgen -destination=./mock_${GOFILE} -package=genx -source=${GOFILE}

package genx

// Generator declare a genx factory
type Generator interface {
	Int64() int64
}

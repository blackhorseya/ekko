//go:generate mockgen -destination=./mock_${GOFILE} -package=${GOPACKAGE} -source=${GOFILE}

package fuglex

import (
	"github.com/blackhorseya/ekko/pkg/contextx"
)

// Dialer is an interface that represents the dialer.
type Dialer interface {
	// IntradayQuote is used to get the intraday quote.
	IntradayQuote(ctx contextx.Contextx, symbol string) (res *IntradayQuote, err error)
}

package grid2

import (
	"github.com/OvictorVieira/promeheux.api/pkg/fixedpoint"
	"github.com/OvictorVieira/promeheux.api/pkg/types"
)

// collectTradeFee collects the fee from the given trade slice
func collectTradeFee(trades []types.Trade) map[string]fixedpoint.Value {
	fees := make(map[string]fixedpoint.Value)
	for _, t := range trades {
		if t.FeeDiscounted {
			continue
		}

		if fee, ok := fees[t.FeeCurrency]; ok {
			fees[t.FeeCurrency] = fee.Add(t.Fee)
		} else {
			fees[t.FeeCurrency] = t.Fee
		}
	}
	return fees
}

func aggregateTradesQuantity(trades []types.Trade) fixedpoint.Value {
	tq := fixedpoint.Zero
	for _, t := range trades {
		tq = tq.Add(t.Quantity)
	}
	return tq
}

// aggregateTradesQuoteQuantity aggregates the quote quantity from the given trade slice
func aggregateTradesQuoteQuantity(trades []types.Trade) fixedpoint.Value {
	quoteQuantity := fixedpoint.Zero
	for _, t := range trades {
		if t.QuoteQuantity.IsZero() {
			quoteQuantity = quoteQuantity.Add(t.Price.Mul(t.Quantity))
		} else {
			quoteQuantity = quoteQuantity.Add(t.QuoteQuantity)
		}
	}

	return quoteQuantity
}

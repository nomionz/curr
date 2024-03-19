package ecb

import (
	"github.com/nomionz/currency-converter/pkg/backoff"
	"github.com/shopspring/decimal"
	"log/slog"
	"time"
)

type Client struct {
	fetchTimer *time.Timer
	backoff    backoff.Backoff
	st         Store
	log        *slog.Logger
}

func NewClient(st Store) *Client {
	return &Client{
		st:         st,
		log:        slog.With("sys", "ecb_provider"),
		fetchTimer: DefaultTicker(),
		backoff:    backoff.NewBackoff(),
	}
}

func (c *Client) Convert(fromS, toS, amountS string) (decimal.Decimal, error) {
	var empty decimal.Decimal

	amount, err := decimal.NewFromString(amountS)

	if err != nil {
		return empty, err
	}

	from, err := c.st.Get(fromS)

	if err != nil {
		return empty, err
	}

	to, err := c.st.Get(toS)

	if err != nil {
		return empty, err
	}

	return amount.Div(from).Mul(to).Round(4), nil
}

func (c *Client) GetCurrencies() []string {
	return c.st.GetKeys()
}

func (c *Client) Fetch() {
	tf := c.fetchTimer
	tb := time.NewTimer(0)

	for {
		select {
		case <-tf.C:
		case <-tb.C:
		}

		if err := c.tryFetch(); err != nil {
			c.log.Warn("failed to fetch ecb data", "err", err)
			tb.Reset(c.backoff.Next())
			tf = nil
		} else {
			tb.Stop()
			tf = c.fetchTimer
		}
	}
}

func (c *Client) tryFetch() error {
	ecb, err := FetchDaily()

	if err != nil {
		return err
	}

	for _, d := range ecb.Data {
		for _, r := range d.Rates {
			dec, err := decimal.NewFromString(r.Rate)

			if err != nil {
				return err
			}

			c.st.Put(r.Currency, dec)
		}
	}

	return nil
}

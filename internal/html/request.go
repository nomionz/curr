package html

const (
	AmountTo   = "amountTo"
	AmountFrom = "amountFrom"
)

type Request struct {
	From       string `form:"from"`
	To         string `form:"to"`
	AmountTo   string `form:"amountTo"`
	AmountFrom string `form:"amountFrom"`
}

func (r *Request) GetAmount() string {
	if r.AmountTo != "" && r.AmountFrom == "" {
		return r.AmountTo
	}

	return r.AmountFrom
}

func (r *Request) GetAmountTagName() string {
	if r.AmountFrom == "" && r.AmountTo == "" {
		return ""
	}

	if r.AmountTo != "" && r.AmountFrom == "" {
		return AmountTo
	}

	return AmountFrom
}

package html

type ViewData struct {
	InputFrom  inputData
	InputTo    inputData
	SelectFrom selectData
	SelectTo   selectData
}

func NewViewData() *ViewData {
	return &ViewData{
		InputFrom: inputData{
			Id:          "from",
			Name:        AmountFrom,
			Placeholder: "Type amount to convert from",
			HXTarget:    "#to",
		},
		InputTo: inputData{
			Id:          "to",
			Name:        AmountTo,
			Placeholder: "Type amount to convert to",
			HXTarget:    "#from",
		},
		SelectFrom: selectData{
			Name:     "from",
			HXTarget: "#to",
		},
		SelectTo: selectData{
			Name:     "to",
			HXTarget: "#to",
		},
	}
}

func (v *ViewData) ResetValues() {
	v.InputFrom.Value = ""
	v.InputTo.Value = ""
}

func (v *ViewData) SetCurrencies(currencies []string) {
	v.SelectFrom.Currencies = currencies
	v.SelectTo.Currencies = currencies
}

type inputData struct {
	Id          string
	Name        string
	Value       string
	Placeholder string
	HXTarget    string
}

type selectData struct {
	Name       string
	HXInclude  string
	HXTarget   string
	Currencies []string
}

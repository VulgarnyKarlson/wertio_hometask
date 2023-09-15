package domain

type Ticker struct {
	id       int
	name     string
	price    float64
	currency string
}

func NewTicker(id int, name, to string, price float64) *Ticker {
	return &Ticker{
		id:       id,
		name:     name,
		price:    price,
		currency: to,
	}
}

func (t *Ticker) ID() int {
	return t.id
}

func (t *Ticker) Name() string {
	return t.name
}

func (t *Ticker) Price() float64 {
	return t.price
}

func (t *Ticker) Currency() string {
	return t.currency
}

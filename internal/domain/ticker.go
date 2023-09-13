package domain

type Ticker struct {
	id    int
	name  string
	price float64
}

func NewTicker(id int, name string, price float64) *Ticker {
	return &Ticker{
		id:    id,
		name:  name,
		price: price,
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

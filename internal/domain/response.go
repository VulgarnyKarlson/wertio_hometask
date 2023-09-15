package domain

type Response struct {
	amount float64
	result float64
	err    error
}

func NewResponse(amount, result float64, err error) *Response {
	return &Response{amount: amount, result: result, err: err}
}

func (r *Response) Amount() float64 {
	return r.amount
}

func (r *Response) Result() float64 {
	return r.result
}

func (r *Response) Err() error {
	return r.err
}

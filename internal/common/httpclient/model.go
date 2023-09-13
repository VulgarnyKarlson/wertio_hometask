package httpclient

type Request struct {
	Method string
	Body   []byte
	Path   string
	Query  map[string]string
}

type Response struct {
	StatusCode int
	Body       []byte
}

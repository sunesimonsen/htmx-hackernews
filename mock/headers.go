package mock

type Headers map[string]string

func (headers Headers) Get(name string) string {
	return headers[name]
}

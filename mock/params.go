package mock

type Params map[string]string

func (params Params) Get(name string) string {
	return params[name]
}

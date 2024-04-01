package request

type Method string

const (
	GET    = "GET"
	PUT    = "PUT"
	POST   = "POST"
	DELETE = "DELETE"
)

func CreateMethod(s string) (Method, error) {
	m := Method(s)
	if !m.IsValid() {
		return "", errInvalidMethod
	}

	return m, nil
}

func (m Method) String() string {
	return string(m)
}

func (m Method) IsValid() bool {
	switch m {
	case GET, PUT, POST, DELETE:
		return true
	default:
		return false
	}
}

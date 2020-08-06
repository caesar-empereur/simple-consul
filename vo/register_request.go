package vo

type RegisterRequest struct {
	Id string `json:"Id"`

	Name string `json:"Name"`

	Address string `json:"Address"`

	Port int `json:"Port"`

	Tag []string `json:"Tag"`

	Check RegisterCheck `json:"Check"`
}

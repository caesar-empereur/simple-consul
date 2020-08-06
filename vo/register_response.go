package vo

type RegisterResponse struct {
	Kind string `json:"Kind"`

	Id string `json:"Id"`

	Service string `json:"Service"`

	Tag []string `json:"Tag"`

	Meta interface{} `json:"Meta"`

	Address string `json:"Address"`

	Port int `json:"Port"`

	EnableTagOverride bool `json:"EnableTagOverride"`

	CreateIndex int `json:"CreateIndex"`

	ModifyIndex int `json:"ModifyIndex"`

	ProxyDestination string `json:"ProxyDestination"`

	Connect interface{} `json:"Connect"`
}

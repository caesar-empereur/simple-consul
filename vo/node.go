package vo

type Node struct {
	ID string `json:"ID"`

	Node string `json:"Node"`

	Address string `json:"Address"`

	Datacenter string `json:"Datacenter"`

	TaggedAddresses interface{} `json:"TaggedAddresses"`

	Meta interface{} `json:"Meta"`

	CreateIndex int `json:"CreateIndex"`

	ModifyIndex int `json:"ModifyIndex"`
}

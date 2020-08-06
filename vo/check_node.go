package vo

type CheckNode struct {
	Node string `json:"Node"`

	CheckID string `json:"CheckID"`

	Name string `json:"Name"`

	Status string `json:"Status"`

	Notes string `json:"Notes"`

	Output string `json:"Output"`

	ServiceID string `json:"ServiceID"`

	ServiceName string `json:"ServiceName"`

	ServiceTags []string `json:"ServiceTags"`

	Definition string `json:"Definition"`

	CreateIndex int `json:"CreateIndex"`

	ModifyIndex int `json:"ModifyIndex"`
}

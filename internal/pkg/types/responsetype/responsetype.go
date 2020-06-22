package responsetype

type (
	Base struct {
		Result interface{} `json:"result,omitempty"`
		Page   interface{} `json:"page,omitempty"`
		ID     interface{} `json:"id,omitempty"`
	}
)

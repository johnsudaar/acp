package params

type Params map[string]Input

type Input struct {
	Type        Type           `json:"type,omitempty"`
	Required    bool           `json:"required,omitempty"`
	Default     interface{}    `json:"default,omitempty"`
	Description string         `json:"description,omitempty"`
	Min         int            `json:"min,omitempty"`
	Max         int            `json:"max,omitempty"`
	Options     []SelectOption `json:"options,omitempty"`
}

type Type string

const (
	IP     Type = "ip"
	Number Type = "number"
	String Type = "string"
	Select Type = "select"
)

type SelectOption struct {
	Value string `json:"value,omitempty"`
	Name  string `json:"name,omitempty"`
}

package params

type Params map[string]Input

type Input struct {
	Type        Type        `json:"type"`
	Required    bool        `json:"required"`
	Default     interface{} `json:"default,omitempty"`
	Description string      `json:"description,omitempty"`
	Min         int         `json:"min,omitempty"`
	Max         int         `json:"max,omitempty"`
}

type Type string

const (
	IP     Type = "ip"
	Number Type = "number"
	String Type = "string"
)

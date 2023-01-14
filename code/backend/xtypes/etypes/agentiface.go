package etypes

type AgentIface struct {
	Name        string                `json:"name,omitempty"`
	Definations map[string]any        `json:"definations,omitempty"`
	Methods     map[string]*Method    `json:"methods,omitempty"`
	Events      map[string]*EventType `json:"events,omitempty"`
	Schemas     map[string]*ValueType `json:"schemas,omitempty"`
}

type Method struct {
	Info       string            `json:"info,omitempty"`
	Arg        ValueType         `json:"arg,omitempty"`
	ReturnType ValueType         `json:"return_type,omitempty"`
	ErrorTypes map[string]string `json:"error_types,omitempty"` // <error_id, error_info>
}

type ValueType struct {
	Type         string      `json:"type,omitempty"`
	Values       []ValueType `json:"values,omitempty"`
	PropertyName string      `json:"property,omitempty"` // only applicable to object
	Ref          string      `json:"ref,omitempty"`      // (type, valye, prop_name) ||  ref to schema key
}

type EventType struct {
	Info       string     `json:"info,omitempty"`
	Async      bool       `json:"async,omitempty"`
	ArgData    ValueType  `json:"arg_data,omitempty"`
	ReturnData *ValueType `json:"return_data,omitempty"`
}

/*

# primitive types
	- null
	- bool
	- integer
	- string
	- any
# complex type
	- array
	- object

# definations example
	{
		"init_cms_engine": {
			"sub_routes": {},
			"global_routes": {},
			"widget_hooks": {},
		}
	}
*/

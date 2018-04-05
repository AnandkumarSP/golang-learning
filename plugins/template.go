package plugins

// Plugin interface with necessary methods
type Plugin interface {
	Execute() (*PluginReturnValue, error)
}

type PluginReturnValue struct {
	Values []interface{}
	Err    error
}

type BasePlugin struct{}

func (base *BasePlugin) Execute() (*PluginReturnValue, error) {
	r := &PluginReturnValue{}
	return r, nil
}

func (base *BasePlugin) ConvertToInterface(a ...interface{}) (b []interface{}) {
	b = make([]interface{}, len(a))
	for i := range a {
		b[i] = a[i]
	}
	return
}

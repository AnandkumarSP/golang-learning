package multiply

import (
	"net/http"

	template ".."
	"../../utils"
)

type Multiply struct {
	template.BasePlugin
	val1 int
	val2 int
}

func (f *Multiply) Execute() (r template.PluginReturnValue, err error) {
	r = template.PluginReturnValue{}
	result := &r
	result.Values = f.ConvertToInterface(f.val1 * f.val2)
	return
}

// New creates a multiply instance and returns the pointer
func New(args ...interface{}) (r *Multiply, err error) {
	defer utils.GenericErrorHandler(nil)

	arg0, ok := args[0].(int)
	if !ok {
		return nil, utils.NewCustomError("Expected 'int' as first argument in 'Multiply' step", http.StatusBadRequest)
	}
	arg1, ok := args[1].(int)
	if !ok {
		return nil, utils.NewCustomError("Expected 'int' as second argument in 'Multiply' step", http.StatusBadRequest)
	}

	r = &Multiply{
		val1: arg0,
		val2: arg1,
	}
	return
}

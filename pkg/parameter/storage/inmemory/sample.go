package inmemory

import "github.com/AlbertMorenoDEV/simple-sidecar/pkg/parameter"

// Parameters contain a bunch of items
var Parameters = map[string]*parameter.Parameter{
	"parameter_1": &parameter.Parameter{
		ID:    "parameter_1",
		Value: "value_1",
	},
	"parameter_2": &parameter.Parameter{
		ID:    "parameter_2",
		Value: "value_2",
	},
	"parameter_3": &parameter.Parameter{
		ID:    "parameter_3",
		Value: "value_3",
	},
	"parameter_4": &parameter.Parameter{
		ID:    "parameter_4",
		Value: "value_4",
	},
}

package inmemory

import (
	"fmt"
	"sync"

	"github.com/AlbertMorenoDEV/simple-sidecar/pkg/parameter"
)

type parameterRepository struct {
	mtx        sync.RWMutex
	parameters map[string]*parameter.Parameter
}

// NewParameterRepository creates a new parameter respository instance
func NewParameterRepository(parameters map[string]*parameter.Parameter) parameter.Repository {
	if parameters == nil {
		parameters = make(map[string]*parameter.Parameter)
	}

	return &parameterRepository{
		parameters: parameters,
	}
}

func (r *parameterRepository) FetchParameters() ([]*parameter.Parameter, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	values := make([]*parameter.Parameter, 0, len(r.parameters))
	for _, value := range r.parameters {
		values = append(values, value)
	}
	return values, nil
}

func (r *parameterRepository) DeleteParameter(ID string) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	delete(r.parameters, ID)

	return nil
}

func (r *parameterRepository) UpdateParameter(param *parameter.Parameter) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.parameters[param.ID] = param
	return nil
}

func (r *parameterRepository) FetchParameterByID(ID string) (*parameter.Parameter, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	for _, v := range r.parameters {
		if v.ID == ID {
			return v, nil
		}
	}

	return nil, fmt.Errorf("The ID %s doesn't exist", ID)
}

package parameter

// Parameter defines the properties of a parameter
type Parameter struct {
	ID    string `json:"ID"`
	Value string `json:"value"`
}

// Repository provides access to the parameter storage
type Repository interface {
	// FetchParameters return all parameters saved in storage
	FetchParameters() ([]*Parameter, error)

	// DeleteParameter remove parameter with given ID
	DeleteParameter(ID string) error

	// UpdateParameter modify parameter with given ID and given new data
	UpdateParameter(param *Parameter) error

	// FetchParameterByID returns the parameter with given ID
	FetchParameterByID(ID string) (*Parameter, error)
}

package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/AlbertMorenoDEV/simple-sidecar/pkg/parameter"
	"github.com/AlbertMorenoDEV/simple-sidecar/pkg/parameter/storage/inmemory"
)

func TestHealth(t *testing.T) {
	parameters := map[string]*parameter.Parameter{}
	repo := inmemory.NewParameterRepository(parameters)
	s := New(repo)

	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	s.Router().ServeHTTP(rr, req)

	assertStatusCode(t, rr, http.StatusOK)
	assertEqualsJSON(t, rr.Body.String(), `{"ok":true}`)
}

func TestFetchParameters(t *testing.T) {
	parameters := map[string]*parameter.Parameter{
		"parameter_1": &parameter.Parameter{
			ID:    "parameter_1",
			Value: "value_1",
		},
		"parameter_2": &parameter.Parameter{
			ID:    "parameter_2",
			Value: "value_2",
		},
	}
	expected := `[{"ID":"parameter_1","value":"value_1"},{"ID":"parameter_2","value":"value_2"}]`

	repo := inmemory.NewParameterRepository(parameters)
	s := New(repo)

	req, err := http.NewRequest(http.MethodGet, "/parameters", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-Session-Token", "00000000")

	rr := httptest.NewRecorder()
	s.Router().ServeHTTP(rr, req)

	assertStatusCode(t, rr, http.StatusOK)
	assertEqualsJSON(t, rr.Body.String(), expected)
}

func TestFetchParameter(t *testing.T) {
	parameters := map[string]*parameter.Parameter{
		"parameter_1": &parameter.Parameter{
			ID:    "parameter_1",
			Value: "value_1",
		},
		"parameter_2": &parameter.Parameter{
			ID:    "parameter_2",
			Value: "value_2",
		},
	}
	expected := `{"ID":"parameter_1","value":"value_1"}`

	repo := inmemory.NewParameterRepository(parameters)
	s := New(repo)

	req, err := http.NewRequest(http.MethodGet, "/parameters/parameter_1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("X-Session-Token", "00000000")

	rr := httptest.NewRecorder()
	s.Router().ServeHTTP(rr, req)

	assertStatusCode(t, rr, http.StatusOK)
	assertEqualsJSON(t, rr.Body.String(), expected)
}

func TestCreateParameter(t *testing.T) {
	eParam := &parameter.Parameter{
		ID:    "parameter_1",
		Value: "value_1",
	}
	nParam := &parameter.Parameter{
		ID:    "test_1",
		Value: "1111",
	}
	params := map[string]*parameter.Parameter{
		"parameter_1": eParam,
	}
	bReq, _ := json.Marshal(nParam)

	repo := inmemory.NewParameterRepository(params)
	s := New(repo)

	req, err := http.NewRequest(http.MethodPut, "/parameters/test_1", bytes.NewBuffer([]byte(bReq)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Session-Token", "00000000")

	rr := httptest.NewRecorder()
	s.Router().ServeHTTP(rr, req)

	assertStatusCode(t, rr, http.StatusAccepted)
	assertParametersCount(t, params, 2)
	assertParametersExist(t, params, *eParam)
	assertParametersExist(t, params, *nParam)
}

func TestUpdateParameter(t *testing.T) {
	eParam := &parameter.Parameter{
		ID:    "parameter_1",
		Value: "value_1",
	}
	nParam := &parameter.Parameter{
		ID:    "parameter_1",
		Value: "updated_value_1",
	}
	params := map[string]*parameter.Parameter{
		"parameter_1": eParam,
	}
	bReq, _ := json.Marshal(nParam)

	repo := inmemory.NewParameterRepository(params)
	s := New(repo)

	req, err := http.NewRequest(http.MethodPut, "/parameters/test_1", bytes.NewBuffer([]byte(bReq)))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Session-Token", "00000000")

	rr := httptest.NewRecorder()
	s.Router().ServeHTTP(rr, req)

	assertStatusCode(t, rr, http.StatusAccepted)
	assertParametersCount(t, params, 1)
	assertParametersExist(t, params, *nParam)
}

func TestDeleteParameter(t *testing.T) {
	eParam := &parameter.Parameter{
		ID:    "parameter_1",
		Value: "value_1",
	}
	params := map[string]*parameter.Parameter{
		"parameter_1": eParam,
	}

	repo := inmemory.NewParameterRepository(params)
	s := New(repo)

	req, err := http.NewRequest(http.MethodDelete, "/parameters/parameter_1", nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Session-Token", "00000000")

	rr := httptest.NewRecorder()
	s.Router().ServeHTTP(rr, req)

	assertStatusCode(t, rr, http.StatusAccepted)
	assertParametersCount(t, params, 0)
}

func assertStatusCode(t *testing.T, rr *httptest.ResponseRecorder, exp int) {
	if status := rr.Code; status != exp {
		t.Errorf("handler returned wrong status code: got %v want %v", status, exp)
	}
}

func assertEqualsJSON(t *testing.T, exp, rec string) {
	var oExp interface{}
	var oRec interface{}
	var err error

	err = json.Unmarshal([]byte(exp), &oExp)
	if err != nil {
		t.Fatal(fmt.Errorf("Error mashalling expected string :: %s", err.Error()))
	}

	err = json.Unmarshal([]byte(rec), &oRec)
	if err != nil {
		t.Fatal(fmt.Errorf("Error mashalling received string :: %s", err.Error()))
	}

	if !reflect.DeepEqual(oExp, oRec) {
		t.Errorf("handler returned unexpected body: got %v want %v", exp, rec)
	}
}

func assertParametersCount(t *testing.T, params map[string]*parameter.Parameter, c int) {
	if len(params) != c {
		t.Errorf("wrong number of parameters: have %v want %v", len(params), c)
	}
}

func assertParametersExist(t *testing.T, params map[string]*parameter.Parameter, e parameter.Parameter) {
	var f *parameter.Parameter

	for i := range params {
		if params[i].ID == e.ID {
			f = params[i]
			break
		}
	}

	if f == nil {
		t.Errorf("parameter does not exist: want %v", e)
	}

	if f.Value != e.Value {
		t.Errorf("wrong parameter value: got %v want %v", f.Value, e.Value)
	}
}

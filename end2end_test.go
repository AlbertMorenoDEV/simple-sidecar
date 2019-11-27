package main_test

import (
	"net/http"
	"testing"

	"gopkg.in/h2non/baloo.v3"
)

var test = baloo.New("http://127.0.0.1:7983")

func TestHealth(t *testing.T) {
	test.Get("/health").
		Expect(t).
		Status(http.StatusOK).
		Type("json").
		JSON(map[string]bool{"ok": true}).
		Done()
}

func TestParamsGetEmpty(t *testing.T) {
	test.Get("/parameters").
		SetHeader("X-Session-Token", "00000000").
		Expect(t).
		Status(http.StatusOK).
		Type("json").
		JSON([]string{}).
		Done()
}

func TestParamsPost(t *testing.T) {
	test.Post("/parameters").
		SetHeader("X-Session-Token", "00000000").
		JSON(map[string]string{"ID": "parameter_1", "Value": "new value"}).
		Expect(t).
		Status(http.StatusAccepted).
		Done()
}

func TestParamsGetNewValue(t *testing.T) {
	test.Get("/parameters").
		SetHeader("X-Session-Token", "00000000").
		Expect(t).
		Status(http.StatusOK).
		Type("json").
		JSON(`[{"ID":"parameter_1","value":"new value"}]`).
		Done()
}

func TestParamsPut(t *testing.T) {
	test.Put("/parameters/parameter_1").
		SetHeader("X-Session-Token", "00000000").
		JSON(map[string]string{"ID": "parameter_1", "Value": "new value edited"}).
		Expect(t).
		Status(http.StatusAccepted).
		Done()
}

func TestParamsGetUpdatedValue(t *testing.T) {
	test.Get("/parameters").
		SetHeader("X-Session-Token", "00000000").
		Expect(t).
		Status(http.StatusOK).
		Type("json").
		JSON(`[{"ID":"parameter_1","value":"new value edited"}]`).
		Done()
}

func TestParamsDelete(t *testing.T) {
	test.Delete("/parameters/parameter_1").
		SetHeader("X-Session-Token", "00000000").
		Expect(t).
		Status(http.StatusAccepted).
		Done()
}

// AssertFunc(func(res *http.Response, req *http.Request) error {
// 	fmt.Println(res.Header.Get("Content-Type"))
// 	fmt.Println(res.Status)
// 	fmt.Println(res.Body)
// 	return nil
// }).

package main

import (
	"encoding/json"
	"reflect"
	"testing"
)

func checkJSONStringsEqual(item1, item2 string) (bool, error) {
	var out1, out2 interface{}

	err := json.Unmarshal([]byte(item1), &out1)
	if err != nil {
		return false, nil
	}

	err = json.Unmarshal([]byte(item2), &out2)
	if err != nil {
		return false, nil
	}

	return reflect.DeepEqual(out1, out2), nil
}

func TestTransform(t *testing.T) {
	err := setupKzm()
	if err != nil {
		t.Error("error setting up kazaam")
	}

	in := `{
		"timings": {
			"start": "2019-06-13T20:37:43.160Z"
		}
	}`

	expected := `{
		"startDate": "2019-06-13T20:37:43.160Z"
	}`

	result, _ := transform(in)

	res, _ := checkJSONStringsEqual(result, expected)
	if !res {
		t.Errorf("something bad\n %v\n is not as we expected\n %v\n", result, expected)
	}
}

package prom

import (
	prom "immich-exp/prometheus"
	"reflect"
	"testing"

	API "immich-exp/api"
)

func TestGetName(t *testing.T) {
	result2 := &API.StructAllUsers{
		{ID: "1", Name: "John", Email: "john@example.com", IsAdmin: true},
		{ID: "2", Name: "Jane", Email: "jane@example.com", IsAdmin: false},
	}

	result := "1"
	expected := API.StructCustomUser{
		ID:      "1",
		Name:    "John",
		Email:   "john@example.com",
		IsAdmin: true,
	}
	actual := prom.GetName(result, result2)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, but got: %v", expected, actual)
	}

	result = "3"
	expected = API.StructCustomUser{}
	actual = prom.GetName(result, result2)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, but got: %v", expected, actual)
	}
}

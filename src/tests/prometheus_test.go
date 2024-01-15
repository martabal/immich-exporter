package prom

import (
	"immich-exp/models"
	prom "immich-exp/prometheus"
	"reflect"
	"testing"
)

func TestGetName(t *testing.T) {
	result2 := &models.StructAllUsers{
		{ID: "1", Name: "John", Email: "john@example.com", IsAdmin: true},
		{ID: "2", Name: "Jane", Email: "jane@example.com", IsAdmin: false},
	}

	result := "1"
	expected := models.StructCustomUser{
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
	expected = models.StructCustomUser{}
	actual = prom.GetName(result, result2)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: %v, but got: %v", expected, actual)
	}
}

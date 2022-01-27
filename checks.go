package testo

import (
	"reflect"
	"testing"
)

func EqualValues(t *testing.T, v1, v2 interface{}) {
	if v1 != v2 {
		t.Errorf("expected equality: %v, %v", v1, v2)
	}
}

func UnequalValues(t *testing.T, v1, v2 interface{}) {
	if v1 == v2 {
		t.Errorf("expected unequalilty: %v, %v", v1, v2)
	}
}

func Nil(t *testing.T, v interface{}, nilExpected bool) {
	val := reflect.ValueOf(v)
	if val.IsNil() && !nilExpected {
		t.Error("unexpected nil")
	}
	if !val.IsNil() && nilExpected {
		t.Errorf("missing expected nil: %v", v)
	}
}

func Error(t *testing.T, err error, errorExpected bool) {
	if err != nil && !errorExpected {
		t.Error("unexpected error:", err)
	}
	if err == nil && errorExpected {
		t.Error("missing expected error")
	}
}

func DeepEqual(t *testing.T, r1, r2 interface{}) {
	if !reflect.DeepEqual(r1, r2) {
		t.Errorf("expected equality: %v, %v", r1, r2)
	}
}

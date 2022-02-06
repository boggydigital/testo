package testo

import (
	"reflect"
	"testing"
)

const (
	Less           = -2
	LessOrEqual    = -1
	Equal          = 0
	GreaterOrEqual = 1
	Greater        = 2
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
	if v == nil && !nilExpected {
		t.Error("unexpected nil")
	}
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

func CompareInt64(t *testing.T, i1, i2 int64, exp int) {
	if i1 < i2 {
		if exp != Less && exp != LessOrEqual {
			t.Errorf("expected %d < %d", i1, i2)
		}
	} else if i1 == i2 {
		if exp != LessOrEqual && exp != Equal && exp != GreaterOrEqual {
			t.Errorf("expected %d = %d", i1, i2)
		}
	} else {
		if exp != Greater && exp != GreaterOrEqual {
			t.Errorf("expected %d > %d", i1, i2)
		}
	}
}

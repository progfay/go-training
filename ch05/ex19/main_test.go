package main_test

import (
	"reflect"
	"testing"

	ex19 "github.com/progfay/go-training/ch05/ex19"
)

func Test_ReturnNonZeroValueWithoutReturnStatement(t *testing.T) {
	out := ex19.ReturnNonZeroValueWithoutReturnStatement()
	if out == reflect.Zero(reflect.TypeOf(out)).Int() {
		t.Errorf("%d is zero value", out)
	}
}

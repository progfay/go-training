package pattern

import "testing"

type T struct {
	PostCode string `pattern:"post-code"`
	Email    string `pattern:"email"`
}

func Test_Pack_Pass(t *testing.T) {
	in := T{
		PostCode: "012-3456",
		Email:    "example@gmail.com",
	}

	err := Validate(in)
	if err != nil {
		t.Error(err)
	}
}

func Test_Pack_Fail(t *testing.T) {
	in := T{
		PostCode: "hoge",
		Email:    "fuga",
	}

	err := Validate(in)
	if err == nil {
		t.Error("no validation error")
	}
}

package browser

import "testing"

func TestReg(t *testing.T) {
	a := Reg("https://exhentai.org/g/")
	if a == "" {
		t.Fail()
	}
}

func TestGetGIDAndToken(t *testing.T) {
	GetGIDAndToken("1761748/882d6987a6")
}

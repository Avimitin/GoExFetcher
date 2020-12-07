package browser

import "testing"

func TestReg(t *testing.T) {
	a := Reg("https://exhentai.org/g/1761748/882d6987a6")
	if a == "" {
		t.Errorf("Expect 1761748/882d6987a6 but got nothing")
		t.FailNow()
	}
	if a != "1761748/882d6987a6" {
		t.Errorf("Expect 1761748/882d6987a6 but got %s", a)
		t.Failed()
	}
}

func TestGetGIDAndToken(t *testing.T) {
	result := GetGIDAndToken("1761748/882d6987a6")
	if result == nil || result[0] != "1761748" || result[1] != "882d6987a6" {
		t.Errorf("Expect result[0] == 1761748, result[1] == 882d6987a6, but got %v", result)
	}
}

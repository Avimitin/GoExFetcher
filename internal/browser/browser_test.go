package browser

import (
	"testing"
)

func TestJsonBrowse(t *testing.T) {
	result, err := Browser("1761748", "882d6987a6")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	if result == nil {
		t.Errorf("Fetch unexpected nil struct")
		t.FailNow()
	}
	if len(result.Gmd) == 0 {
		t.Errorf("Can't fetch result")
	}
}

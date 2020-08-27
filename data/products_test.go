package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Abhishek ",
		Price: 4.5,
		SKU:   "abs-abs-abs",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}

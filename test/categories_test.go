package test

import (
	"github.com/CeoFred/goingecko"
	"testing"
)

func TestCategoriesList(t *testing.T) {

		cgClient := goingecko.NewClient(nil,"CG-QSYXVe6kc41uvas2kyrEr5Lz")

	categoriesList, err := cgClient.CategoriesList()
	if categoriesList == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

func TestCategories(t *testing.T) {

		cgClient := goingecko.NewClient(nil,"")

	categoriesList, err := cgClient.Categories("market_cap_desc")
	if categoriesList == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

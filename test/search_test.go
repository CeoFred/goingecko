package test

import (
	"testing"

	"github.com/CeoFred/goingecko"
)

func TestSearch(t *testing.T) {
	cgClient := goingecko.NewClient(nil, "")
	data, err := cgClient.Search("bitcoin")
	if data == nil {
		t.Errorf("Error")
	}
	if err != nil {
		t.Errorf("Error: %s", err)
	}
}

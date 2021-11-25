package maps_test

import (
	"github.com/matryer/is"
	"testing"
	"web-scraper/pkgs/maps"
)

func TestPinAddress(t *testing.T) {
	t.Skip()
	t.Run("given and address in the form of a string, then I expect a pin on a google map", func(t *testing.T) {
		is := is.New(t)
		address := "1600 Amphitheatre Parkway, Mountain View, CA"

		err := maps.PinAddress(address)

		is.NoErr(err)
	})
}

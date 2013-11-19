package govat

import (
	"strings"
	"testing"
)

const (
	CC  = Country("GB")
	UID = "117223643"
)

func TestCheckId(t *testing.T) {
	ok := Country("US").MustChargeVAT()
	if ok {
		t.Fatal()
	}
	ok = CC.MustChargeVAT()
	if !ok {
		t.Fatal()
	}
	res, err := CC.CheckId(UID)
	if err != nil {
		t.Fatal(err)
	}
	if x := res.CountryCode; x != "GB" {
		t.Fatal(x)
	}
	if x := res.VatNumber; x != "117223643" {
		t.Fatal(x)
	}
	if x := res.RequestDate; x == "" {
		t.Fatal()
	}
	if x := res.Valid; x != true {
		t.Fatal(x)
	}
	if x := res.Name; !strings.HasPrefix(x, "APPLE DISTRIBUTION INTERNATIONAL") {
		t.Fatal(x)
	}
	if x := res.Address; !strings.HasPrefix(x, "2 FURZEGROUND WAY") {
		t.Fatal(x)
	}
}

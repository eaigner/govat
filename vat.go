package govat

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

var invalidIdErr = errors.New("invalid VAT id")

var countries = map[Country]string{
	"AT": "Austria",
	"BE": "Belgium",
	"BG": "Bulgaria",
	"HR": "Croatia",
	"CY": "Cyprus",
	"CZ": "Czech Republic",
	"DK": "Denmark",
	"EE": "Estonia",
	"FI": "Finland",
	"FR": "France",
	"DE": "Germany",
	"EL": "Greece",
	"HU": "Hungary",
	"IE": "Ireland",
	"IT": "Italy",
	"LV": "Latvia",
	"LT": "Lithuania",
	"LU": "Luxembourg",
	"MT": "Malta",
	"NL": "Netherlands",
	"PL": "Poland",
	"PT": "Portugal",
	"RO": "Romania",
	"SK": "Slovakia",
	"SI": "Slovenia",
	"ES": "Spain",
	"SE": "Sweden",
	"GB": "United Kingdom and Isle of Man",
}

type Country string

func (c Country) Code() string {
	return strings.ToUpper(string(c))
}

func (c Country) Name() string {
	return countries[Country(c.Code())]
}

func (c Country) MustChargeVAT() bool {
	return c.Name() != ""
}

func (c Country) CheckId(vatId string) (*Result, error) {
	if len(vatId) < 2 {
		return nil, invalidIdErr
	}
	body := fmt.Sprintf(soapReqFmt, c, vatId)
	resp, err := http.Post(port, "text/xml", strings.NewReader(body))
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var v soapEnv
	err = xml.Unmarshal(b, &v)
	if err != nil {
		return nil, err
	}

	return &v.Body.CheckVatResponse, nil
}

type soapEnv struct {
	XMLName string   `xml:"Envelope"`
	Body    soapBody `xml:"Body"`
}

type soapBody struct {
	CheckVatResponse Result `xml:"checkVatResponse"`
}

type Result struct {
	CountryCode string `xml:"countryCode"`
	VatNumber   string `xml:"vatNumber"`
	RequestDate string `xml:"requestDate"`
	Valid       bool   `xml:"valid"`
	Name        string `xml:"name"`
	Address     string `xml:"address"`
}

const (
	port       = `http://ec.europa.eu/taxation_customs/vies/services/checkVatService`
	soapReqFmt = `<?xml version="1.0" encoding="UTF-8"?>
<SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns1="urn:ec.europa.eu:taxud:vies:services:checkVat:types">
  <SOAP-ENV:Body>
    <ns1:checkVat>
      <ns1:countryCode>%s</ns1:countryCode>
      <ns1:vatNumber>%s</ns1:vatNumber>
    </ns1:checkVat>
  </SOAP-ENV:Body>
</SOAP-ENV:Envelope>`
)

govat
=====

Validate EU VAT numbers in Go

```
if c := govat.Country("GB"); c.MustChargeVAT() {
  res, err := c.CheckId("VAT-ID-TO-CHECK")
  if err != nil {
    return false, err
  }
  return res.Valid, nil
}
```
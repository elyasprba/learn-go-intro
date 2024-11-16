package basic

import "fmt"

type Address struct {
	City, Province, Country string
}

func Method() {
	var address = Address{
		City:     "Jakarta",
		Province: "DKI Jakarta",
		Country:  "Indonesia",
	}

	fmt.Println(address.GetCity())
	fmt.Println(address.GetProvince())
	fmt.Println(address.GetCountry())
}

func (a *Address) GetCity() string {
	return "City : " + a.City
}

func (a *Address) GetProvince() string {
	return "Province : " + a.Province
}

func (a *Address) GetCountry() string {
	return "Country : " + a.Country
}

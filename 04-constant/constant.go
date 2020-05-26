package main

import "fmt"

// Brand is a mobile company type
type Brand string

// Constant for companies
const (
	APPLE   Brand = "Apple"
	GOOGLE  Brand = "Google"
	SAMSUNG Brand = "Samsung"
)

// GetBrand returns brand names
func GetBrand(brand Brand) {
	fmt.Println(brand)
}

func main() {
	GetBrand(APPLE)
	GetBrand(GOOGLE)
	GetBrand(SAMSUNG)
}

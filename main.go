package main

import (
	"fmt"
	"log"
)

func find(item string, slice []string) bool {
	for _, sliceItem := range slice {
		if item == sliceItem {
			return true
		} // retorna verdadero si la cadena se encuentra en el slice
	} // fin forange
	return false
}

type Refrigerador []string

func (r Refrigerador) Open() {
	fmt.Println(" Opening  refrigerator")
}

func (r Refrigerador) Close() {
	fmt.Println(" Closing refrigerator")
}

func (r Refrigerador) FindFood(food string) error {
	r.Open()
	defer r.Close()
	if find(food, r) {
		fmt.Println("Found ", food)
	} else {
		return fmt.Errorf(" %s not found ", food)
	}
	//defer r.Close()
	return nil
}

func main() {
	refrigerador := Refrigerador{"leche", "pizza", "salsa"}

	for _, food := range []string{"leche", "Bananas"} {
		error := refrigerador.FindFood(food)
		if error != nil {
			log.Fatal(error)
		}
	}

}

package main

import (
	"CoffeeShop/domain"
	"fmt"
)

func main() {
	coffees, err := domain.GetCoffees()

	if err != nil {
		fmt.Println("Error while getting coffeelist ", err)
		return
	}

	fmt.Println("Printing Coffees")

	for _, v := range coffees.CoffeeList {
		result := fmt.Sprintf("%s for $%v", v.Name, v.Price)
		fmt.Println(result)
	}
}

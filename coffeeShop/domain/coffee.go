package domain

import (
	"fmt"

	"github.com/spf13/viper"
)

type CoffeeDetails struct {
	Name  string
	Price float32
}

type CoffeeList struct {
	CoffeeList []CoffeeDetails
}

var Coffees CoffeeList

func GetCoffees() (*CoffeeList, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("fatal error loading config file: %w", err))
	}

	err = viper.Unmarshal(&Coffees)

	if err != nil {
		panic(fmt.Errorf("fatal error parse config file: %w", err))
	}

	return &Coffees, nil
}

func IsCoffeeAvailable(coffee string) string {
	for _, c := range Coffees.CoffeeList {
		if c.Name == coffee {
			result := fmt.Sprintf("%s for $%v", c.Name, c.Price)
			return result
		}
	}
	return ""
}

package main

import (
	"flag"
	"fmt"
	"os"
)

// CLI commands:
// go build .
// ./gocli get --all
// ./gocli get --zip 94158
// ./gocli add --name "Spices" --zip 94118 --url "https://www.ubereats.com/store/spices/lszSibfTSlu-uoUiB1g2aA" --dish "Dry Hot Pot" --mp 43.99
// open eats.json

func main() {
	fmt.Println("Celia's Go CLI App for UberEats!")

	// create am initial JSON slice
	myeats := []Eat{
		{
			Name:     "Philz Coffee",
			Zipcode:  94158,
			Url:      "https://www.ubereats.com/store/philz-coffee-civic-center/GKkC3kmlQlW-QwaBGWiGrw",
			Dish:     "Blueberry Muffin",
			MaxPrice: 21.00,
		},
		{
			Name:     "Shake Shack",
			Zipcode:  94103,
			Url:      "https://www.ubereats.com/store/shake-shack-845-market-street/jj_Fs4JhRZqTgj7m5hThPQ",
			Dish:     "ShackBurger",
			MaxPrice: 13.19,
		},
	}
	// save JSON file
	EncodeJSON(myeats)

	// 'get' subcommands
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)
	getAll := getCmd.Bool("all", false, "Get all eats")
	getZip := getCmd.Int("zip", 0, "Get ubereats by zipcode")

	// 'add' subcommands
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addName := addCmd.String("name", "", "Add ubereats' name")
	addZip := addCmd.Int("zip", 0, "Add ubereats' zipcode")
	addUrl := addCmd.String("url", "", "Add ubereats' UberEats url")
	addDish := addCmd.String("dish", "", "Add ubereats' famous dishes")
	addMaxPrice := addCmd.Float64("mp", 0.00, "Add ubereats' maximum price")

	// Validate 'get' or 'add' commands
	fmt.Println(os.Args[0:])
	if len(os.Args) < 2 {
		fmt.Println("expecting 'get' or 'add' commands")
		os.Exit(1)
	}

	// Switch statements to toggle commands
	switch os.Args[1] {
	case "get":
		// if get cmd, handle get
		HandleGet(getCmd, getAll, getZip)
	case "add":
		// if add cmd, handle add
		HandleAdd(addCmd, addName, addZip, addUrl, addDish, addMaxPrice)
	default:
		// if none of these
	}

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

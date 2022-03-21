package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const JsonFileName = "eats.json"

type Eat struct {
	Name     string  `json:"name"`
	Zipcode  int     `json:"zipcode"`
	Url      string  `json:"url"`
	Dish     string  `json:"dish"`
	MaxPrice float64 `json:"maxprice"`
}

func EncodeJSON(eats []Eat) {
	// convert slices to bytes
	var bytes []byte
	bytes, err := json.Marshal(eats)
	CheckError(err)
	fmt.Println(string(bytes))

	// save bytes to json files
	err = ioutil.WriteFile("./"+JsonFileName, bytes, 0644) //filemode:linux permission - owner read/write
	CheckError(err)
}

func DecodeJSON() (eats []Eat) {
	// read JSON data into bytes
	bytes, err := ioutil.ReadFile("./" + JsonFileName)
	CheckError(err)

	// transform bytes into a slice
	err = json.Unmarshal(bytes, &eats)
	CheckError(err)

	// fmt.Print(eats)
	return eats
}

func PrintColumnName() {
	fmt.Printf("Name \t Zipcode \t Url \t Dish \t MaxPrice \n")
}

func PrintJSONValue(eat Eat) {
	fmt.Printf("%v \t %v \t %v \t %v \t %v \n",
		eat.Name, eat.Zipcode, eat.Url, eat.Dish, eat.MaxPrice)
}

func HandleGet(getCmd *flag.FlagSet, all *bool, zip *int) {
	// parse arguments into app
	getCmd.Parse(os.Args[2:])

	// input Validation
	if *all == false && *zip == 0 {
		fmt.Print("use --all for all eats or specify zipcode")
		getCmd.PrintDefaults()
		os.Exit(1)
	} else if *all {
		// return all eats
		eats := DecodeJSON()
		PrintColumnName()
		for _, eat := range eats {
			PrintJSONValue(eat)
		}
		// stop app execution
		return
	} else if *zip != 0 {
		eats := DecodeJSON()
		zip := *zip
		for _, eat := range eats {
			if zip == eat.Zipcode {
				PrintColumnName()
				PrintJSONValue(eat)
			}
		}
	}
}

func ValidateAddEats(addCmd *flag.FlagSet, name *string, zip *int, url *string, dish *string, mp *float64) {
	// parse args into app
	addCmd.Parse(os.Args[2:])
	// fmt.Println(*name)
	// fmt.Println(*zip)
	// fmt.Println(*url)
	// fmt.Println(*dish)
	// fmt.Println(*mp)
	if *name == "" || *zip == 0 || *url == "" || *dish == "" || *mp == 0.00 {
		fmt.Print("Please enter data for all fields.")
		addCmd.PrintDefaults()
		os.Exit(1)
	}
}

func HandleAdd(addCmd *flag.FlagSet, name *string, zip *int, url *string, dish *string, mp *float64) {
	// validate input
	ValidateAddEats(addCmd, name, zip, url, dish, mp)

	// create a new struct
	eat := Eat{
		Name:     *name,
		Zipcode:  *zip,
		Url:      *url,
		Dish:     *dish,
		MaxPrice: *mp,
	}

	// append to original arrays
	eats := DecodeJSON()
	eats = append(eats, eat)

	// save JSON file
	EncodeJSON(eats)
}

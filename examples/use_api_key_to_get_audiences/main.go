package main

import (
	"fmt"
	"os"
)

func main() {
	apiKey := os.Getenv("DIGISEG_API_KEY")
	if apiKey == "" {
		panic("Please provide a Digiseg API key in environment variable DIGISEG_API_KEY")
	}

	audiences, err := GetAudiences(apiKey, "152.115.123.174")
	if err != nil {
		println()
		panic(err)
	}

	println()
	fmt.Printf("Status: %s", audiences.Status)

	for _, audience := range audiences.Audiences {
		println()
		fmt.Printf("Audience:\t%s\t%s", *audience.Code, *audience.Name)
	}
}

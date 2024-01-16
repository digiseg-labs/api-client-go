package main

import "fmt"

func main() {
	accessToken, err := Authenticate()
	if err != nil {
		println()
		panic(err)
	}
	fmt.Printf("Access token: %s", accessToken)

	audiences, err := GetAudiences(accessToken, "152.115.123.174")
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

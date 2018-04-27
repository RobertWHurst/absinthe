package main

import (
	absinthe "github.com/RobertWHurst/Absinthe"
)

func main() {
	_, err := absinthe.Connect(
		absinthe.DefaultURL,
		absinthe.Name("service"),
		absinthe.Version("0.1.0"),
	)

	if err != nil {
		panic(err)
	}

	select {}
}

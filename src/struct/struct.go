package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type User struct {
	Name     string   `json:"name"`
	Age      int64    `json:"age"`
	Hobbies  []string `json:"hobbies,omitempty"`
	Password string   `json:"-"`
}

func main() {
	u := &User{
		Name:     "lilei",
		Age:      12,
		Password: "123455",
	}

	if out, err := json.MarshalIndent(u, "", "  "); err != nil {
		log.Println(err)
		os.Exit(1)
	} else {
		fmt.Println(string(out))
	}
}

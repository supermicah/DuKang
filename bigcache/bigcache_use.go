package main

import (
	"fmt"
)

func main() {

	_ = GetCache().Set("my-unique-key", []byte("value"))

	if entry, err := cache.Get("my-unique-key"); err == nil {
		fmt.Println(string(entry))
	}
}

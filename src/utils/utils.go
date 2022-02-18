package utils

import (
	"fmt"
	"os"
)

func Err(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}

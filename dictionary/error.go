package dictionary

import (
	"fmt"
	"os"
)

func HandleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error: %v", err)
		os.Exit(1)
	}
}

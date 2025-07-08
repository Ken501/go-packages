package file

import (
	"fmt"
	"os"
	"time"
)

func ViewFileContents() {
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Printf("Error: %v occurred at: %v", err, time.Now())
	}
	fmt.Println(f)

}

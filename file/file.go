package file

import (
	"fmt"
	"os"
	"time"
)

func ViewFileContents(n string) {
	s := make([]byte, 128)
	f, err := os.Open(n)
	if err != nil {
		fmt.Printf("Error: %v occurred at: %v", err, time.Now())
	}
	fmt.Printf("Value: %v Type: %T\n", f, f)
	data, err := f.Read(s)
	fmt.Println(data)

}

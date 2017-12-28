package sec1_2

import (
	"os"
	"fmt"
)

func main() {

	for index, param := range os.Args {
		fmt.Printf("index: %d, param: %s\n", index, param)
	}
}
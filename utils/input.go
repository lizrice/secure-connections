package utils

import (
	"bufio"
	"fmt"
	"os"
)

// Wait holds up proceedings until the user presses return
func Wait() {
	fmt.Printf("[Press enter to proceed]")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println()
}

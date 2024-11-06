package shared

import (
	"bufio"
	"fmt"
)

func GetInput(reader *bufio.Scanner, prompt string) string {
	fmt.Print(prompt)
	reader.Scan()
	return reader.Text()
}

package console_apps

import "fmt"

func multiplicationTables(start int, max int) []string {
	var output []string
	for i := 0; i <= max; i++ {
		output = append(output, fmt.Sprintf("%d x %d = %d", start, i, start*i))
	}
	return output
}

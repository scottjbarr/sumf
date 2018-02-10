package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	sum := 0.0

	reader := bufio.NewReader(os.Stdin)

	for {
		t, err := reader.ReadString(' ')
		if err == io.EOF && len(t) == 0 {
			// we've finished reading stdin when there is an EOF and no more
			// values
			break
		}

		// get rid of spaces
		s := strings.TrimSpace(t)

		// convert to a float, ignoring errors. If you want to pump text
		// into this, that is your choice :)
		f, _ := strconv.ParseFloat(s, 64)

		sum += f
	}

	// format the sum to least number sigificant digits
	out := strconv.FormatFloat(sum, 'f', -1, 64)

	fmt.Printf("%s\n", out)
}

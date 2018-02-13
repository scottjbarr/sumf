package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sum := 0.0

	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)

	scanner.Split(split)

	for scanner.Scan() {
		tok := scanner.Bytes()
		f, _ := strconv.ParseFloat(string(tok), 64)
		sum += f
	}

	// format the sum to least number sigificant digits
	out := strconv.FormatFloat(sum, 'f', -1, 64)

	fmt.Printf("%s\n", out)
}

func split(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if len(data) == 0 {
		return
	}

	switch data[0] {
	case ' ', '\n', '\r', '\t':
		advance, token, err = parseWhitespace(data)

	default:
		advance, token, err = parseToken(data)
	}
	return
}

func parseToken(data []byte) (int, []byte, error) {
	var accum []byte
	for i, b := range data {
		if b == ' ' || b == '\n' || b == '\t' || b == '\r' {
			return i, accum, nil
		} else {
			accum = append(accum, b)
		}
	}
	return 0, nil, nil
}

func parseWhitespace(data []byte) (int, []byte, error) {
	var accum []byte
	for i, b := range data {
		if b == ' ' || b == '\n' || b == '\t' || b == '\r' {
			accum = append(accum, b)
		} else {
			return i, accum, nil
		}
	}
	return 0, nil, nil
}

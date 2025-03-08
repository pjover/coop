package cmd

import (
	"fmt"
	"strconv"
)

func ParseInteger(input string, argName string) (int, error) {
	id, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("invalid arg `%s`: %s", argName, input)
	}
	return id, nil
}

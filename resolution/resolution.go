package resolution

import (
	"strconv"
	"strings"
)

type Resolution struct {
	X int
	Y int
}

func New(x int, y int) Resolution {
	return Resolution{x, y}
}

func Parse(resolution string) (Resolution, error) {
	r := strings.Split(resolution, "x")

	if len(r) != 2 {
		return Resolution{}, nil
	}

	x, err := strconv.Atoi(r[0])

	if err != nil {
		return Resolution{}, err
	}

	y, err := strconv.Atoi(r[1])

	if err != nil {
		return Resolution{}, err
	}

	return New(x, y), nil
}

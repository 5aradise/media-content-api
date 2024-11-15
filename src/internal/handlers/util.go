package handlers

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

func getId(idS string) (int32, error) {
	id, err := strconv.Atoi(idS)
	if err != nil {
		return 0, errors.New("invalid id value: invalid number")
	}

	if id < 1 {
		return 0, errors.New("invalid id value: id cannot be lower than 1")
	}

	if id > math.MaxInt32 {
		return 0, fmt.Errorf("invalid id value: id cannot be greater than %d", math.MaxInt32)
	}

	return int32(id), nil
}

package bsearch

import (
	"errors"
)

var (
	NotFoundErr = errors.New("Not found")
)

func LookUp(s []int, v int) (int, error) {
	low := 0
	high := len(s) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := s[mid]
		if guess == v {
			return mid, nil
		}
		if guess > v {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, NotFoundErr
}

func LinearLookUp(s []int, v int) (int, error) {
	for i := 0; i < len(s); i++ {
		if s[i] == v {
			return i, nil
		}
	}
	return 0, NotFoundErr
}

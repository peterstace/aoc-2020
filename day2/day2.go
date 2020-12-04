package day2

import (
	"strconv"
	"strings"
)

func CountValid(input string) int {
	input = strings.TrimSpace(input)
	var count int
	for _, line := range strings.Split(input, "\n") {
		if IsValid(line) {
			count++
		}
	}
	return count
}

func IsValid(input string) bool {
	parts := strings.Split(input, ":")
	pw := strings.TrimSpace(parts[1])
	parts = strings.Split(parts[0], " ")
	ch := parts[1]
	parts = strings.Split(parts[0], "-")
	min, _ := strconv.Atoi(parts[0])
	max, _ := strconv.Atoi(parts[1])

	count := strings.Count(pw, ch)
	return count >= min && count <= max
}

func CountValid2(input string) int {
	input = strings.TrimSpace(input)
	var count int
	for _, line := range strings.Split(input, "\n") {
		if IsValid2(line) {
			count++
		}
	}
	return count
}

func IsValid2(input string) bool {
	parts := strings.Split(input, ":")
	pw := strings.TrimSpace(parts[1])
	parts = strings.Split(parts[0], " ")
	ch := parts[1]
	parts = strings.Split(parts[0], "-")
	min, _ := strconv.Atoi(parts[0])
	max, _ := strconv.Atoi(parts[1])

	w1 := pw[min-1 : min]
	w2 := pw[max-1 : max]

	var count int
	if w1 == ch {
		count++
	}
	if w2 == ch {
		count++
	}
	return count == 1
}

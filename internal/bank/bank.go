package bank

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bank struct {
	Name    string
	BinFrom int
	BinTo   int
}

func LoadFromFile(path string) ([]Bank, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	var banks []Bank
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, ",")
		if len(parts) != 3 {
			continue
		}

		name := strings.TrimSpace(parts[0])

		binFrom, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			continue
		}

		binTo, err := strconv.Atoi(strings.TrimSpace(parts[2]))
		if err != nil {
			continue
		}

		banks = append(banks, Bank{
			Name:    name,
			BinFrom: binFrom,
			BinTo:   binTo,
		})
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return banks, nil
}

func FindBank(banks []Bank, cardNumber string) *Bank {
	if len(cardNumber) < 6 {
		return nil
	}

	prefix, err := strconv.Atoi(cardNumber[:6])
	if err != nil {
		return nil
	}

	for i := range banks {
		if prefix >= banks[i].BinFrom && prefix <= banks[i].BinTo {
			return &banks[i]
		}
	}

	return nil
}

package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Money struct {
	totalCents int
}

func (m Money) Dollars() int {
	return m.totalCents / 100
}

func (m Money) Cents() int {
	return m.totalCents % 100
}

func (m Money) String() string {
	return fmt.Sprintf("%d.%02d", m.Dollars(), m.Cents())
}

func (m Money) Add(m2 Money) Money {
	m.totalCents += m2.totalCents
	return m
}

func ParseMoney(s string) (m Money, err error) {
	parts := strings.Split(s, ".")
	if len(parts) != 2 {
		return m, errors.New("string must be of form 0.00")
	}
	dollars, err := strconv.Atoi(parts[0])
	if err != nil {
		return m, err
	}
	m.totalCents, err = strconv.Atoi(parts[1])
	if err != nil {
		return m, err
	}
	m.totalCents += dollars * 100
	return m, nil
}

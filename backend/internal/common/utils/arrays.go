package utils

import "github.com/shopspring/decimal"

func MinInArray(array ...float64) float64 {
	value := array[0]
	for _, v := range array[1:] {
		if value > v {
			value = v
		}
	}
	return value
}

func MaxInArray(array ...float64) float64 {
	value := array[0]
	for _, v := range array[1:] {
		if value < v {
			value = v
		}
	}
	return value
}

func MaxMinInArray(array ...float64) (max float64, min float64) {
	max = array[0]
	min = array[0]
	for _, v := range array[1:] {
		if max < v {
			max = v
		} else if min > v {
			min = v
		}
	}
	return
}

func MaxMinInDecimalArray(array ...decimal.Decimal) (max decimal.Decimal, min decimal.Decimal) {
	max = array[0]
	min = array[0]
	for _, v := range array[1:] {
		if max.LessThan(v) {
			max = v
		} else if min.GreaterThan(v) {
			min = v
		}
	}
	return
}

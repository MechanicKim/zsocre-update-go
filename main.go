package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		Fields that we should keep
		m1 : mean of a1
		l1 : length of a1
		s1 : sum of square of each values
	*/

	var a1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var l1 = len(a1)
	var m1 = getMean(a1)
	var s1 = getSquareSum(a1)
	var sd1 = getStandardDeviation(l1, s1, m1)
	fmt.Println("m1 : ", m1)
	fmt.Println("sd1 : ", sd1)

	var a2 = []int{11, 12, 13, 14, 15, 17, 20}
	var m2 = updateMean(m1, l1, a2)
	var sd2 = updateStandardDeviation(l1, s1, m2, a2)
	fmt.Println("m2 : ", m2)
	fmt.Println("sd2 :", sd2)

	var a3 = append(a1, a2...)
	var l3 = len(a3)
	var m3 = getMean(a3)
	var s3 = getSquareSum(a3)
	var sd3 = getStandardDeviation(l3, s3, m3)
	fmt.Println("m3 : ", m3)
	fmt.Println("sd3 :", sd3)
}

func getSum(a []int) int {
	sum := 0
	for _, v := range a {
		sum += v
	}
	return sum
}

func getMean(a []int) float64 {
	l := len(a)
	s := getSum(a)
	m := float64(s) / float64(l)
	return m
}

func updateMean(m float64, l int, a []int) float64 {
	result := (m * float64(l)) + float64(getSum(a))
	result /= float64(l + len(a))
	return result
}

func getSquareSum(a []int) float64 {
	sum := 0.0
	for _, v := range a {
		sum += math.Pow(float64(v), 2.0)
	}
	return sum
}

func getStandardDeviation(l int, s float64, m float64) float64 {
	result := (s / float64(l)) - math.Pow(m, 2.0)
	return math.Sqrt(result)
}

func updateStandardDeviation(l int, s float64, m float64, a []int) float64 {
	result := (s + getSquareSum(a)) / float64(l+len(a))
	result -= math.Pow(m, 2.0)
	return math.Sqrt(result)
}

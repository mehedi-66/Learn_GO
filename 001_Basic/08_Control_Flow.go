package main

import "fmt"

func average(x []float64) (avg float64) {
	total := 0.0

	if len(x) == 0 {
		avg = 0
	} else {
		// _ is used to ingore index that provide the range function
		for _, v := range x {
			total += v
		}

		avg = total / float64(len(x))
	}
	return
}

func average2(x []float64) (avg float64) {
	total := 0.0
	switch len(x) {
	case 0:
		avg = 0
	default:
		for _, v := range x {
			total += v
		}
		avg = total / float64(len(x))
	}
	return
}

func main() {
	x := []float64{1.15, 3.14, 42.0, 29.5}
	fmt.Println(average(x)) // 19.19749999999998
}

package main

import "fmt"

func main() {
	firstReleases := map[string]int{
		"Go":         2007,
		"Java":       1995,
		"Python":     1991,
		"JavaScript": 1995,
		"Ruby":       1995,
	}

	/* Loop through each entry and output the name and relase year */
	for k, v := range firstReleases {
		fmt.Printf("%s was first realse in %d\n", k, v)
	}
}

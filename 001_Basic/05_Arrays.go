package main

import "fmt"

func main() {

	// ************************ 1st Array User *************************
	/* Define an array of size 4 that stores deployment options */
	var DeveloymentOptions = []string{"R-pi", "AWS", "GCP", "Azure"}

	/* Loop through the deployment options array */
	for i := 0; i < len(DeveloymentOptions); i++ {
		option := DeveloymentOptions[i]
		fmt.Println(i, option)
	}

	// ******************* 2nd Example *************************

	DepOpstions := [...]string{"R-pi", "AWS", "GCP", "Azure"}

	for index, op := range DepOpstions {
		fmt.Println(index, op)
	}
}

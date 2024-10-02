package maps

import "fmt"

func RunExampleAverageAgeOfMaps() {
	fmt.Println("\nAverage age of maps:")

	people := map[string]int{
		"Alice":   30,
		"Bob":     25,
		"Charlie": 35,
		"David":   28,
	}

	result := averageAgeOfMaps(people)
	fmt.Printf("Средний возраст: %.2f\n", result)
}

func averageAgeOfMaps(people map[string]int) float64 {
	if len(people) == 0 {
		return 0
	}

	totalAge := 0
	for _, age := range people {
		totalAge += age
	}

	average := float64(totalAge) / float64(len(people))
	return average
}

package main

import "fmt"

type map1 map[string]string

func main() {
	nicknames := make(map1, 5)
	nicknames["Charles"] = "Chuck"
	nicknames["Robert"] = "Bob"
	nicknames["Richard"] = "Rick"
	nicknames["Teddy"] = "Ted"
	nicknames["Mohammad"] = "Mo"

	for key, value := range nicknames {
		fmt.Printf("\nThe nickname for %s is %s", key, value)
	}

	// Test for the presence of a key in map
	_, present := nicknames["James"]
	fmt.Println("\nThe key James is present: ", present)

	_, present = nicknames["Teddy"]
	fmt.Println("\nThe key Teddy is present: ", present)
	delete(nicknames, "Robert")

	_, present = nicknames["Robert"]
	fmt.Println("\nThe key Robert is present: ", present)

	nicknames["Charles"] = "Charlie"
}

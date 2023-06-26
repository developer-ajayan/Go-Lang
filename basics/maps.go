package main

import "fmt"

func main() {
	// Initialize an empty map
	emptyMap := make(map[string]int)
	fmt.Println("Empty Map:", emptyMap)

	// Insert key-value pairs into the map
	emptyMap["apple"] = 3
	emptyMap["banana"] = 5
	emptyMap["cherry"] = 2
	fmt.Println("Map after insertion:", emptyMap)

	// Update the value for a key
	emptyMap["banana"] = 7
	fmt.Println("Map after updating value:", emptyMap)

	// Delete a key-value pair from the map
	delete(emptyMap, "cherry")
	fmt.Println("Map after deletion:", emptyMap)

	// Check if a key exists in the map
	_, present := emptyMap["apple"]
	if present {
		fmt.Println("Key 'apple' exists in the map")
	} else {
		fmt.Println("Key 'apple' does not exist in the map")
	}

	// Iterate over keys and values in the map
	for key, value := range emptyMap {
		fmt.Println("Key:", key, "Value:", value)
	}

	// Nested maps
	nestedMap := make(map[string]map[string]string)
	nestedMap["person1"] = make(map[string]string)
	nestedMap["person1"]["name"] = "John"
	nestedMap["person1"]["age"] = "30"
	nestedMap["person2"] = make(map[string]string)
	nestedMap["person2"]["name"] = "Jane"
	nestedMap["person2"]["age"] = "25"
	fmt.Println("Nested Map:", nestedMap)
}

package orderedmap

import "fmt"

func ExampleOrderedMap() {
	data := New[int, string]()
	data.Set(1, "one")
	data.Set(3, "three")
	data.Set(2, "two")
	data.Set(-1, "neg one")
	data.Del(1)

	for _, key := range data.Keys() {
		value, _ := data.Get(key)
		fmt.Println(key, value)
	}

	data.Range(func(key int, value string) bool {
		fmt.Println(key, value)
		return true
	})

	// Output:
	// 3 three
	// 2 two
	// -1 neg one
	// 3 three
	// 2 two
	// -1 neg one
}

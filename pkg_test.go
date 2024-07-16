package ordered

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

func ExampleOrderedSet() {
	set1 := NewSet[int]()
	set1.Add(1)
	set1.Add(3)
	set1.Add(2)
	set1.Add(4)
	set1.Del(1)

	fmt.Println(set1.Elements())

	set2 := NewSet[int]()
	set2.Add(3)
	set2.Add(4)
	set2.Add(5)

	fmt.Println(set2.Elements())
	fmt.Println(set1.Union(set2).Elements())
	fmt.Println(set1.Intersection(set2).Elements())
	fmt.Println(set1.Difference(set2).Elements())

	// Output:
	// [3 2 4]
	// [3 4 5]
	// [3 2 4 5]
	// [3 4]
	// [2]
}

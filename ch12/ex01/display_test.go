package display

type S struct {
	a int
	b int
}

func ExampleDisplay_StructKeyMap() {
	m := make(map[S]int)
	m[S{a: 0, b: 0}] = 0
	m[S{a: 1, b: 1}] = 1

	Display("m", m)
	// Unordered output:
	// Display m (map[display.S]int):
	// m[{"a": 0, "b": 0}] = 0
	// m[{"a": 1, "b": 1}] = 1
}

func ExampleDisplay_ArrayKeyMap() {
	m := make(map[[2]int]int)
	m[[2]int{0, 0}] = 0
	m[[2]int{1, 1}] = 1

	Display("m", m)
	// Unordered output:
	// Display m (map[[2]int]int):
	// m[[0, 0]] = 0
	// m[[1, 1]] = 1
}

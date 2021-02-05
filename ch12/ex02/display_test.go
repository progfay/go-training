package display

type A struct{ t string }
type B struct {
	t string
	a A
}
type C struct {
	t string
	b B
}
type D struct {
	t string
	c C
}

func ExampleDisplay_Recursive() {
	d := D{
		t: "d",
		c: C{
			t: "c",
			b: B{
				t: "b",
				a: A{
					t: "a",
				},
			},
		},
	}
	Display("d", d)
	// Output:
	// Display d (display.D):
	// d.t = "d"
	// d.c.t = "c"
}

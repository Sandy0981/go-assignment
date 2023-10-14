package temperature

var c int

func FahrenheitToCelcius(f int) int {
	c = (f - 32) * 5 / 9
	return c
}

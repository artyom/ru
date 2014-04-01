package ru

// Pluralize возвращает правильную множественную форму для заданного числа
func Pluralize(i int, item0, item1, item3 string) string {
	if i < 0 {
		i = -i
	}
	switch {
	case i%10 == 0:
		return item0
	case i >= 11 && i <= 20:
		return item0
	case i%10 == 1:
		return item1
	case i%10 > 1 && i%10 < 5:
		return item3
	default:
		return item0
	}
}

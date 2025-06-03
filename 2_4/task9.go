package main

import (
	"fmt"
)

func main() {
	var meta [3]int
	for i := 0; i < 3; i++ {
		fmt.Scan(&meta[i])
	}
	var a, b, n int
	a = meta[0] // рубли за один пирожок
	b = meta[1] // копейки за один пирожок
	n = meta[2] // число пирожков

	// Рассчитываем общую стоимость в копейках
	totalKopecks := (a*100 + b) * n

	// Конвертируем в рубли и копейки
	rubles := totalKopecks / 100
	kopecks := totalKopecks % 100

	// Выводим рубли и копейки (обычно сначала рубли)
	fmt.Println(rubles, kopecks)
}

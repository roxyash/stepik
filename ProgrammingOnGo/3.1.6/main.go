// В ходе анализа результатов переписи населения информация была сохранена в объекте типа map:

// groupCity := map[int][]string{
// 	10:   []string{...}, // города с населением 10-99 тыс. человек
// 	100:  []string{...}, // города с населением 100-999 тыс. человек
// 	1000: []string{...}, // города с населением 1000 тыс. человек и более
// }
// При подготовке важного отчета о городах с населением 100-999 тыс.
// человек был подготовлен другой объект типа map:

// cityPopulation := map[string]int{
// 	город: население города в тысячах человек,
// }
// Однако база данных с информацией о точной численности населения содержала ошибки, поэтому в cityPopulation в т.ч.
// была сохранена информация о городах, которые входят в другие группы из groupCity.

// Ваша программа имеет доступ к обоим указанным отображениям, требуется исправить cityPopulation,
// чтобы в ней была сохранена информация только о городах из группы groupCity[100].

// Функция main() уже объявлена, доступ к отображениям осуществляется по указанным именам.
// По результатам выполнения ничего больше делать не требуется, проверка будет осуществлена автоматически.
package main // Delete this string before push stepik

import "fmt" // Delete this string before push stepik

func main() { // Delete this string before push stepik
	groupCity := map[int][]string{ // Delete this string before push stepik
		10:   []string{"Ивановка", "Грязь"},                           // Delete this string before push stepik
		100:  []string{"Королев", "Электроугли", "Село", "Миллионик"}, // Delete this string before push stepik
		1000: []string{"Москва"},                                      // Delete this string before push stepik
	}
	cityPopulation := map[string]int{ // Delete this string before push stepik
		"Село":        123456,  // Delete this string before push stepik
		"Миллионик":   500,     // Delete this string before push stepik
		"Королев":     444444,  // Delete this string before push stepik
		"Электроугли": 1000000, // Delete this string before push stepik
		"Москва":      1000000, // Delete this string before push stepik
		"Питер":       1000000, // Delete this string before push stepik
	} // Delete this string before push stepik

	for key := range cityPopulation {
		flag := false
		for _, value := range groupCity[100] {
			if key == value {
				flag = true
				continue
			}
		}
		if !flag {
			delete(cityPopulation, key)
		}
	}

	fmt.Println(cityPopulation) // Delete this string before push stepik
} // Delete this string before push stepik

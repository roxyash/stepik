// Представьте что вы работаете в большой компании где используется модульная архитектура.
// Ваш коллега написал модуль с какой-то логикой (вы не знаете) и передает в вашу программу какие-то данные.
// Вы же пишете функцию которая считывает две переменных типа string ,
// а возвращает число типа int64 которое нужно получить сложением этих строк.

// Но не было бы так все просто, ведь ваш коллега не пишет на Go, и он зол из-за того, что гоферам платят больше.
// Поэтому он решил подшутить над вами и подсунул вам свинью.
//  Он придумал вставлять мусор в строки перед тем как вызывать вашу функцию.

// Поэтому предварительно вам нужно убрать из них мусор и конвертировать в числа.
// Под мусором имеются ввиду лишние символы и спец знаки.
//  Разрешается использовать только определенные пакеты: fmt, strconv, unicode, strings,  bytes.
// Они уже импортированы, вам ничего импортировать не нужно!

// Считывать и выводить ничего не нужно!

// Ваша функция должна называться adding() !

// Считайте что функция и пакет main уже объявлены!
package main //Remove string before push stepik

import ( //Remove string before push stepik
	"strconv" //Remove string before push stepik
) //Remove string before push stepik

func adding(s1, s2 string) int64 {
	arraysOfRune := [][]int32{[]rune(s1), []rune(s2)}

	var numberStr string
	var strArr []string

	for _, runeArray := range arraysOfRune {
		numberStr = ""
		for _, value := range runeArray {
			if value >= 48 && value <= 57 {
				numberStr += string(value)
			}
		}
		strArr = append(strArr, numberStr)
	}

	var sum int64
	for _, value := range strArr {
		number, _ := strconv.Atoi(value)
		sum += int64(number)
	}
	return sum
}

// 48 49 50 51 52 53 54 55 56 57
func main() { //Remove string before push stepik
	adding("asdas12asdasd", "123asdasd") //Remove string before push stepik
} //Remove string before push stepik

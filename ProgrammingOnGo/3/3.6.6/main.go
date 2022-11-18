package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type AverageRating struct {
	Average float32
}

// {
//     "ID":134,
//     "Number":"ИЛМ-1274",
//     "Year":2,
//     "Students":[
//         {
//             "LastName":"Вещий",
//             "FirstName":"Лифон",
//             "MiddleName":"Вениаминович",
//             "Birthday":"4апреля1970года",
//             "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
//             "Phone":"+7(948)709-47-24",
//             "Rating":[1,2,3,4,5,6,7,8,9]
//         },
// 		{
//             "LastName":"Вещий",
//             "FirstName":"Лифон",
//             "MiddleName":"Вениаминович",
//             "Birthday":"4апреля1970года",
//             "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
//             "Phone":"+7(948)709-47-24",
//             "Rating":[1,2,3]
//         },
// 		{
//             "LastName":"Вещий",
//             "FirstName":"Лифон",
//             "MiddleName":"Вениаминович",
//             "Birthday":"4апреля1970года",
//             "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
//             "Phone":"+7(948)709-47-24",
//             "Rating":[1,2,3]
//         }
//     ]
// }

func main() {

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(string(data))
	}

	group := Group{}

	err = json.Unmarshal(data, &group)
	if err != nil {
		fmt.Printf("Umarshall err: %v", err)
	}

	var cntAllRating int
	for _, student := range group.Students {
		cntAllRating += len(student.Rating)
	}

	averageRating := AverageRating{
		Average: float32(cntAllRating/len(group.Students)),
	}

	res, err := json.MarshalIndent(averageRating, "", "    ")
	if err != nil {
		fmt.Printf("Marshall err: %v", err)
	}

	_, err = os.Stdout.Write(res)
	

}

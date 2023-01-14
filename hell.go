// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"os"
// )

// type Person struct {
// 	LastName string
// 	FirstName string
// 	MiddleName string
// 	Birthday string
// 	Address string
// 	Phone string
// 	Rating []int
// }

// type myStruct struct {
// 	ID int
// 	Number string
// 	Year int
// 	Students []Person
// }

// type Answer struct {
// 	Average float64
// }

// func main() {
//     data, err := ioutil.ReadAll(os.Stdin)
//     if err != nil {
//         panic(err)
//     }

// 	var s myStruct
// 	if err := json.Unmarshal(data, &s); err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// fmt.Printf("%v", s)

// 	var count_grade int

// 	for _, value := range s.Students {
// 		count_grade += len(value.Rating)
// 	}

// 	an := Answer{Average: float64(count_grade) / float64(len(s.Students))}

// 	d, err := json.MarshalIndent(an, "", "    ")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	// fmt.Printf("%s", d)
// 	_, err = os.Stdout.Write(d)

// }

// // {
// //     "ID":134,
// //     "Number":"ИЛМ-1274",
// //     "Year":2,
// //     "Students":[
// //         {
// //             "LastName":"Вещий",
// //             "FirstName":"Лифон",
// //             "MiddleName":"Вениаминович",
// //             "Birthday":"4апреля1970года",
// //             "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
// //             "Phone":"+7(948)709-47-24",
// //             "Rating":[1,2,3]
// //         },
// // 		{
// //             "LastName":"Вещий",
// //             "FirstName":"Лифон",
// //             "MiddleName":"Вениаминович",
// //             "Birthday":"4апреля1970года",
// //             "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
// //             "Phone":"+7(948)709-47-24",
// //             "Rating":[1,2,3, 4, 5, 6]
// //         }
// //     ]
// // }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	_ "fmt"
// 	"net/http"
// )

// type dataGov struct {
// 	Global_id int64
// }

// func main() {
// 	resp, _ := http.Get("https://raw.githubusercontent.com/semyon-dev/stepik-go/master/work_with_json/data-20190514T0100.json")
// 	gov := []dataGov{}
// 	json.NewDecoder(resp.Body).Decode(&gov)

// 	var sum int64

// 	for _, value := range gov {
// 		sum += value.Global_id
// 	}

// 	fmt.Println(sum)
// }

package main

import "fmt"
import "strconv"


func main() {
	fmt.Println(delete(727178))
}
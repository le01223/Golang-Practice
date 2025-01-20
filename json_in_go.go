package main

import (
	"encoding/json" // подключаем парсер для работы с json
	"fmt"
)

type Worker struct {
	Name string `json:"name"` /* этот тег говорит парсеру, что Name
	должно быть представлено в JSON как поле name*/
	Age int    `json:"age"`
	Job string `json:"job"`
}

func main() {
	//Example of Unmarshalling
	var worker Worker
	jsDt := `{"name":"Michael", "age":50, "job":"A doctor"}`
	err := json.Unmarshal([]byte(jsDt), &worker)
	if err != nil {
		fmt.Println("Ошибка!!", err)
	}
	fmt.Println(worker.Age) // возьмет из json'а 50
	//Example of Marshalling
	goDt := Worker{Name: "Lev", Age: 35, Job: "Lawyer"}
	jsDt2, err2 := json.Marshal(goDt)
	if err2 != nil {
		fmt.Println("Ошибка!!", err2)
	}
	fmt.Println(string(jsDt2)) // выведет json goDt
}

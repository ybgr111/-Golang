package main

import (
	"net/http"
	"ybgr111/tasks"
)

func main() {

	http.HandleFunc("/task/Циклическая ротация", tasks.ServeHttpCyclic)
	http.HandleFunc("/task/Чудные вхождения в массив", tasks.ServeHttpWonderful)
	http.HandleFunc("/task/Проверка последовательности", tasks.ServeHttpSequence)
	http.HandleFunc("/task/Поиск отсутствующего элемента", tasks.ServeHttpSearching)
	http.HandleFunc("/tasks/", tasks.ServeHttpAll)

	err := http.ListenAndServe("localhost:3000", nil)

	if err != nil {
		panic(err)
	}
}

package tasks

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"
	"sort"
	"ybgr111/types"
)

func ServeHttpSearching(w http.ResponseWriter, r *http.Request) {
	log.Println(string(SearchingTusk()))
}

func SearchingTusk() []byte {
	response, err := http.Get("https://kuvaev-ituniversity.vps.elewise.com/tasks/Поиск отсутствующего элемента")

	if err != nil {
		log.Println(err)
		return nil
	}

	var data [][]interface{}

	err = json.NewDecoder(response.Body).Decode(&data)

	if err != nil {
		log.Println(err)
		return nil
	}

	var results []types.TasksResult

	for _, arr := range data {
		var child types.TasksResult

		for _, item := range arr {
			el := reflect.ValueOf(item)

			switch el.Type().Kind() {
			case reflect.Slice:
				for i := 0; i < el.Len(); i++ {
					child.Slice = append(child.Slice, int(el.Index(i).Elem().Float()))
				}
			default:
			}
		}
		results = append(results, child)
	}

	var answer []interface {
	}

	for i := 0; i < len(results); i++ {
		answer = append(answer, SolutionS(results[i].Slice))
	}

	var send types.SendData

	send = types.SendData{
		User_name: "ybgr111",
		Task:      "Поиск отсутствующего элемента",
		Results: types.ChildData{
			Payload: data,
			Results: answer,
		},
	}

	body, err := json.Marshal(send)

	if err != nil {
		log.Println(err)
		return nil
	}

	resp, err := http.Post("https://kuvaev-ituniversity.vps.elewise.com/tasks/solution", "application/json", bytes.NewBuffer(body))

	if err != nil {
		log.Println(err)
		return nil
	}

	respData, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return nil
	}
	return respData
}

func SolutionS(A []int) int {

	N := len(A)

	sort.Ints(A)

	if A[N-1] != N+1 {
		return N + 1
	}

	for i := 1; i < N; i++ {
		if A[i-1] != i {
			return i
		}
	}
	return N
}

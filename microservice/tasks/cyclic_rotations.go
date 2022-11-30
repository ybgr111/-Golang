package tasks

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"reflect"
	"ybgr111/types"
)

func ServeHttpCyclic(w http.ResponseWriter, r *http.Request) {
	log.Println(string(CyclicTusk()))
}

func CyclicTusk() []byte {
	response, err := http.Get("https://kuvaev-ituniversity.vps.elewise.com/tasks/Циклическая ротация")

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

	var results []types.TasksResultWithK

	for _, arr := range data {
		var child types.TasksResultWithK

		for _, item := range arr {
			el := reflect.ValueOf(item)

			switch el.Type().Kind() {
			case reflect.Slice:
				for i := 0; i < el.Len(); i++ {
					child.Slice = append(child.Slice, int(el.Index(i).Elem().Float()))
				}
			case reflect.Float64:
				child.K = int(el.Float())
			default:
			}
		}
		results = append(results, child)
	}

	var answer []interface {
	}

	for i := 0; i < len(results); i++ {
		answer = append(answer, SolutionC(results[i].Slice, results[i].K))
	}

	var send types.SendData

	send = types.SendData{
		User_name: "ybgr111",
		Task:      "Циклическая ротация",
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

func SolutionC(A []int, K int) []int {

	if K < 0 || len(A) == 0 {
		return A
	}

	rs := len(A) - K%len(A)

	A = append(A[rs:], A[:rs]...)

	return A
}

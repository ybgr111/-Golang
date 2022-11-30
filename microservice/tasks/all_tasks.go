package tasks

import (
	"log"
	"net/http"
)

func ServeHttpAll(w http.ResponseWriter, r *http.Request) {

	var all_answers []string

	all_answers = append(all_answers, string(SequenceTask()))
	all_answers = append(all_answers, string(CyclicTusk()))
	all_answers = append(all_answers, string(SearchingTusk()))
	all_answers = append(all_answers, string(WonderfulTask()))

	log.Println(all_answers)
}

package types

type TasksResult struct {
	Slice []int
}

type TasksResultWithK struct {
	Slice []int
	K     int
}

type SendData struct {
	User_name string    `json:"user_name"`
	Task      string    `json:"task"`
	Results   ChildData `json:"results"`
}

type ChildData struct {
	Payload interface{} `json:"payload"`
	Results interface{} `json:"results"`
}

type GetData struct {
	Percent int `json:"percent"`
	Fails   []struct {
		OriginalResult int `json:"OriginalResult"`
		ExternalResult int `json:"ExternalResult"`
	} `json:"fails"`
}

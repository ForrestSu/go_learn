package proto

type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type JobInfo struct {
	Skills []string `json:"skills"`
}

type Employee struct {
	Basic BasicInfo `json:"basic"`
	Job   JobInfo   `json:"job"`
}

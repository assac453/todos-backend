package entity

type Todo struct {
	Id        int32  `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

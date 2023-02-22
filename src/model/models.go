package models

type Gender struct {
	Id		string		`json:"genId"`
	Name	string		`json:"name"`
}

type Category struct {
	Id		string		`json:"catId"`
	Name	string		`json:"name"`
}

type Album struct {
	Id			string 		`json:"albId"`
	Title		string		`json:"title"`
	Gender		Gender		`json:"Gender"`
	Category	Category	`json:"Category"`
}
package models

type Tag struct {
	ID   int    `json:"id_emp"`
	IDem int    `json:"id_no_emp"`
	Name string `json:"name_th"`
}

type ArrUser struct {
	ID      int
	IDEmp   int
	NameEmp string
}

package models

type TagCompany struct {
	Id_comp      int32  `json:"id_emp"`
	Name_comp_th string `json:"Name_comp_th"`
	Name_comp_en string `json:"name_comp_en"`
	Noaddress    string `json:"noaddress"`
	Soi          string `json:"soi"`
	Road         string `json:"road"`
	District     string `json:"district"`
	Amphur       string `json:"amphur"`
	Province     string `json:"province"`
	Zipcode      string `json:"zipcode"`
	Tel_call     string `json:"tel_call"`
	Tel_fax      string `json:"tel_fax"`
	Ss_number    string `json:"ss_number"`
	Tax_number   string `json:"tax_number"`
	Bank_number  string `json:"bank_number"`
	Co_number    string `json:"co_number"`
	Fund_number  string `json:"fund_number"`
	Status_comp  string `json:"status_comp"`
}

type ArrComp struct {
	Id_comp      int32
	Name_comp_th string
	Name_comp_en string
	Noaddress    string
	Soi          string
	Road         string
	District     string
	Amphur       string
	Province     string
	Zipcode      string
	Tel_call     string
	Tel_fax      string
	Ss_number    string
	Tax_number   string
	Bank_number  string
	Co_number    string
	Fund_number  string
	Status_comp  string
}

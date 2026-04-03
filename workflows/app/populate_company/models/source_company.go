package models


type SourceCompany struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	Website         string `json:"website"`
	LongDescription string `json:"long_description"`
}

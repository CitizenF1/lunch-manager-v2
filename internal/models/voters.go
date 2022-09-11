package models

type Voter struct {
	UserName string `json:"user_name"`
	UserID   string `json:"user_id"`
}

type Voters struct {
	Voters []Voter `json:"voters"`
}

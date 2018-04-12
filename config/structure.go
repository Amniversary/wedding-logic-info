package config

type GetProcessList struct {
	WeddingId int64 `json:"weddingId"`
}


type ProcessRes struct {
	ID        int64  `json:"id"`
	Time      string `json:"time"`
	PlanText  string `json:"planText"`
	Principal string `json:"principal"`
	CreateAt  int64  `json:"create_at"`
}

type DelProcess struct {
	Id int64 `json:"id"`
}
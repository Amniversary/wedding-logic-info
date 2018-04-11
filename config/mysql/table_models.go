package mysql

type Process struct {
	ID        int64  `gorm:"primary_key" json:"id"`
	WeddingId int64  `gorm:"not null; default:0; type:int; index" json:"weddingId"`
	Time      string `grom:"not null; default:''; type:varchar(128)" json:"time"`
	PlanText  string `gorm:"not null; default:''; type:text" json:"planText"`
	Principal string `gorm:"not null; default:''; type:text" json:"principal"`
	UpdateAt  int64  `gorm:"not null; default:0; type:int" json:"update_at"`
	CreateAt  int64  `gorm:"not null; default:0; tyoe:int" json:"create_at"`
}

func (Process) TableName() string {
	return "cProcess"
}
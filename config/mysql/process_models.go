package mysql

import (
	"log"
	"time"
)

func NewProcess(req *Process) bool {
	req.CreateAt = time.Now().Unix()
	if err := db.Create(&req).Error; err != nil {
		log.Printf("newProcess create err: [%v]", err)
		return false
	}
	return true
}

func UpProcess(req *Process) bool {
	req.UpdateAt = time.Now().Unix()
	if err := db.Model(&Process{}).Where("id = ?", req.ID).Update(&req).Error; err != nil {
		log.Printf("upProcess update err: [%v]", err)
		return false
	}
	return true
}
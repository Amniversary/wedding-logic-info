package mysql

import (
	"log"
	"time"
	"github.com/Amniversary/wedding-logic-info/config"
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

func GetProcessList(weddingId int64) ([]config.ProcessRes, bool) {
	var list []config.ProcessRes
	err := db.Table("cProcess").
		Select("id, `time`, plan_text, principal, create_at").
		Where("wedding_id = ?", weddingId).
		Order("time asc").Find(&list).Error
	if err != nil {
		log.Printf("getProcessList query err: [%v]", err)
		return nil, false
	}
	return list, true
}

func DelProcess(id int64) bool {
	if err := db.Where("id = ?", id).Delete(&Process{}).Error; err != nil {
		log.Printf("delProcess query err: [%v]", err)
		return false
	}
	return true
}
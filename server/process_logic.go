package server

import (
	"net/http"
	"github.com/Amniversary/wedding-logic-info/config"
	"github.com/Amniversary/wedding-logic-info/config/mysql"
	"encoding/json"
	"log"
)

/**
	TODO: 创建婚礼流程
 */
func (s *Server) NewProcess(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code:config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &mysql.Process{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("newProcess json decode err : [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if req.WeddingId == 0 || req.Time == "" {
		Response.Msg = "params can not be empty."
		log.Printf("%v : [%v]", Response.Msg, req)
		return
	}
	if ok := mysql.NewProcess(req); !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}
/**
	TODO: 更新婚礼流程
 */
func (s *Server) UpProcess(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &mysql.Process{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("upProcess json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if req.ID == 0 {
		Response.Msg = "params can not be empty."
		log.Printf("%v : [%v]", Response.Msg, req)
		return
	}
	if ok := mysql.UpProcess(req); !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 获取婚礼流程列表
 */
func (s *Server) GetProcessList(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetProcessList{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("getProcessList json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if req.WeddingId == 0 {
		Response.Msg = "params can not be empty."
		log.Printf("%v : [%v]", Response.Msg, req.WeddingId)
		return
	}
	list, ok := mysql.GetProcessList(req.WeddingId)
	if !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Data = list
	Response.Code = config.RESPONSE_OK
}
/**
	TODO: 删除婚礼流程
 */
func (s *Server) DelProcess(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.DelProcess{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("delProcess json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if ok := mysql.DelProcess(req.Id); !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}
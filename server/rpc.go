package server

import (
	"log"
	"time"
	"net/http"
	"encoding/json"
	"fmt"

	"github.com/Amniversary/wedding-logic-server/config"
)

const (
	ServerName = "weddingInfo"
)

type MethodFunc func(w http.ResponseWriter, r *http.Request)

func (s *Server) rpc(w http.ResponseWriter, r *http.Request) {
	res := &config.Response{Code: config.RESPONSE_OK}
	if r.Method != "POST" {
		log.Printf("Method not be Post Request [%s]\n", r.Method)
		EchoJson(w, http.StatusOK, res)
		return
	}
	serverName := r.Header.Get("ServerName")
	if serverName != ServerName {
		log.Printf("ServerName: [%s]  request -> ServerName: [%s] Method: [%s]\n", ServerName, serverName, r.Method)
		EchoJson(w, http.StatusOK, res)
		return
	}

	methodName := r.Header.Get("MethodName")
	start := time.Now()
	defer func() {
		log.Printf("Request MethodName: [%s], Rtime[%v]\n", methodName, time.Now().Sub(start))
	}()
	methodExc, ok := s.methodMap[methodName]
	if !ok {
		res.Code = 1
		res.Msg = fmt.Sprintf("Can't find the interface: [%s]", methodName)
		EchoJson(w, http.StatusOK, res)
		return
	}
	methodExc(w, r)
}

// TODO @ 输出Json数据
func EchoJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type,servername,methodname,userid,msgid")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func (s *Server) initMap() {
	var MethodMap = map[string]MethodFunc{
		"newProcess":     s.NewProcess,
		"upProcess":      s.UpProcess,
		"getProcessList": s.GetProcessList,
		"delProcess":     s.DelProcess,
		"getProcessInfo": s.GetProcessInfo,
	}
	s.methodMap = MethodMap
}

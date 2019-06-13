package master

import (
	"net"
	"net/http"
	"time"
)

type ApiServer struct {
	httpServer *http.Server
}

var (
	G_apiServer * ApiServer
)

// 保存任务接口
// POST job={"name": "job1", "command": "echo hello", "cronExpr": "* * * * *"}
func handleJobSave(resp http.ResponseWriter, req *http.Request) {



}


func InitApiServer()(err error) {
	var (
		mux        *http.ServeMux
		listener   net.Listener
		httpServer *http.Server
	)

	mux = http.NewServeMux()
	mux.HandleFunc("/job/save",handleJobSave)

	//启动TCP监听
	if listener,err = net.Listen("tcp",":8090"); err != nil {
		return
	}

	httpServer = &http.Server{
		ReadTimeout:5 * time.Second,
		WriteTimeout:5 * time.Second,
		Handler:mux,
	}

	//单例模式初始化
	G_apiServer = &ApiServer{
		httpServer:httpServer,
	}

	go httpServer.Serve(listener)

	return
}
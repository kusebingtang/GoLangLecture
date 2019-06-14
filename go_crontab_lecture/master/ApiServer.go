package master

import (
	"GoLecture/go_crontab_lecture/common"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"
)

type ApiServer struct {
	httpServer *http.Server
}

var (
	G_apiServer *ApiServer
)

// 保存任务接口
// POST job={"name": "job1", "command": "echo hello", "cronExpr": "* * * * *"}
func handleJobSave(resp http.ResponseWriter, req *http.Request) {

	var (
		err     error
		postJob string
		addJob  common.Job
		bytes   []byte
		oldJob  *common.Job
	)

	if err = req.ParseForm(); err != nil {
		goto ERR
	}

	postJob = req.PostForm.Get("job")

	//fmt.Println(postJob)

	//json 反序列化addJob
	if err = json.Unmarshal([]byte(postJob), &addJob); err != nil {
		goto ERR
	}

	//fmt.Println(addJob)

	if oldJob, err = G_jobMgr.SaveJob(&addJob); err != nil {
		goto ERR
	}
	//返回成功应答
	if bytes, err = common.BuildResponse(0, "success", oldJob); err == nil {
		resp.Write(bytes)
	}
	return

ERR:
	fmt.Println(err)
}



//删除任务接口
//Post请求  /job/delete name=job1
func handleJobDelete(resp http.ResponseWriter, req *http.Request) {

	var(
		err error
		deleteJobName string
		oldJob  *common.Job
		bytes   []byte
	)

	fmt.Println("enter handleJobDelete")

	if err = req.ParseForm(); err != nil {
		goto ERR
	}

	deleteJobName = req.PostForm.Get("name")

	if oldJob,err = G_jobMgr.DeleteJob(deleteJobName); err!=nil {
		goto ERR
	}

	//正常应答
	if bytes,err = common.BuildResponse(0,"success",oldJob);err == nil {
		resp.Write(bytes)
	}

	return

ERR:
	if bytes,err = common.BuildResponse(-1,err.Error(),nil);err == nil {
		resp.Write(bytes)
	}
}

func InitApiServer() (err error) {
	var (
		mux        *http.ServeMux
		listener   net.Listener
		httpServer *http.Server
	)

	mux = http.NewServeMux()
	mux.HandleFunc("/job/save", handleJobSave)
	mux.HandleFunc("/job/delete", handleJobDelete)

	//启动TCP监听
	if listener, err = net.Listen("tcp", ":"+strconv.Itoa(G_config.ApiPort)); err != nil {
		return
	}

	httpServer = &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      mux,
	}

	//单例模式初始化
	G_apiServer = &ApiServer{
		httpServer: httpServer,
	}

	go httpServer.Serve(listener)

	return
}

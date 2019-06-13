package master

import (
	"GoLecture/go_crontab_lecture/common"
	"context"
	"encoding/json"
	"go.etcd.io/etcd/clientv3"
	"time"
)

//任务管理器
type JobMgr struct {
	client *clientv3.Client
	kv     clientv3.KV
	lease  clientv3.Lease
}

var (
	G_jobMgr *JobMgr
)

func InitJobMgr() (err error) {

	var (
		config clientv3.Config
		client *clientv3.Client
		kv     clientv3.KV
		lease  clientv3.Lease
	)

	config = clientv3.Config{
		Endpoints:   G_config.EtcdEndpoints,                                     //集群地址
		DialTimeout: time.Duration(G_config.EtcdDialTimeout) * time.Millisecond, //连接超时

	}

	if client, err = clientv3.New(config); err != nil {
		return
	}

	kv = clientv3.NewKV(client)
	lease = clientv3.NewLease(client)

	G_jobMgr = &JobMgr{
		client: client,
		kv:     kv,
		lease:  lease,
	}

	//fmt.Println(client)

	return
}

// 保存任务
func (jobMgr *JobMgr) SaveJob(job *common.Job) (oldJob *common.Job, err error) {
	// 把任务保存到/cron/jobs/任务名 -> json

	var (
		jobKey        string
		jobBytesValue []byte
		putResp       *clientv3.PutResponse
		oldJobObject  common.Job
	)

	jobKey = common.JOB_SAVE_DIR + job.Name

	if jobBytesValue, err = json.Marshal(job); err != nil {
		return
	}

	//保存到etcd
	if putResp, err = jobMgr.kv.Put(context.TODO(), jobKey, string(jobBytesValue), clientv3.WithPrevKV()); err != nil {
		return
	}

	//put success 返回久值
	if putResp.PrevKv != nil {
		if err = json.Unmarshal(putResp.PrevKv.Value, &oldJobObject); err != nil {
			return
		}
		oldJob = &oldJobObject
	}
	return
}

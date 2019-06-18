package master

import (
	"GoLecture/go_crontab_lecture/common"
	"context"
	"encoding/json"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
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


func (jobMgr *JobMgr) DeleteJob(jobName string) (oldJob *common.Job, err error) {

	var (
		jobKey        string
		deleteResp    *clientv3.DeleteResponse
		oldJobObject  common.Job
	)

	jobKey = common.JOB_SAVE_DIR + jobName


	if deleteResp ,err = jobMgr.client.Delete(context.TODO(),jobKey,clientv3.WithPrevKV()); err!=nil {
		return
	}

	//返回被删除的任务信息
	if len(deleteResp.PrevKvs) >0 {
		if err = json.Unmarshal(deleteResp.PrevKvs[0].Value, &oldJobObject); err!=nil {
			err = nil
			return
		}
		oldJob = &oldJobObject
	}
	return
}



func (jobMgr *JobMgr) ListWorkers() (jobList []*common.Job, err error) {

	var (
		dirKey string
		getRes *clientv3.GetResponse
		kvPair *mvccpb.KeyValue
		jobObj *common.Job
	)

	dirKey = common.JOB_SAVE_DIR

	if getRes,err = jobMgr.kv.Get(context.TODO(),dirKey,clientv3.WithPrefix());err !=nil {
		return
	}

	jobList = make([]*common.Job, 0)

	for _,kvPair = range  getRes.Kvs {
		jobObj = &common.Job{}
		if err = json.Unmarshal(kvPair.Value,jobObj); err != nil {
			err = nil
			continue
		}

		jobList = append(jobList,jobObj)
	}
	return
}




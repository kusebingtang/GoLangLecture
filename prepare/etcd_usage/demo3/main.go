package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main()  {

	var (
		config      clientv3.Config
		client      *clientv3.Client
		err         error
		kv          clientv3.KV
		getResponse *clientv3.GetResponse

	)

	config = clientv3.Config{
		Endpoints: []string{"47.103.88.171:2379"},
		DialTimeout: time.Second * 5,
	}

	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	kv = clientv3.NewKV(client)

	if getResponse, err = kv.Get(context.Background(),"/cron/jobs/job1"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(getResponse.Kvs,"   ",  getResponse.Count)
	}
}

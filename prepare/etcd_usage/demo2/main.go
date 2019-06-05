package main

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"time"
)

func main()  {

	var (
		config clientv3.Config
		client *clientv3.Client
		err error
		putResponse *clientv3.PutResponse

	)

	config  = clientv3.Config{
		Endpoints: []string{"47.103.88.171:2379"},
		DialTimeout: time.Second * 5,
	}


	if client, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}

	kv :=  clientv3.NewKV(client)

	if putResponse ,err = kv.Put(context.Background(),"/cron/jobs/job1","hello world 2!", clientv3.WithPrevKV()); err !=nil {
		fmt.Println(err)
	} else {
		fmt.Println(putResponse.Header.Revision)
		if putResponse.PrevKv != nil {
			fmt.Println("PrevKV value:", string(putResponse.PrevKv.Value))
		}
	}
}

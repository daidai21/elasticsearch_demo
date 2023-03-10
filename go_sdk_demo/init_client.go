package main

import (
	"fmt"
	es "github.com/olivere/elastic"
	"log"
)

var ESClient *es.Client

func init() {
	fmt.Println("ES version: ", es.Version)

	var err error
	ESClient, err = es.NewClient(
		es.SetBasicAuth("", ""),
		// TODO: 设置参数
		es.SetTraceLog(log.Default()),
	)

	if err != nil {
		fmt.Println(err)
		//panic(err)
	}
	fmt.Println("ESClient: ", ESClient)

}

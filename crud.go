package main

import (
	"context"
	"fmt"

	es6 "github.com/olivere/elastic"
)

var ES6Client *es6.Client

// initClient 初始化客户端
func initClient() {
	var err error
	ES6Client, err = es6.NewClient(es6.SetHttpClient(nil))
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
}

// CRUD
func crud() {
	{
		// select * from ${index} where id = 123;
		query := es6.NewBoolQuery().Must(es6.NewTermQuery("id", 123))
		ctx := context.Background()
		result, err := ES6Client.Search("index").
			Query(query).
			Do(ctx)
		fmt.Println(result)
		fmt.Println(err)
	}

	{
		// select * from ${index} where name like "%123%";
		query := es6.NewBoolQuery().Must(es6.NewMatchQuery("id", 123))
		ctx := context.Background()
		result, err := ES6Client.Search("index").
			Query(query).
			Do(ctx)
		fmt.Println(result)
		fmt.Println(err)
	}

	{
		// select sum(amount) as sum_amount from ${index} where user_id = 123;
		query := es6.NewBoolQuery().Filter(es6.NewTermQuery("user_id", 123))
		agg := es6.NewSumAggregation().Field("amount")
		ctx := context.Background()
		result, err := ES6Client.Search("index").
			Query(query).
			Aggregation("sum_amount", agg).
			Do(ctx)
		fmt.Println(result)
		fmt.Println(err)
		fmt.Println(result.Aggregations.Sum("sum_amount"))
	}

	{
		// select user_id, sum(amount) as sum_amount
		// from ${index}
		// where user_id = 123
		// 		and create_time between '19901010' and '20201231'
		// group by category
		// order by sum_amount
		// limit 10
		// ;
		query := es6.NewBoolQuery().Filter(
			es6.NewTermQuery("user_id", 123),
			es6.NewRangeQuery("create_time").Format("yyyyMMdd").Gte("19901010").Lt("20201231"),
		)
		agg := es6.NewTermsAggregation().
			Field("category").
			Size(10).
			Order("sum_amount", false).
			SubAggregation("sum_amount", es6.NewSumAggregation().Field("amount"))
		ctx := context.Background()
		result, err := ES6Client.Search("index").
			Query(query).
			Aggregation("agg_user_id", agg).
			Do(ctx)
		fmt.Println(result)
		fmt.Println(err)
	}

	{
		// 增+删+改 建议走flink流式写入
	}
}

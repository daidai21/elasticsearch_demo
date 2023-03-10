package main

import (
	"context"
	"fmt"
	es6 "github.com/olivere/elastic"
)

/*

先条件查询 再模糊查询 再定义打分函数

- 性别为: 男
- 匹配自我介绍： 篮球
- 成绩越好越拍前： 语文 * 1 + 数学 * 2 分数

- 第 5 页开始
- 每页10项
*/

func main() {

	score := es6.NewFunctionScoreQuery().
		BoostMode("multiply").
		ScoreMode("sum").
		Add(es6.NewQueryStringQuery("chinese"), es6.NewWeightFactorFunction(1)).
		Add(es6.NewQueryStringQuery("math"), es6.NewWeightFactorFunction(2))

	query := es6.NewBoolQuery().
		Must(es6.NewBoolQuery().QueryName("sex")).
		Must(es6.NewMatchQuery("intro", "篮球")).Must(score)

	fmt.Println(query.Source())
	fmt.Println(query)
	fmt.Println(query.Query)

	do, e := ESClient.Search().
		Query(query).
		From(5).
		Size(10).
		Do(context.Background())

	fmt.Println(do)
	fmt.Println(e)
}

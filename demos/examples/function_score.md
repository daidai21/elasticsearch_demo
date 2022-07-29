# function_score 评分函数


房子的坐标和价格

```shell
# 创建索引：
> curl -X PUT "localhost:9200/london_house_info?pretty" -H 'Content-Type: application/json' -d'
    {
        "mappings": {
            "_doc": {
                "properties": {
                    "location": {
                        "type": "geo_point"
                    },
                    "price": {
                        "type": "integer"
                    }
                }
            }
        }
    }
    '


# 查看创建的索引
> curl -X GET 'localhost:9200/_cat/indices?v'
> curl 'localhost:9200/london_house_info/_mapping?pretty=true'

# 插入mock数据
./london_house_info_mock_data.sh

```


自定义查询评价函数：

距离这里近`"location" : { "lat" : 11, "lon" : 12 }`，且价格不超过100

```shell
> curl 'localhost:9200/london_house_info/_search?pretty=true' -H 'Content-Type: application/json' -d '
    {
        "query": {
            "function_score": {
                "functions": [
                    {
                        "gauss": {
                            "location": {
                                "origin": {
                                    "lat": 11,
                                    "lon": 12
                                },
                                "offset": "2km",
                                "scale": "3km"
                            }
                        }
                    },
                    {
                        "gauss": {
                            "price": {
                                "origin": "50",
                                "offset": "50",
                                "scale": "20"
                            }
                        },
                        "weight": 2
                    }
                ]
            }
        }
    }
'
```

result:

```json
{
  "took" : 42,
  "timed_out" : false,
  "_shards" : {
    "total" : 5,
    "successful" : 5,
    "skipped" : 0,
    "failed" : 0
  },
  "hits" : {
    "total" : 6,
    "max_score" : 2.0,
    "hits" : [
      {
        "_index" : "london_house_info",
        "_type" : "_doc",
        "_id" : "2",
        "_score" : 2.0,
        "_source" : {
          "location" : {
            "lat" : 11,
            "lon" : 12
          },
          "price" : 50
        }
      },
      {
        "_index" : "london_house_info",
        "_type" : "_doc",
        "_id" : "1",
        "_score" : 2.0,
        "_source" : {
          "location" : {
            "lat" : 11,
            "lon" : 12
          },
          "price" : 100
        }
      },
      {
        "_index" : "london_house_info",
        "_type" : "_doc",
        "_id" : "3",
        "_score" : 2.0,
        "_source" : {
          "location" : {
            "lat" : 11.01,
            "lon" : 12
          },
          "price" : 100
        }
      },
      {
        "_index" : "london_house_info",
        "_type" : "_doc",
        "_id" : "5",
        "_score" : 1.0,
        "_source" : {
          "location" : {
            "lat" : 11,
            "lon" : 12
          },
          "price" : 120
        }
      },
      {
        "_index" : "london_house_info",
        "_type" : "_doc",
        "_id" : "6",
        "_score" : 0.026278013,
        "_source" : {
          "location" : {
            "lat" : 11,
            "lon" : 12
          },
          "price" : 150
        }
      },
      {
        "_index" : "london_house_info",
        "_type" : "_doc",
        "_id" : "4",
        "_score" : 0.0,
        "_source" : {
          "location" : {
            "lat" : 13,
            "lon" : 15
          },
          "price" : 100
        }
      }
    ]
  }
}
```

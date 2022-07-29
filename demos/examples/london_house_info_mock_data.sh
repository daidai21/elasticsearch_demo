#!/usr/bin/env bash

set -ex

# insert
curl -X PUT 'localhost:9200/london_house_info/_doc/1?pretty' -H 'Content-Type: application/json' -d '
    {
        "location": {
            "lat": 11,
            "lon": 12
        },
        "price": 100
    }
'

# 查询
curl 'localhost:9200/london_house_info/_doc/1?pretty=true'

# insert
curl -X PUT 'localhost:9200/london_house_info/_doc/2?pretty' -H 'Content-Type: application/json' -d '
    {
        "location": {
            "lat": 11,
            "lon": 12
        },
        "price": 50
    }
'

# insert
curl -X PUT 'localhost:9200/london_house_info/_doc/3?pretty' -H 'Content-Type: application/json' -d '
    {
        "location": {
            "lat": 11.01,
            "lon": 12
        },
        "price": 100
    }
'

# insert
curl -X PUT 'localhost:9200/london_house_info/_doc/4?pretty' -H 'Content-Type: application/json' -d '
    {
    "location": {
        "lat": 13,
        "lon": 15
    },
    "price": 100
    }
'

# insert
curl -X PUT 'localhost:9200/london_house_info/_doc/5?pretty' -H 'Content-Type: application/json' -d '
    {
    "location": {
        "lat": 11,
        "lon": 12
    },
    "price": 120
    }
'

# insert
curl -X PUT 'localhost:9200/london_house_info/_doc/6?pretty' -H 'Content-Type: application/json' -d '
    {
    "location": {
        "lat": 11,
        "lon": 12
    },
    "price": 150
    }
'

# 查看所有数据
curl 'localhost:9200/london_house_info/_search?pretty=true'

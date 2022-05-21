#!/usr/bin/env bash


# #############################################################################
# File Name   : start.sh
# Author      : DaiDai
# Mail        : daidai4269@aliyun.com
# Created Time: 日  5/15 21:56:24 2022
# #############################################################################


set -ex

# 下载并启动容器，且为 后台 自动启动
docker-compose up -d

# 查看正在运行中的容器
docker-compose ps

# 显示 running 容器
docker ps

# 停止所有up命令启动的容器
# docker-compose down

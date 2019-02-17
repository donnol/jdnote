#!/bin/bash

# 构建镜像
sudo docker build -t jdnote-server -f cmd/server/Dockerfile .

# 运行
sudo docker run -it jdnote-server 

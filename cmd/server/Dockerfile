# 构建
FROM golang:1.15.2-alpine AS builder

ARG PROJECT_PATH=Projects/jdnote

RUN mkdir -p /$PROJECT_PATH
WORKDIR /$PROJECT_PATH
COPY . .
WORKDIR /$PROJECT_PATH/cmd/server

RUN CGO_ENABLED=0 go build -o /server .

# 为所有标准证书颁发机构添加根证书, FROM: https://hub.docker.com/r/centurylink/ca-certs/
FROM centurylink/ca-certs

COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=builder /server .

# 静态文件夹
ADD cmd/server/dist ./dist

ENTRYPOINT ["./server"]

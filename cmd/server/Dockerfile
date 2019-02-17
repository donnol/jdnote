# 构建
FROM golang AS builder

ARG PROJECT_PATH=github.com/donnol/jdnote

RUN mkdir -p /go/src/$PROJECT_PATH
WORKDIR /go/src/$PROJECT_PATH
COPY ../../ .
WORKDIR /go/src/$PROJECT_PATH/cmd/server

RUN CGO_ENABLED=0 go build -o /server .

# 为所有标准证书颁发机构添加根证书, FROM: https://hub.docker.com/r/centurylink/ca-certs/
FROM centurylink/ca-certs

COPY --from=builder /usr/local/go/lib/time/zoneinfo.zip /usr/local/go/lib/time/zoneinfo.zip
COPY --from=builder /server .

ENTRYPOINT ["./server"]

# 构建
FROM golang:1.18-alpine AS builder

RUN go env -w GOPROXY=https://goproxy.io,https://goproxy.cn,direct

ARG PROJECT_PATH=Projects/jdnote

RUN mkdir -p /$PROJECT_PATH
WORKDIR /$PROJECT_PATH

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
WORKDIR /$PROJECT_PATH/cmd/server

RUN CGO_ENABLED=0 go build -o /server .

# 为所有标准证书颁发机构添加根证书, FROM: https://hub.docker.com/r/centurylink/ca-certs/
FROM centurylink/ca-certs

COPY --from=builder /server .

ENTRYPOINT ["./server"]

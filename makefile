.PHONY:server_start

unitTest=true
server_start:db_start
	go clean -cache && \
	cd cmd/server/ && \
	go install && \
	env UNIT_TEST_ENV=$(unitTest) /mnt/d/go/bin/server

db_start:
	sudo service postgresql start

dir=.
testName=TestXXX
server_unit_test:install
	go clean -testcache && \
	gotest $(dir) -v -timeout=1200s -run $(testName)

install:
	go install ./...

# 压力测试，看情况决定是否加'-k'标志
usekflag=false
ifeq ($(usekflag), true)
	kflag=-k
else
	kflag=
endif
path=/
ab:
	ab -n 1000 -c 100 $(kflag) http://localhost:8810$(path)

# 构建镜像
docker_build:
	sudo docker build -t jdnote-server -f cmd/server/Dockerfile .

# 运行
docker_run:
	sudo docker run -d --net=host --restart=unless-stopped jdnote-server 

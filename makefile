.PHONY:server_start

go_clean:
	go clean -cache

server_install:install
	cd cmd/server/ && \
	go install

unitTest=true
server_start:db_start server_install
	env PROJECT_ENV=PROJECT_ENV_DEV UNIT_TEST_ENV=$(unitTest) $(GOBIN)/server

timer_start:db_start
	cd cmd/timer/ && \
	go install && \
	env PROJECT_ENV=PROJECT_ENV_DEV $(GOBIN)/timer

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
	ab -n 1000 -c 100 $(kflag) http://localhost:8890$(path)

docker_ps:
	sudo docker ps

# 构建镜像
docker_build:
	sudo docker build -t jdnote-server -f cmd/server/Dockerfile .

# 运行
docker_run:
	sudo docker run -d --net=host --restart=unless-stopped -e="PROJECT_ENV=PROJECT_ENV_PROD" jdnote-server 

# === k8s ===
deployName=jdnote
#

# 在microk8s的registry里构建镜像
registry_docker_build:
	sudo docker build -t localhost:32000/jdnote-server -f cmd/server/Dockerfile .

# k8s create deploy
kubectl_create_deploy:
	sudo microk8s.kubectl create deployment $(deployName) --image=localhost:32000/jdnote-server

# k8s delete deploy
kubectl_delete_deploy:
	sudo microk8s.kubectl delete deployment $(deployName)

# 获取k8s deploy配置
kubectl_get_deploy_config:
	sudo microk8s.kubectl get deploy $(deployName) -o yaml > zeus/data/deployment.yml

# k8s scale deploy
kubectl_scale_deploy:
	sudo microk8s.kubectl scale deployment $(deployName) --replicas=2

# k8s deploy service
kubectl_expose_service:
	# kubectl expose deployment jdnote-server --type=NodePort --name=jdnote-server-service --port=8890
	sudo microk8s.kubectl expose deployment $(deployName) --type=NodePort --port=80 --name=jdnote-service

# k8s delete service
kubectl_delete_service:
	sudo microk8s.kubectl delete services/jdnote-service

# k8s get all namespaces
kubectl_get_all_namespaces:
	sudo microk8s.kubectl get all --all-namespaces | less

podName=jdnote-76db6f897c-kgvvq

# k8s pod log
kubectl_get_pod_logs:
	sudo microk8s.kubectl logs $(podName) -p

yamlFile=zeus/data/postgresql.yaml

# k8s apply -f
kubectl_apply_byfile:
	sudo microk8s.kubectl apply -f $(yamlFile)


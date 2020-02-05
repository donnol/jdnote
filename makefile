.PHONY:server_start

unitTest=true
server_start:
	go clean -cache && \
	cd cmd/server/ && \
	go install && \
	env UNIT_TEST_ENV=$(unitTest) /mnt/d/go/bin/server

db_start:
	sudo service postgresql start

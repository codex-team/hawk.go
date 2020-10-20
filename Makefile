test: SHELL:=/bin/bash
test:
	set -xeuo pipefail;\
	cd tests;\
	mkdir client/hawk;\
	cp ../*go* client/hawk/;\
	cp -r client/hawk server/;\
	docker-compose -f docker-compose-test.yaml up --build -d;\
	test_status_code=0;\
	docker-compose -f docker-compose-test.yaml logs > logs.txt;\
	docker-compose -f docker-compose-test.yaml run client ./client || test_status_code=$$?;\
	docker-compose -f docker-compose-test.yaml down;\
	if grep -q fail logs.txt; then \
		cat logs.txt | grep fail;\
		test_status_code=1;\
	fi ;\
	rm -rf client/hawk server/hawk;\
	exit $$test_status_code

ut:
	go test -v -count=1 -race -gcflags=-l -timeout=30s .

.PHONY: test ut

ut:
	go test -v -count=1 -race -gcflags=-l -timeout=20s .

.PHONY: ut

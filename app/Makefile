.PHONY: install-watcher
install-watcher:
	go get github.com/canthefason/go-watcher
	go install github.com/canthefason/go-watcher/cmd/watcher

.PHONY: install
install:
	go get -u github.com/gorilla/mux
	go get -u github.com/lib/pq
	go get -u html/template
	go get -u gopkg.in/oleiade/reflections.v1

.PHONY: run
run: install install-watcher
	watcher

.PHONY: test
test: install
	go test


# go_install
# go_compile

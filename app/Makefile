.PHONY: install-watcher
install-watcher:
	go get github.com/canthefason/go-watcher
	go install github.com/canthefason/go-watcher/cmd/watcher

.PHONY: install
install: install-watcher
	go get -u github.com/gorilla/mux

.PHONY: run
run: install
	watcher

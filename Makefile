PACKAGE  = github.com/thejasbabu/keep-it-dry
BASE     = $(GOPATH)/src/$(PACKAGE)

.PHONY: test
test: | $(BASE)
	cd $(BASE) && dep ensure && go test ./...
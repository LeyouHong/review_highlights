PACKAGES = $(shell go list ./... | grep -v '/vendor/')

build:
	go build -o rh github.com/LeyouHong/review_highlights

clean:
	rm rh

.PHONY: test
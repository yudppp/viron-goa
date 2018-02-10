run:
	go run *.go
gen:
	goagen main -d github.com/yudppp/viron-goa/design
	goagen app -d github.com/yudppp/viron-goa/design
	goagen swagger -d github.com/yudppp/viron-goa/design
install:
	git clone --depth 1 --branch gh-pages --single-branch git@github.com:cam-inc/viron.git viron-client
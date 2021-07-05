.PHONY: all

all: clean build

build:
	@echo =================================================================
	@echo building...
	@echo =================================================================
	GOOS=windows GOARCH=amd64 $(MAKE) proxy

	@echo =================================================================
	@echo testing...
	@echo =================================================================
	cd main; go test

	@echo =================================================================
	@echo starting server...
	@echo =================================================================
	@read -p "Enter port number [default : 8080]:" module; \
	build/proxy $${module};
	

proxy: ./main/proxy.go
	go build -o ./build/proxy ./main



.PHONY: clean
clean: 
	rm -fr ./build

.PHONY: run
run:
	build/proxy
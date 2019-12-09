build_run:
	docker build -t lagashmx/go-workshop:dev .
	docker run -it lagashmx/go-workshop:dev

build:
	docker build -t lagashmx/go-workshop:dev .

release:
	docker build -t lagashmx/go-workshop:release .

run:
	docker run -it lagashmx/go-workshop:dev

gointo:
	docker run -it lagashmx/go-workshop:dev /bin/bash

test:
	docker run -it lagashmx/go-workshop:dev go test

build_debug:
	docker build --target debug -t lagashmx/go-workshop:debug .
debug:		
	docker build --target debug -t lagashmx/go-workshop:debug .
	docker run -p 40000:40000 -p 1541:1541 --security-opt=seccomp:unconfined lagashmx/go-workshop:debug
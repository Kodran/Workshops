build_run:
		docker build -t lagashmx/go-workshop:dev .
		docker run -it -p 8080:8081 lagashmx/go-workshop:dev
build:
		docker build -t lagashmx/go-workshop:dev .
run:		
		docker run -it -p 8080:8081 lagashmx/go-workshop:dev
gointo:
		docker run -it lagashmx/go-workshop:dev /bin/bash
test:
		docker run -it lagashmx/go-workshop:dev go test

build_debug:
	docker build -t lagashmx/app-debug:dev .
debug:	
	docker run -p 40000:40000 -p 1541:1541 --security-opt=seccomp:unconfined lagashmx/app-debug:dev
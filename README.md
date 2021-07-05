


## Objective
Create a proxy server in Golang.

## Constraints
- Proxy server should add a key and its value to request header of server API.
- Client shoul receive the response returned by actual server.

## How to build and use 
There are two ways to build it

* [Automated using Makefile](#Automated)
* [Manual](#Manual)

## Automated
- Go to root directory (proxy)
- make
- It will build the project in build directory, perform testing and ask for the port number to start proxy server.
- If no port is entered server will start on defaul port 8080.
- Use http client to test end point : http://127.0.0.1:PORT_NUMBER/test 


## Manual
- Go to directory proxy/main
- Run tests : go test -v
- Start server : go run proxy.go PORT_NUMBER (eg. 3000)
- If no port is entered server will start on defaul port 8080.
- Use http client to test end point : http://127.0.0.1:PORT_NUMBER/test 

<!-- TABLE OF CONTENTS -->
## Directory structure

* [main](#main)
  * [proxy.go](#proxy)
  * [proxy_test.go](#proxy_test)
* [go.mod](#mod)
* [Makefile](#Makefile)
* [README.md](#README)
* [build](#build)
  * [proxy](#proxybin)


## main
Main directory that contains program logic and testing logic.

#### proxy
proxy.go file contains the proxy server logic.

#### proxy_test
proxy_test.go file contains the test cases.

## mod
go.mod file contains all the external module dependencies.

## Makefile
Instructions to build, test and start proxy server.

## build
This directory is generated when we build project and contains the binary code.

#### proxybin
proxy is the executale file generated after building the project.


## Design and features
Overall design of proxy server is like if a client will send a request to proxy server, it will add an secret key to it and send to actual server.
Receives response and send it to client.

Some of the design points are listed below-

- Proxy servers port is configurable. If no port is provided server will start on port 8080.
- Each request from client is handled by Handler function.
- We have given single API end point (/test) of actual server, so  any other end point will return 404.
- If server URL is provided we can pass all the requests to actual server.
- In proxy server create and initialize a request parameter with necessory fields. [host, scheme, path, request URI, client IP and Key]
- Make request to server with this request object.
- Copy response code and body of server response to clients responce.
- Copy all headers of server response to clients responce.
- If server sends streaming data, proxy should flush this data in reguler interval until the complete data is received.
- If server sends trailer, proxy needs to read its key and send it before setting status code, after that pass the trailer values after copying body. 

## Limitation/Improvement
- Secret key is hard coded in logic. This is not a good practice. We should store Key in environment file.
- We can provide a cacheing layer with this proxy server to to lower the load of server.
- We can use secure HTTP between client and proxy server to make request more secure.

<!-- LICENSE -->
## License

Distributed under the MIT License. 

<!-- CONTACT -->
## Contact

Name: Dinesh Dev Pandey

LinkedIn : [https://www.linkedin.com/in/dinesh-dev-pandey-51317229]

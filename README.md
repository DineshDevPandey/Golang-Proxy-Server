


## Objective
Create a proxy server in Golang.

## Constraints
- Proxy server should add a key and its value to request header of server request.
- Client should receive the response returned by actual server.

## How to build and use 
There are two ways to build it

* [Automated using Makefile](#Automated)
* [Manual](#Manual)

## Automated
- Go to the root directory (proxy)
- Run command **make**
- It will build the project in **build** directory, perform testing and ask for the port number to start the proxy server.
- If no port is entered, the server will be started on default port 8080.
- Use http client to test end point : http://127.0.0.1:PORT_NUMBER/test


## Manual
- Go to directory proxy/main
- Run tests : go test -v
- Start server : go run proxy.go PORT_NUMBER (eg. 3000)
- If no port is entered, the server will be started on default port 8080.
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
Main directory that contains program logic and testing logic.

#### proxy
proxy.go file contains the proxy server logic.

#### proxy_test
proxy_test.go file contains the test cases.

## mod
go.mod file contains all the external module dependencies.

## Makefile
Instructions to build, test and start proxy server.

## build
This directory is generated when we build project and contains the binary code.

#### proxybin
proxy is the executable file generated after building the project.

## Design and features
Overall design of proxy server is like if a client will send a request to proxy server, it will add a secret key to it and send it to actual server.Receives response and sends it to client.

Some of the design points are listed below-

- Proxy servers port is configurable. If no port is provided, the server will be started on port 8080.
- Each request from the client is handled by the Handler function.
- We have given a single API endpoint (/test) of the actual server, so any other endpoint will return 404.
- If the server URL is provided we can pass all the requests to the actual server.
- In proxy server create and initialize a request parameter with necessary fields. [host, scheme, path, request URI, client IP and Key]
- Make a request to the server with this request object.
- Copy response status code and body of server response to client's response.
- Copy all headers of server response to client's response.
- If the server sends streaming data, proxy should flush this data in a regular interval until the complete data is received.
- If the server sends a trailer, proxy needs to read its key and send it before setting status code, after that pass the trailer values after copying the body.


## TestCases
There are two main functions. Following test cases have been added -
**Handler Function**

- test endpoint check                     : Test to check /test URL returns correct response 
- other endpoint check                    : Test to check other endpoints returns 404
- empty server URL check                  : Test to check empty server address/url returns 500
- remoteAddr without port check           : Test to check remoteAddr(without port) returns 500 
- response header secret KEY/VALUE check  : Test to check proxy response contains added secret KEY and VALUE

**Port Function**

- user defined port                       : Test for check it returns correct port provided by user 
- default port                            : Test for check it returns default port if user doesn't provide port 
- non integer port                        : Test for check it returns default port if user provides non numeric port 


## Limitations / Improvements
- Secret key is hard coded in logic. This is not a good practice. We should store the Key in the environment file.
- We can provide a caching layer with this proxy server to lower the load of the server.
- We can use secure HTTP between client and proxy server to make requests more secure.

<!-- LICENSE -->
## License

Distributed under the MIT License. 

<!-- CONTACT -->
## Contact

Name: Dinesh Dev Pandey

LinkedIn : [https://www.linkedin.com/in/dinesh-dev-pandey-51317229]

Project Link: [https://github.com/DineshDevPandey/Golang-Proxy-Server](https://github.com/DineshDevPandey/Golang-Proxy-Server.git)

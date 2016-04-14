# hashService
A web service thats accepts POST requests and returns a the SHA513 hash of a string.


## Installing
go get github.com/sullirobert/hashService


## Configuring Port
uses port 9999 by default
flags:
    -port sets the listen port


## Example
go install github.com/sullirobert/hashService
hashService -port 8080

send a request:
curl --data "password=angryMonkey" http://localhost:8080

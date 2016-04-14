# hashService
A web service thats accepts POST requests and returns the base64 encoded SHA512 hash of a string.


## Installing
go get github.com/sullirobert/hashService


## Configuring Port
uses port 9999 by default<br/>
flags:<br/>
    -port sets the listen port<br/>


## Example
go install github.com/sullirobert/hashService<br/>
hashService -port 8080<br/>
<br/>
send a request:<br/>
curl --data "password=angryMonkey" http://localhost:8080

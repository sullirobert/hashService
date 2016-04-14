# hashService
service to return a password hash

install:
go get github.com/sullirobert/hashService

uses port 9999 by default

flags:
    -port sets the listen port


example:
go install github.com/sullirobert/hashService
hashService -port 8080

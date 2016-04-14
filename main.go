package main

import (
	"crypto/sha512"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 9999, "set listening port")
	flag.Parse()
	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: hashStringhandler{},
	}
	fmt.Printf("Listening on port %d\n", port)
	log.Fatal(s.ListenAndServe())
}

type hashStringhandler struct{}

func (h hashStringhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("request recieved")
	time.Sleep(time.Second * 5)
	pw := r.FormValue("password")
	if pw == "" {
		w.WriteHeader(400)
		w.Write([]byte(`empty string error`))
		return
	}
	hashedPwsha := EncodedHash([]byte(pw))
	log.Printf("returning string: %s", hashedPwsha)
	w.Write([]byte(hashedPwsha))
}

// EncodedHash returns encoded SHA512 hash of a string
func EncodedHash(value []byte) string {
	hasher := sha512.New()
	hasher.Write(value)
	return base64.StdEncoding.EncodeToString(hasher.Sum(nil))

}

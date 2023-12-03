package handlers
import (
	"io/ioutil"
	"net/http"
	"log"
	"fmt"
)
//log.Logger is from package log which takes in a struct type.

type Hello struct {
l *log.Logger    //a memory address, for dependency injection.
}
//this function points to the Hello Struct and creates a new instance that will be passed into that same address of l.
func NewHello(l *log.Logger) *Hello{
	return &Hello{l} //
	}

//define a method which satisfies the handler interface.
func(h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
    h.l.Println("endpointhit")
	d, err := ioutil.ReadAll(r.Body) //store the reads from the body to the variable d.	
	if err != nil {
		http.Error(rw, "oops", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(rw, "hello %s", d)

	
}


package main
import (
	"net/http"
	"log"
	"fmt"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/goodbye", goodBye)
	http.ListenAndServe(":10000", nil)
}
func goodBye(w http.ResponseWriter, r *http.Request) {
	
	d, err := ioutil.ReadAll(r.Body) //store the reads from the body to the variable d.	
	if err != nil {
		http.Error(w, "oops", http.StatusBadRequest)
		return
	}
	fmt.Println("hello world")
	log.Println("endpointhit")
	fmt.Fprintf(w, "data %s\n", d)
}
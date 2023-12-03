package main
import (
	"net/http"
	"log"
	"os"
	"github.com/marviigrey/golang-microservice-projects/web-server/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", helloHandler)

	
http.ListenAndServe(":8087", sm)
}

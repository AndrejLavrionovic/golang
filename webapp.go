package GolangTest

import(
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main(){
	http.Handle("/", handler)
	http.ListenAndServe(":9999", nil)
}
package main
import(
	"net/http"
	"fmt"
	"strings"
)


type MyServer struct {

}
func (p MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	
	router := http.NewServeMux()
	router.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		content := "AHHH"
		s := fmt.Sprintf("<h1> %s </h1>", content)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, s)
	}))
	router.Handle("/players/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		player := strings.TrimPrefix(r.URL.Path, "/players/")
		s := fmt.Sprintf("<h1> %s </h1>", player)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, s)
	}))
	
	router.ServeHTTP(w, r)
}
func main(){
	server := MyServer{}
	http.ListenAndServe(":5000", server)
}
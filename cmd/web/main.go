package main
import (
	"flag"
"log"
"net/http"
)
func main() {
	// flag variables
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	
	mux := http.NewServeMux()
	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.

	
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// Register the other application routes as normal.
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Printf("Starting server on %s",*addr)
	err := http.ListenAndServe(*addr, mux)
log.Fatal(err)
}
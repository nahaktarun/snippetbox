package main
import (
	"flag"
"log"
"net/http"
"os"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}


func main() {
	// flag variables
	addr := flag.String("addr", ":4000", "HTTP network address")

	flag.Parse()

	
	
	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.

	infoLog := log.New(os.Stdout, "/INFO\t",log.Ldate|log.Ltime)

	errorLog := log.New(os.Stderr, "ERROR\t",log.Ldate|log.Ltime|log.Lshortfile)
	

	app := &application{
		errorLog: errorLog,
		infoLog: infoLog,
	}
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// Register the other application routes as normal.
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)


	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}

	infoLog.Printf("Starting server on %s",*addr)
	err := srv.ListenAndServe()
errorLog.Fatal(err)
}
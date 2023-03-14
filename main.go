package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){

	if r.URL.Path != "/"{
		http.NotFound(w,r)
		return
	}

	w.Write([]byte("Hello from snippetbox"))
}



// Add a SnippetView handler function
func snippetView(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Display a specific snippet"))
}

// add a snippetCreate Handler function
func snippetCreate(w http.ResponseWriter, r *http.Request){
	
	if r.Method != http.MethodPost{
		w.Header().Set("Allow", "POST")
		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed"))
		http.Error(w, "Method not found",http.StatusMethodNotAllowed)
		return
	}
	
	w.Write([]byte("Create a new Snippet"))

}
func main(){

	mux := http.NewServeMux()
	mux.HandleFunc("/",home)
	mux.HandleFunc("/snippet/view",snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Print("Starting server on: 4000")
	err := http.ListenAndServe(":4000",mux)

	log.Fatal(err)

}
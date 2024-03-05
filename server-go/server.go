package main

import(
	"net/http"
)

func main(){
	http.HandleFunc("/", Hello)
	//name:= "edson"
	//age:= 43
	//fmt.Fprintf( "Hello, I'm %s. I'm %s", name, age)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request){
    //name := os.Getenv("NAME")
    // age := os.Getenv("AGE")
	w.Write([]byte("teste"))
}
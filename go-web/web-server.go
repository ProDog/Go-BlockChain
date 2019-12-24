package main

import(
	"fmt"
	"net/http"
	"strings"
	"log"
	"html/template"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v := range r.Form {
		fmt.Println("key:",k)
		fmt.Println("val:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"Hello grip!")
}

func testFunc(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()

	fmt.Fprintf(w,"test!")
}

func login(w http.ResponseWriter,r *http.Request) {
	fmt.Println("method:",r.Method)
	if r.Method=="GET" {
		t,_:=template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w,nil))
	} else {
		r.ParseForm()
		
		fmt.Println("username:",r.Form["username"])
		fmt.Println("pwd:",r.Form["password"])
	}
}

func main() {
	http.HandleFunc("/",sayHelloName)
	http.HandleFunc("/test",testFunc)
	http.HandleFunc("/login",login)
	err:=http.ListenAndServe(":9090",nil)
	if err!=nil{
		log.Fatal("ListenAndServe: ",err)
	}

	// mux :=&MyMux{}
	// http.ListenAndServe(":9090",mux)
}


type MyMux struct {
}

func (p *MyMux) ServeHTTP(w http.ResponseWriter,r *http.Request) {
	if r.URL.Path =="/" {
		sayHello(w,r)
		return
	}

	http.NotFound(w,r)
	return
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w,"Hello my route!")
}


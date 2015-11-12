package blog

import (
	"net/http"
	"html/template"
	"log"
)

//注册actions
func init(){
	http.HandleFunc("/",index)
}

func index(w http.ResponseWriter,r *http.Request){
	
	if r.URL.Path == "/" {
		return
	}
	
	//not found
	t,err := template.ParseFiles("blog/templates/404.html")
	if err!=nil{
		log.Fatal(err)
	}
	t.Execute(w,nil)
}

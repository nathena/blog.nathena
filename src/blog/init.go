package blog

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func init(){
	
	initEnv()
	
	if strings.Trim(DefaultEnv.Static," ")  == "on" {
		http.Handle("/static/",http.FileServer(http.Dir("blog")))
		println(" set /static/ FileServer ")
	}
}

type env struct {
	Port int
	Static string
}

var DefaultEnv env

func initEnv(){
	
	data,err := ioutil.ReadFile("config")
	if err!=nil {
		println("read data error ",err)
	}
	
	err = json.Unmarshal(data,&DefaultEnv)
	if err != nil {  
        println("error in translating,", err)  
        return
    }
}


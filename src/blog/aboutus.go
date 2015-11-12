package blog

import (
	"net/http"
	"utils"
)

//注册actions
func init(){
	http.HandleFunc("/aboutus",aboutus)
}


func aboutus(rep http.ResponseWriter,req *http.Request){
	rep.Write(utils.String2Byte("aboutus"))
}





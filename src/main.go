package main

import (
	_ "blog"
	_ "net/http"
	"net"
	_ "strings"
	//"bytes"
	"strconv"
)

func main(){
	
	//println(blog.DefaultEnv.Port)
	//http.ListenAndServe(":8181",nil)
	conn,err := net.Dial("tcp","baidu.com:80")
	defer conn.Close()
	if err!=nil {
		println("error : ",err)
		return
	}
	
	println("ip = ",conn.LocalAddr().String())
	println("ip = ",conn.RemoteAddr().String())
	
	
	
	//data := bytes.NewBufferString("挖掘机哪家强")
	//for _,v := range data.Bytes() {
		//print(strconv.FormatInt(int64(v),16)," ")
	//}
	val,err := strconv.ParseInt("e6",16,8)
	
	println(val)
}
package main

import (
	"controllers"
	"net/http"
)

var resp=make(chan string)
var out=make(chan string)




func main() {
	//code:="#include<stdio.h>\n int main()\n{printf(\"Hello Go Lang\");return 0;}"
	//start := time.Now()
	//var wg sync.WaitGroup
	//for i:=0;i<1000 ;i++ {
	//	wg.Add(1)
	//	go func(wg*sync.WaitGroup,i int) {
	//		go saveToFile(code)
	//		fName:=<-resp
	//		go runCode(wg,fName)
	//		ret := <-out
	//		fmt.Println(ret)
	//	}(&wg,i)
	//}
	//wg.Wait()
	//elapsed := time.Since(start)
	//fmt.Println(elapsed)
	http.HandleFunc("/execute",controllers.CodeHandler)
	http.ListenAndServe(":8083",nil)
}
package controllers

import (
	"encoding/json"
	"fmt"
	"models"
	"net/http"
	"sync"
	u "utils"
)

func CodeHandler(w http.ResponseWriter,r *http.Request) {
	var code =  models.Code{}
	var resp = make(chan string)
	var out = make(chan string)
	json.NewDecoder(r.Body).Decode(&code)
		var wg sync.WaitGroup
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			go code.SaveToFile(resp)
			fName:=<-resp
			go code.RunCode(fName,out,r.Context())
			ret := <-out
			fmt.Println(ret)
			msg := u.Message(true, "success")
			msg["output"]=ret
			u.Respond(w,msg)
		}(&wg)
		wg.Wait()
}


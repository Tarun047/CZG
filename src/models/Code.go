package models

import (
	"bytes"
	"context"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Code struct {
	Source string
	Language string
}

func (code Code)SaveToFile(resp chan string) {

	f, err := ioutil.TempFile(".",strconv.FormatInt(time.Now().Unix(),10)+"*"+".c")
	if err!=nil{
		panic(err.Error())
	}
	defer f.Close()
	f.WriteString(code.Source)
	resp<-f.Name()
}

func (code Code) RunCode(fName string, out chan string, rctx context.Context) {
	execName:=strings.TrimSuffix(fName, filepath.Ext(fName))
	cmd1:=exec.Command("gcc",fName,"-o",execName)
	err1:=cmd1.Run()
	if err1!=nil{
		panic(err1.Error())
	}
	ctx,cancel := context.WithTimeout(rctx,time.Second*2)
	defer cancel()
	cmd:=exec.CommandContext(ctx,"./"+execName)
	var x bytes.Buffer
	cmd.Stdout = &x
	if err := cmd.Start(); err != nil {
		if cmd.Process != nil {
			cmd.Process.Kill()
		}
		panic(err.Error())
	}
	err2:=cmd.Wait()
	if err2!=nil{
		panic(err2.Error())
	}
	out<-x.String()
	defer os.Remove(fName)
	defer os.Remove(execName)
}





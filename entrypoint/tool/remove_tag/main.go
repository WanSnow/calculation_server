package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var fileName = "F:\\File\\学校\\毕业设计\\calculation_server\\model\\mission\\Mission.pb.go"

func main() {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
		return
	}
	data := strings.ReplaceAll(string(file), ",omitempty", "")
	file = []byte(data)
	err = ioutil.WriteFile(fileName, file, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("end")
}

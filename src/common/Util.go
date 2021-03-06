package common

import (
	"os"
	"io/ioutil"
	"strconv"
	"github.com/gin-gonic/gin/json"
	"fmt"
)

func ReadConfFile(fileName string) []byte {
	fi, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	recover()
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	return fd
}

func WriteConfFile(fileName string, confStr string) {
	fi, err := os.OpenFile(fileName, os.O_CREATE, 0)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	ioutil.WriteFile(fileName, []byte(confStr), 0)
}

func AtoI(source string) int {
	result, err := strconv.Atoi(source)
	if err != nil {
		return 0
	} else {
		return result
	}
}

func ToJson(obj interface{}) string {
	ret, err := json.Marshal(obj)
	if err != nil {
		fmt.Println("err: ", err)
		recover()
		return "{}"
	}
	return string(ret)
}


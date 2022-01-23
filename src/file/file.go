package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	legacyPath := filepath.Join("src", "testnet")
	log.Println(legacyPath)
}

func echoFileInfo() {
	// os.Stat 返回文件信息 如果文件不存在则返回错误 The system cannot find the path specified
	fileinfo, err := os.Stat("src/testnet")
	if fileinfo != nil {
		log.Panicln(fileinfo)
	}
	if err != nil {
		log.Println(err)
	}
}

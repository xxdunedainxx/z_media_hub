package util

import (
		"log"
    "os"
)


func GetCwd() string {
	path, pattherr := os.Getwd()
	if pattherr != nil {
	    log.Println(pattherr)
	}
	return path
}
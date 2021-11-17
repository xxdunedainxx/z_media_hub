//https://www.honeybadger.io/blog/golang-logging/

package util

import (
	"log"
)

func Setup() {
	log.Println("Start project...")
	initLogger()
}
package utils

import (
	"corpotecht_client/config"
	"log"
)

func DebugLog(payload ...any) {
	if config.Logging {
		log.Println(payload...)
	}
}

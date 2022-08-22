package utils

func HandleError(e error) {
	if e != nil {
		DebugLog(e.Error())
		panic(e)
	}
}

func RecoverFromCrash() {
	recover()
}

package common

func PanicErr(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckErrAndCallback(err error, err_call func()) {
	if err != nil {
		err_call()
	}
}

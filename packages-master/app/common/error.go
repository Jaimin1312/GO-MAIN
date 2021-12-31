package common

//CheckError is check error
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

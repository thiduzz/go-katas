package utils

func HasError(e error) {
	if e != nil {
		panic(e)
	}
}

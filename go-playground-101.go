package goplayground

func testLog() string {
	return "> test log"
}

func Hello(name string) string {
	message := testLog()
	return message
}

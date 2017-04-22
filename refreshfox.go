package main

func main() {
	var caller Caller = telnet.StandardCaller

	//@TOOD: replace "example.net:5555" with address you want to connect to.
	telnet.DialToAndCall("example.net:5555", caller)
}

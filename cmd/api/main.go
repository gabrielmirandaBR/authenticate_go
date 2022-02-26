package main

func main() {
	server := MustNewServer()
	router := server.Listen()

	router.Run()
}

package main

func main() {
	// get from config
	HOST := "localhost"
	PORT := "8554"

	server := Server{Host: HOST, Port: PORT}
	server.Start()
}

package main

func main() {
	// Loads configuration from ../config/default.json
	configuration := LoadConfiguration()

	server := NewServer(configuration)
	server.Start()
}

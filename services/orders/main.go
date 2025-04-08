package main

func main() {
	gRPCServer := NewGRPCServer(":9000")
	gRPCServer.Run()
}

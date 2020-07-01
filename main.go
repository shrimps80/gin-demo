package main

func main() {
	r := initRouter()
	
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

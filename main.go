package main

import (
	"os"

	// "cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	name := os.Getenv("NAME")
	hello := "Hello " + name + " and the World!"

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": hello,
		})
	})

	r.Run()

	// http.HandleFunc("/", indexHandler)

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// 	log.Printf("Defaulting to port %s", port)
	// }

	// log.Printf("Listening on port %s", port)
	// if err := http.ListenAndServe(":"+port, nil); err != nil {
	// 	log.Fatal(err)
	// }
}

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.NotFound(w, r)
// 		return
// 	}

// 	name := os.Getenv("NAME")
// 	hello := "Hello " + name + " and the World"
// 	fmt.Fprint(w, hello)
// }

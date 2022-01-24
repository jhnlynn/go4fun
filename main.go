package main

import (
	"log"
	"new-go/routers"
	"os"
)

func main()  {
	router := routers.Routers()

	i := os.Getenv("en")

	log.Println(i)

	_ = router.Run()
}

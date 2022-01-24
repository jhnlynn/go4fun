package main

import "new-go/routers"

func main()  {
	router := routers.Routers()

	_ = router.Run(":4000")
}

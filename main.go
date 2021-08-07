package main

import (
	"fmt"

	route "github.com/StayAlert/blogByF/backend/router"
)

func main() {
	fmt.Println("Hello Start Server ...")

	route.Router()
}

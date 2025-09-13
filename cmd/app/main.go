package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tierthree/go-finance/internal/api/routes"
)

type test struct {
	ID string
}

var ids = []test{
	{ID: "1"},
}
var tracker = 2

func testValues(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ids)
}

func setValues(c *gin.Context) {
	ids = append(ids, test{ID: string(tracker)})
	tracker = tracker + 1
}

func main() {
	router := gin.Default()
	fmt.Sprintf("Type is: %T\n", router)
	router.GET("/test", testValues)
	router.POST("/addtest", setValues)

	router.Run("localhost:8080")
}

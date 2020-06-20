package main

import (
	"fmt"
	"grpc-demo/calculatorpb"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:3030", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := calculatorpb.NewCalculatorClient(connection)

	g := gin.Default()
	g.GET("calculate/:val1/:op/:val2/", func(ctx *gin.Context) {
		val1, err := strconv.ParseFloat(ctx.Param("val1"), 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Paramater val1"})
			return
		}

		val2, err := strconv.ParseFloat(ctx.Param("val2"), 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Request Paramater val2"})
			return
		}

		req := &calculatorpb.Request{Value1: float32(val1), Value2: float32(val2), Operation: ctx.Param("op")}
		response, err := client.Calculate(ctx, req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(response.Result),
		})
	})

	runErr := g.Run(":9002")
	if runErr != nil {
		fmt.Print("Error running server")
	}
}

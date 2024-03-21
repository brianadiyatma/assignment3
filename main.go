package main

import (
	"encoding/json"
	"fmt"

	"math/rand/v2"
	"net/http"
	"os"
	"third-assignment/model"

	"github.com/gin-gonic/gin"
)

func main() {

	///mengarahkan view
	ginEngine := gin.Default()
	ginEngine.LoadHTMLGlob("views/*")
	ginEngine.GET("/", func(ctx *gin.Context) {
		report := model.StatusReport{}
		report.Status.Water = rand.IntN(100-1) + 1
		report.Status.Wind = rand.IntN(100-1) + 1

		jsonData, err := json.Marshal(report)
		if err != nil {
			fmt.Println("Error marshalling to JSON:", err)
			return
		}

		filePath := "status_report.txt"

		err = os.WriteFile(filePath, jsonData, 0644)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}

		fmt.Println("File written successfully")

		var status string
		if report.Status.Water < 5 {
			status = "aman"
		} else if report.Status.Water >= 6 && report.Status.Water <= 8 {
			status = "siaga"
		} else if report.Status.Water > 8 {
			status = "bahaya"
		}

		if report.Status.Wind < 6 {
			if status != "siaga" && status != "bahaya" {
				status = "aman"
			}
		} else if report.Status.Wind >= 7 && report.Status.Wind <= 15 {
			if status != "bahaya" {
				status = "siaga"
			}
		} else if report.Status.Wind > 15 {
			status = "bahaya"
		}

		ctx.HTML(http.StatusOK, "index.tmpl", gin.H{
			"water":  report.Status.Water,
			"wind":   report.Status.Wind,
			"status": status,
		})
	})
	ginEngine.Run(fmt.Sprintf("%s:%d", "localhost", 8000))
}

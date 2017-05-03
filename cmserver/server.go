package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jcelliott/lumber"
	"github.com/schollz/closestmatch"
	"gopkg.in/urfave/cli.v1"
)

var version string
var log *lumber.ConsoleLogger
var cm *closestmatch.ClosestMatch

func main() {

	app := cli.NewApp()
	app.Name = "cmserver"
	app.Usage = "fancy server for connecting to a closestmatch db"
	app.Version = version
	app.Compiled = time.Now()
	app.Action = func(c *cli.Context) error {
		listfile := c.GlobalString("list")
		verbose := c.GlobalBool("debug")
		port := c.GlobalString("port")

		if verbose {
			log = lumber.NewConsoleLogger(lumber.TRACE)
		} else {
			log = lumber.NewConsoleLogger(lumber.WARN)
		}

		log.Info("Loading closestmatch...")
		var errcm error
		cm, errcm = closestmatch.Load(listfile + ".cm")
		if errcm != nil {
			log.Warn(errcm.Error())
			log.Info("...loading data file...")
			var intArray []int
			for _, intStr := range strings.Split(c.GlobalString("bags"), ",") {
				intInt, _ := strconv.Atoi(intStr)
				intArray = append(intArray, intInt)
			}
			keys, err := ioutil.ReadFile(listfile)
			if err != nil {
				log.Error(err.Error())
				return err
			}
			log.Info("...computing cm...")
			cm = closestmatch.New(strings.Split(string(keys), "\n"), intArray)
			log.Info("...computed.")
			//log.Info("Saving...")
			//cm.Save(listfile + ".cm")
			//log.Info("...saving.")
		}

		startTime := time.Now()

		gin.SetMode(gin.ReleaseMode)
		r := gin.Default()
		r.GET("/v1/api", func(c *gin.Context) {
			c.String(200, `

				// Get map of buckets and the number of keys in each
				GET /uptime
	`)
		})
		r.GET("/uptime", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"uptime": time.Since(startTime).String(),
			})
		})
		r.POST("/match", handleMatch)

		fmt.Printf("cmserver (v.%s) running on :%s\n", version, port)
		r.Run(":" + port) // listen and serve on 0.0.0.0:8080
		return nil
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "port, p",
			Value: "8051",
			Usage: "port to use to listen",
		},
		cli.StringFlag{
			Name:  "list,l",
			Value: "",
			Usage: "list of phrases to load into closestmatch",
		},
		cli.StringFlag{
			Name:  "bags,b",
			Value: "2,3",
			Usage: "comma separated bags",
		},
		cli.BoolFlag{
			Name:  "debug,d",
			Usage: "turn on debug mode",
		},
	}
	app.Run(os.Args)

}

// test with
// http POST localhost:8051/match s='The War of the Worlds by HG Wells'
func handleMatch(c *gin.Context) {
	type QueryJSON struct {
		SearchString string `json:"s"`
		N            int    `json:"n"`
	}
	var json QueryJSON
	if c.BindJSON(&json) != nil {
		log.Trace("Got %v", json)
		c.String(http.StatusBadRequest, "Must provide search_string")
		return
	}
	log.Trace("Got %v", json)
	if json.N == 0 {
		c.JSON(http.StatusOK, gin.H{"r": cm.Closest(json.SearchString)})
	} else {
		c.JSON(http.StatusOK, gin.H{"r": cm.ClosestN(json.SearchString, json.N)})
	}
}

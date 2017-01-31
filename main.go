package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/dtan4/stimer-dashboard/systemd"
	"github.com/gin-gonic/gin"
)

func main() {
	conn, err := systemd.NewConn()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer conn.Close()

	client := systemd.NewClient(conn)

	r := gin.Default()

	tmpl, err := template.New("stimer-dashboard").Funcs(templateHelpers).ParseGlob("views/*")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	r.SetHTMLTemplate(tmpl)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "stimer-dashboard")
	})

	r.GET("/timers", func(c *gin.Context) {
		timers, err := client.ListTimers()
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}

		c.HTML(http.StatusOK, "timers.tmpl", gin.H{
			"timers": timers,
		})
	})

	r.Run(":8080")
}

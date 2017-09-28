package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
    r.StaticFS("/files", http.Dir("files"))

	templateBase := "templates"
	templateFileNames := []string{
		fmt.Sprintf("%s/%s", templateBase, "headerpart.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "navigation_bar.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "functions.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "functions_nodes.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "edit_node.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "index.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "footerpart.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "nodes/liquid_metals.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "nodes/physical.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "nodes/physical_diagram.tmpl"),
		fmt.Sprintf("%s/%s", templateBase, "nodes/networks.tmpl"),
	}
	r.LoadHTMLFiles(templateFileNames...)

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{"category": "top"})
	})
	nodes := r.Group("/nodes")
	{
		nodes.GET("/liquid_metals", func(c *gin.Context) {
			c.HTML(http.StatusOK, "nodes/liquid_metals.tmpl", gin.H{"category": "nodes.liquid_metals"})
		})

		nodes.GET("/physical", func(c *gin.Context) {
			c.HTML(http.StatusOK, "nodes/physical.tmpl", gin.H{"category": "nodes.physical"})
		})

		nodes.GET("/physical_diagram", func(c *gin.Context) {
			c.HTML(http.StatusOK, "nodes/physical_diagram.tmpl", gin.H{"category": "nodes.physical_diagram"})
		})

		nodes.GET("/networks", func(c *gin.Context) {
			c.HTML(http.StatusOK, "nodes/networks.tmpl", gin.H{"category": "nodes.networks"})
		})

	}

	r.Run(":8080")
}

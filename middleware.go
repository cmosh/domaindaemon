package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/gin-gonic/gin"
)


func hook(c *gin.Context) {

	defer c.Request.Body.Close()
	if body, err := ioutil.ReadAll(c.Request.Body); err == nil {
		
		if err := json.Unmarshal(body, &webhook); err != nil {
			fmt.Printf("error while parsing json (%s)", err.Error())
		} else {
			fmt.Printf("received webhook body: %s", string(body))					
		}

	} else {
		fmt.Printf("error while reading webhook request (%s)", err.Error())		
	}

  c.Next()
}



	
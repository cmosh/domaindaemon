package main

import (
	"net/http"
	"runtime"
	"github.com/gin-gonic/gin"
)

func rootview(c *gin.Context) {    	
        c.HTML(http.StatusOK, "index.tmpl", gin.H{
            "runtime": runtime.Version(),
            "bot":  *me.Result.Username,
            "messages": "We are Torchwood",
        })
    }

// func messageview(c *gin.Context) {		
//         c.HTML(http.StatusOK, "index.tmpl", gin.H{
//             "runtime": runtime.Version(),
//             "bot":  *me.Result.Username,
//             "messages": getmessages(),
//         })
// }
package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"flag"
	bot "github.com/meinside/telegram-bot-go"
)


const (  
    TypingDelaySeconds = 1
)


var (
    theApiToken = *flag.String("bot", os.Getenv("BOT_TOKEN"), "bot")
    b = *bot.NewClient(theApiToken)
    webhook bot.Update
    me = b.GetMe()
)


func main() {	 	
	router := gin.Default()	
	router.LoadHTMLGlob("public/*") 
    router.GET("/",rootview)
    // router.GET("/messages",messageview)
    router.POST(theApiToken,hook,botmessage)
	router.Run(":8080")	
}

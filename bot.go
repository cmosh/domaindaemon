package main

import (
    "fmt"  
    "log"
    "time"
    "github.com/gin-gonic/gin"
    bot "github.com/meinside/telegram-bot-go"
)

type Lambda func()

func botmessage(c *gin.Context) {

    if webhook.HasMessage(){    
      b.SendChatAction(webhook.Message.Chat.Id, bot.ChatActionTyping)
      time.Sleep(TypingDelaySeconds * time.Second)


       message := fmt.Sprintf("Thank you %s! But I'm not sure why you sent me that.", (*webhook.Message.From.FirstName))
                options := map[string]interface{}{}
                sendmsg(message, options)
     
        // switch {
        //     case webhook.HasCallbackQuery(): 

        //          ctxtest(ctx,docallback,*webhook.CallbackQuery.Data)

        //     case webhook.HasEditedMessage(): 

        //         ctxtest(ctx,doeditmessage,*webhook.Message.Text)

        //     case webhook.HasMessage() && webhook.Message.HasText():

        //         ctxtest(ctx,domessage,*webhook.Message.Text)             

        //     default: 

               

        // }

      } 

    }


    func sendmsg(msg string,options map[string]interface {}) bot.Message{   
			var sent bot.ApiResponseMessage 

            if sent = b.SendMessage(webhook.Message.Chat.Id, &msg, options); !sent.Ok {
               log.Printf("*** failed to send message: %s\n", *sent.Description)
            }

            return *sent.Result
            
}


//   func ctxtest(ctx string,lam Lambda,msg string) {
//       delcontext()
//        if ctx == "none" ||   strings.Contains(msg, "/") {
//                       lam()
//                 }else if msg == "cancel" || *webhook.Message.Text == "cancel" {                                                  
//                       cancelmsg()      
//                 }else{                      
//                       ctxswitch(ctx)      
//                 }
//   }
 

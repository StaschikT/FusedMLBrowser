var express = require("express")
var app = express()
app.get("/", function(req, resp){
  resp.send('logic{-\n var i = 2 + 4;  -}*@default "Default text" *@red "red text" *@green "green text" *@yellow "yellow text" *@default "The value of i is:" *@blue "$i"')
})
app.listen(3000)

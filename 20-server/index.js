// This is server created for GO
// we will use this server to fetch Data using GO
// install node and try to run this server by npm install and command "npm start"
// make sure to run this server and proceed to next folders
const express = require("express")
const app = express()
const port = 3000

app.use(express.json())
app.use(express.urlencoded({extended:true}))

app.get("/",(req,res)=>{
    res.send("Hello, Welcome to GO Lang Demo Server")
})

app.get("/get",(req,res)=>{
    res.json({"message":"Hello, I am Pradeep Varma"})
})

app.post("/post",(req,res)=>{
    let body = req.body
    res.send(body)
})

app.post("/postform",(req,res)=>{
    res.send(JSON.stringify(req.body))
})


app.listen(port,()=>{
    console.log(`Server is Running on  http://localhost:${port}`);
    
})
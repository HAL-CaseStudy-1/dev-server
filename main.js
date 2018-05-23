const express = require("express")
const app = express()
const controllers = require("./controllers")

app.get("/", (req, res) => {
  res.send(true)
})

app.get("/top", controllers.top)

const PORT = process.env.PORT || 4000
const HOST = process.env.HOST || "127.0.0.1"
app.listen(PORT, HOST, () => {
  console.log(`listen on port: http://${HOST}:${PORT}`)
})
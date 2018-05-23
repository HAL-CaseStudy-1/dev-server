const express = require("express")
const app = express()

app.get("/", (req, res) => {
  res.json({
    hoge: true
  })
})

const PORT = process.env.PORT || 3000
app.listen(PORT, () => {
  console.log(`listen on port: http://localhost:${PORT}`)
})
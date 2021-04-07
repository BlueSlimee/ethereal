const { createCanvas, registerFont } = require('canvas')
registerFont('../public/fonts/montserrat/Montserrat-Bold.ttf', { family: 'Montserrat', weight: 'bold' })
module.exports = async (req, res) => {
  const { name = "jesus" } = req.query
  res.setHeader('content-type', 'application/png')

  const canvas = createCanvas(1000, 300)
  const ctx = canvas.getContext('2d')
  ctx.font = 'bold 12px Montserrat'
  
  ctx.fillText(name, 500, 150)
  res.status(200).send(await canvas.toBuffer('application/png'))
}


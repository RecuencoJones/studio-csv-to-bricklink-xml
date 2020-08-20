const fs = require('fs')
const csv = require('csv-parser')
const xmlbuilder = require('xmlbuilder2')

async function loadCSV(file) {
  return new Promise((resolve, reject) => {
    const rows = []

    fs.createReadStream(file)
      .on('error', reject)
      .pipe(csv({ separator: '\t' }))
      .on('data', (row) => rows.push(row))
      .on('end', () => resolve(rows))
      .on('error', reject)
  })
}

(async () => {
  const file = process.argv[2]
  const data = await loadCSV(file)
  const doc = xmlbuilder.create()

  data.reduce((accum, next) =>
    accum
      .ele('item')
        .ele('itemtype')
          .txt('P')
        .up()
        .ele('itemid')
          .txt(next.BLItemNo)
        .up()
        .ele('color')
          .txt(next.BLColorId)
        .up()
        .ele('minqty')
          .txt(next.Qty)
        .up()
      .up(),
    doc.ele('inventory')
  )

  const xml = doc.end({ prettyPrint: true })

  console.log(xml)
})()

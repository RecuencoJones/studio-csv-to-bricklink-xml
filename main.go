package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"

	"github.com/beevik/etree"
	"github.com/jszwec/csvutil"
)

func main() {
	file := os.Args[1]

	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	r := csv.NewReader(f)
	r.Comma = '\t'

	type Row struct {
		BLItemNo      string `csv:"BLItemNo"`
		ElementID     string `csv:"ElementId"`
		LdrawID       string `csv:"LdrawId"`
		PartName      string `csv:"PartName"`
		BLColorID     string `csv:"BLColorId"`
		LDrawColorID  string `csv:"LDrawColorId"`
		ColorName     string `csv:"ColorName"`
		ColorCategory string `csv:"ColorCategory"`
		Qty           string `csv:"Qty"`
		Weight        string `csv:"Weight"`
	}

	var rows []Row

	dec, err := csvutil.NewDecoder(r)
	if err != nil {
		panic(err)
	}

	for {
		var r Row
		if err := dec.Decode(&r); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		rows = append(rows, r)
	}

	doc := etree.NewDocument()
	doc.CreateProcInst("xml", `version="1.0"`)

	inventory := doc.CreateElement("inventory")

	for _, row := range rows {
		item := inventory.CreateElement("item")

		item.CreateElement("itemtype").CreateText("P")
		item.CreateElement("itemid").CreateText(row.BLItemNo)
		item.CreateElement("color").CreateText(row.BLColorID)
		item.CreateElement("minqty").CreateText(row.Qty)
	}

	doc.Indent(2)
	doc.WriteTo(os.Stdout)
}

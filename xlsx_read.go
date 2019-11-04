package main

import (
  "fmt"

  "github.com/tealeg/xlsx"
)

func main() {
  excel, err := xlsx.OpenFile("ExcelFile.xlsx") 

  if err != nil {
    fmt.Printf(err.Error())
  }

  sheet := excel.Sheets[0]
  cell_A1 := sheet.Cell(0, 0).Value
  cell_A2 := sheet.Cell(1, 0).Value
  cell_B1 := sheet.Cell(0, 1).Value
  cell_B2 := sheet.Cell(1, 1).Value

  fmt.Println(cell_A1)
  fmt.Println(cell_A2)
  fmt.Println(cell_B1)
  fmt.Println(cell_B2)

}

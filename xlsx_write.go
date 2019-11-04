package main

import (
    "fmt"

    "github.com/tealeg/xlsx"
)

func main() {
    file := xlsx.NewFile() 
    sheet, err := file.AddSheet("Sheet1") // シートの追加
    if err != nil {
        fmt.Printf(err.Error())
    }
    // 指定のシートのセルへ値挿入
    sheet.Cell(0, 0).Value = "A1"
    sheet.Cell(1, 0).Value = "A2"
    sheet.Cell(0, 1).Value = "B1"
    sheet.Cell(1, 1).Value = "B2"
    err = file.Save("ExcelFile.xlsx") // exelファイルを作成
    if err != nil {
        fmt.Printf(err.Error())
    }
}

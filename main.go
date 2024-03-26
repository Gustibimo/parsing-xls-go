package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

type Cities struct {
	Name       string
	Population float64
}

func main() {
	// header := []string{"Name", "Country", "Population"}
	sheetName := "cities"
	wb, err := xlsx.OpenFile("test.xlsx")
	if err != nil {
		panic(err)
	}

	cities := []Cities{}
	// for _, sheet := range wb.Sheets {
	// 	for _, row := range sheet.Rows {
	// 		city := Cities{}
	// 		for i, cell := range row.Cells {
	// 			switch i {
	// 			case 0:
	// 				city.Name = cell.String()
	// 			case 2:
	// 				city.Population, _ = cell.Float()
	// 			}
	// 		}
	// 		cities = append(cities, city)
	// 	}
	// }
	//
	// fmt.Println("print cities")
	// for _, city := range cities {
	// 	fmt.Println(city)
	// }

	// sheet, ok := wb.Sheet[sheetName]
	// if !ok {
	// 	fmt.Println("Sheet not found")
	// 	return
	// }
	// for _, row := range sheet.Rows {
	// 	city := Cities{}
	// 	for i, cell := range row.Cells {
	// 		switch i {
	// 		case 0:
	// 			city.Name = cell.String()
	// 		case 2:
	// 			city.Population, _ = cell.Float()
	// 		}
	// 	}
	// 	cities = append(cities, city)
	// }
	//
	// fmt.Println("print cities")
	//
	// for _, city := range cities {
	// 	fmt.Println(city)
	// }
	//
	//

	sheet, ok := wb.Sheet[sheetName]
	if !ok {
		fmt.Println("Sheet not found")
		return
	}

	for i, row := range sheet.Rows {
		if i == 0 {
			continue
		}
		city := Cities{}
		for i, cell := range row.Cells {
			switch i {
			case 0:
				city.Name = cell.String()
			case 2:
				city.Population, _ = cell.Float()
			}
		}
		cities = append(cities, city)
	}

	for _, city := range cities {
		fmt.Println(city)
	}
}

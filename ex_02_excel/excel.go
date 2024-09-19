package main

import (
	"fmt"
	"log"

	"github.com/x-funs/go-fun"
	"github.com/xuri/excelize/v2"
)

type project struct {
	ProjectCity  string
	ProjectArea  string
	ProjectName  string
	Construction int
	PdfUrl       string
}

func main() {
	// xlsx 格式，2007 之前老版本不支持
	f, err := excelize.OpenFile("20240918.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 子表
	rows, err := f.GetRows("江西房建市政等")
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取区间行号
	var start int
	var end int
	for i, row := range rows {
		if fun.SliceContains(row, "今日新增项目") {
			start = i
		}
		if fun.SliceContains(row, "今日变动项目") {
			end = i
		}
	}
	fmt.Println(start)
	fmt.Println(end)

	// 遍历数据
	datas := make([]project, 0)
	for i, row := range rows {
		if i > start && i < end {
			// 是否施组，完整项目，或其他条件
			if len(row) > 6 && row[6] == "是" {
				var data project
				for col, colCell := range row {
					data.Construction = 1
					if col == 3 {
						data.ProjectCity = colCell
					}
					if col == 4 {
						data.ProjectArea = colCell
					}
					if col == 5 {
						data.ProjectName = colCell
					}
					if col == 17 && colCell == "招" {
						if res, link, _ := f.GetCellHyperLink("江西房建市政等", "R"+fun.ToString(i)); res == true {
							data.PdfUrl = link
						}
					}

				}
				datas = append(datas, data)
			}
			fmt.Println()
		}
	}

	log.Println(datas)
}

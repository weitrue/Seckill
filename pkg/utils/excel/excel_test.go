package excel

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuri/excelize/v2"
)

func TestGetRows(t *testing.T) {
	f, err := os.Open("./testdata/员工导入模板.xlsx")
	assert.NoError(t, err)
	defer f.Close()

	// 20MB
	rows, err := GetRows(f, "Sheet1", excelize.Options{UnzipSizeLimit: 20 << 20, UnzipXMLSizeLimit: 20 << 20})
	assert.NoError(t, err)

	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

func TestGetFilteredRows(t *testing.T) {
	f, err := os.Open("./testdata/员工导入模板.xlsx")
	assert.NoError(t, err)
	defer f.Close()

	rows, err := GetFilteredRows(f, "Sheet1", 3, 1)
	assert.NoError(t, err)

	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

func TestReadRows(t *testing.T) {
	f, err := os.Open("./testdata/员工导入模板.xlsx")
	assert.NoError(t, err)
	defer f.Close()

	err = ReadRows(f, "Sheet1", func(rowNum int, columns []string) bool {
		if rowNum == 4 {
			return false
		}
		if rowNum > 1 && len(columns) >= 3 {
			for _, column := range columns {
				fmt.Print(column, "\t")
			}
			fmt.Println()
		}
		return true
	})
	assert.NoError(t, err)
}

func TestWriteRows(t *testing.T) {
	f, err := os.Open("./testdata/员工导入模板.xlsx")
	assert.NoError(t, err)
	defer f.Close()

	handler := func(rowNum int) (columns []interface{}, needWrite, isContinue bool) {
		rows := [][]interface{}{
			{"赵六", "Java开发", "13400000000"},
			{"陈七", "产品", "13500000000"},
			{"杨八", "财务", "13600000000"},
		}

		if rowNum > 1 && rowNum < 5 {
			return rows[rowNum-2], true, true
		} else if rowNum >= 5 {
			return nil, false, false
		}
		return nil, false, true
	}

	err = WriteRows(f, "Sheet1", "./testdata/员工导入模板_写入测试.xlsx", handler)
	assert.NoError(t, err)
}

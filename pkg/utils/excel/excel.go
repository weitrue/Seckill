package excel

import (
	"io"
	"log"

	"github.com/pkg/errors"
	"github.com/xuri/excelize/v2"
)

// GetRows 获取 excel 表上的所有行数据
func GetRows(r io.Reader, sheet string, opt ...excelize.Options) ([][]string, error) {
	f, err := excelize.OpenReader(r, opt...)
	if err != nil {
		return nil, errors.WithMessage(err, "excelize.OpenReader err")
	}
	defer f.Close()

	rows, err := f.GetRows(sheet, opt...)
	if err != nil {
		return nil, errors.WithMessage(err, "f.GetRows err")
	}

	return rows, nil
}

// GetFilteredRows 获取 excel 表上的所有大于等于指定行长度的行数据（忽略前 skipRow 行）
func GetFilteredRows(r io.Reader, sheet string, rowLength, skipRow int, opt ...excelize.Options) ([][]string, error) {
	rows, err := GetRows(r, sheet, opt...)
	if err != nil {
		return nil, err
	}

	var filteredRows [][]string
	if skipRow < 0 {
		skipRow = 0
	}
	if l := len(rows); l > skipRow {
		for i := skipRow; i < l; i++ {
			if len(rows[i]) >= rowLength {
				filteredRows = append(filteredRows, rows[i])
			}
		}
	}

	return filteredRows, nil
}

// ReadHandler 流式读取处理器
type ReadHandler func(rowNum int, columns []string) (isContinue bool)

// ReadRows 流式读取处理 excel 表上的行数据
func ReadRows(r io.Reader, sheet string, handler ReadHandler, opt ...excelize.Options) error {
	f, err := excelize.OpenReader(r, opt...)
	if err != nil {
		return errors.WithMessage(err, "excelize.OpenReader err")
	}
	defer f.Close()

	rows, err := f.Rows(sheet)
	if err != nil {
		return errors.WithMessage(err, "f.Rows err")
	}
	defer rows.Close()

	rowNum := 0
	for rows.Next() {
		rowNum++
		columns, err := rows.Columns(opt...)
		if err != nil {
			log.Printf("[excel] rows.Columns err, err: %v", err)
			continue
		}
		if isContinue := handler(rowNum, columns); !isContinue {
			return nil
		}
	}

	return nil
}

// WriteHandler 流式写入处理器
type WriteHandler func(rowNum int) (columns []interface{}, needWrite, isContinue bool)

// WriteRows 流式写入行数据至指定 excel 表中
func WriteRows(r io.Reader, sheet, saveAs string, handler WriteHandler, opt ...excelize.Options) error {
	f, err := excelize.OpenReader(r, opt...)
	if err != nil {
		return errors.WithMessage(err, "excelize.OpenReader err")
	}
	defer f.Close()

	sw, err := f.NewStreamWriter(sheet)
	if err != nil {
		return errors.WithMessage(err, "f.NewStreamWriter err")
	}

	rowNum := 0
	for {
		rowNum++
		columns, needWrite, isContinue := handler(rowNum)
		if !isContinue {
			break
		}
		if needWrite {
			cellName, err := excelize.CoordinatesToCellName(1, rowNum)
			if err != nil {
				log.Printf("[excel] excelize.CoordinatesToCellName err, err: %v", err)
				continue
			}
			if err := sw.SetRow(cellName, columns); err != nil {
				log.Printf("[excel] sw.SetRow err, err: %v", err)
				continue
			}
		}
	}

	if err := sw.Flush(); err != nil {
		return errors.WithMessage(err, "sw.Flush err")
	}

	if err := f.SaveAs(saveAs); err != nil {
		return errors.WithMessage(err, "f.SaveAs err")
	}

	return nil
}

package utils

import (
	"base-system-backend/enums/errmsg"
	"bytes"
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
)

func ToExcel(tableName string, titleList []string, dataList []interface{}) (content io.ReadSeeker, err error, debugInfo interface{}) {
	//生成一个新的文件
	file := xlsx.NewFile()
	//添加sheet页
	sheet, _ := file.AddSheet(tableName)
	//插入表头
	titleRow := sheet.AddRow()
	for _, v := range titleList {
		cell := titleRow.AddCell()
		cell.Value = v
	}
	//插入内容
	for _, v := range dataList {
		row := sheet.AddRow()
		row.WriteStruct(v, -1)
	}
	if err = file.Save("~/programming/go/base-system-backend/static/aaaaa.xlsx"); err != nil {
		return
	}
	//if err != nil{
	//	return nil, fmt.Errorf("%s表%w", tableName, errmsg.ReadFailed), err.Error()
	//}
	var buffer bytes.Buffer
	if err = file.Write(&buffer); err != nil {
		return nil, fmt.Errorf("%s表%w", tableName, errmsg.ReadFailed), err.Error()
	}
	content = bytes.NewReader(buffer.Bytes())
	return
}

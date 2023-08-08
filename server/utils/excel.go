package utils

import (
	"base-system-backend/constants/errmsg"
	"base-system-backend/global"
	"base-system-backend/utils/common"
	"bytes"
	"fmt"
	"github.com/tealeg/xlsx"
	"io"
	"strings"
)

// ToExcel
//  @Description: 下载 Excel
//  @param tableName 表名
//  @param titleList 标题 Slice
//  @param dataList 数据 Slice
//  @return content io 上下文
//  @return err 创建路径/保存文件失败异常
//  @return debugInfo 错误调试信息

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
	savePath := strings.Join(global.CONFIG.Static.Log, "")
	err, debugInfo = common.FileCheck(savePath)
	if err != nil {
		return nil, err, debugInfo
	}
	if err = file.Save(strings.Join([]string{savePath, "/", tableName, ".xlsx"}, "")); err != nil {
		return
	}
	var buffer bytes.Buffer
	if err = file.Write(&buffer); err != nil {
		return nil, fmt.Errorf("%s表%w", tableName, errmsg.ReadFailed), err.Error()
	}
	content = bytes.NewReader(buffer.Bytes())
	return
}

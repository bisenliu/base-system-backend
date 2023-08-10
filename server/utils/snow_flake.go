package utils

import "base-system-backend/global"

// GenID
//  @Description: 基于雪花算法生成分布式 ID
//  @return int64 分布式 ID

func GenID() int64 {
	return global.Node.Generate().Int64()
}

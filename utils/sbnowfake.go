package utils

import "base-system-backend/global"

func GenID() int64 {
	return global.Node.Generate().Int64()
}

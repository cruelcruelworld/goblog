package types

import (
	"goblog/pkg/logger"
	"strconv"
)

func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

func StringToUint64(str string) uint64 {
	num, err := strconv.ParseUint(str, 10, 64)

	if err != nil {
		logger.LogError(err)
	}

	return num
}

func Uint64ToString(num uint64) string {
	return strconv.FormatUint(num, 10)
}

func StringToInt(str string) int {
	num, err := strconv.Atoi(str)

	if err != nil {
		logger.LogError(err)
	}

	return num
}

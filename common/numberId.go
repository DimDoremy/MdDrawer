package common

import (
	"hash/crc32"

	"github.com/gofrs/uuid"
)

//随机生成uuid
func RandomUUID() string {
	uid, err := uuid.NewV4()
	if err != nil {
		Logger.Error("Get UUID failed", ErrPacker(err))
	}
	return uid.String()
}

//通过字符串获取纯数字id
func GetNumberId(s string) uint32 {
	return crc32.Checksum([]byte(s), crc32.IEEETable)
}

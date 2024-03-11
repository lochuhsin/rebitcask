package dao

import (
	"rebitcask/internal/util"
	"strconv"
	"strings"
)

// CRC::TimeStamp::KeyLen::Key::ValueLen::Value
func Serialize(k Entry) (string, error) {
	var builder strings.Builder
	builder.WriteString("CRC::")
	builder.WriteString(strconv.Itoa(int(k.CreateTime)))
	builder.WriteString("::")
	builder.WriteString(strconv.Itoa(len(k.Key)))
	builder.WriteString("::")
	builder.Write(k.Key)
	builder.WriteString("::")
	builder.WriteString(strconv.Itoa(len(k.Key)))
	builder.WriteString("::")
	builder.WriteString(util.BytesToString(k.Val))
	return builder.String(), nil
}

func DeSerialize(line string) (Entry, error) {
	// CRC::TimeStamp::KeyLen::Key::ValueLen::Value
	strList := strings.Split(line, "::")
	// crc = strList[0] for validation usage
	timestamp := strList[1]
	// KeyLen := strList[2] for validation usage
	key := strList[3]
	// ValueLen := strList[4] for validation usage
	val := strList[5]

	ts, err := strconv.Atoi(timestamp)
	if err != nil {
		panic(err)
	}
	return Entry{
		Key:        util.StringToBytes(key),
		Val:        util.StringToBytes(val),
		CreateTime: int64(ts),
	}, nil

}

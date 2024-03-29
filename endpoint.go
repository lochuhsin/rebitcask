package rebitcask

import (
	"rebitcask/internal"
	"rebitcask/internal/dao"
	"rebitcask/internal/setting"
	"rebitcask/internal/util"
)

func Get(k string) (string, bool) {
	/**
	 * First, check does the value exist in memory
	 *
	 * Second, check does the value exist in segment
	 *
	 * Note: exists meaning that the key exists, and the value is not tombstone
	 */
	bytes := util.StringToBytes(k)
	m, status := internal.GetMemoryManager().Get(bytes)
	if status {
		return checkTombstone(m)
	}

	s, status := internal.GetSegmentManager().GetValue(bytes)
	if status {
		return checkTombstone(s)
	}
	return "", false
}

func Set(k string, v string) error {
	manager := internal.GetMemoryManager()
	entry := dao.InitEntry(util.StringToBytes(k), util.StringToBytes(v))

	manager.SetRequestQ() <- entry
	<-manager.SetResponseQ()
	return nil
}

func Delete(k string) error {
	manager := internal.GetMemoryManager()
	entry := dao.InitTombEntry(util.StringToBytes(k))
	manager.SetRequestQ() <- entry
	<-manager.SetResponseQ()
	return nil
}

func Exist() (bool, error) {
	panic("Not implemented error")
}

func BulkCreate(k string) error {
	panic("Not implemented error")
}

func BulkUpdate(k string) error {
	panic("Not implemented error")
}

func BulkUpsert(k string) error {
	panic("Not implemented error")
}

func BulkDelete(k string) error {
	panic("Not implemented error")
}

func BulkGet(k ...string) ([]string, error) {
	panic("Not implemented error")
}

func checkTombstone(entry dao.Entry) (string, bool) {
	val := util.BytesToString(entry.Val)
	if val == setting.Config.TOMBSTONE {
		return "", false
	}
	return val, true
}

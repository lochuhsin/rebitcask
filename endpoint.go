package rebitcask

import (
	"rebitcask/internal/dao"
	"rebitcask/internal/memory"
	"rebitcask/internal/segment"
	"rebitcask/internal/setting"
	"rebitcask/internal/transaction"
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
	m, status := memory.GetMemoryManager().Get(bytes)
	if status {
		return checkTombstone(m)
	}

	s, status := segment.GetSegmentManager().GetValue(bytes)
	if status {
		return checkTombstone(s)
	}
	return "", false
}

func Set(k string, v string) error {
	manager := memory.GetMemoryManager()
	entry := dao.InitEntry(util.StringToBytes(k), util.StringToBytes(v))
	serialize, err := dao.Serialize(entry)
	if err != nil {
		panic(err)
	}
	transaction.GetCommitLogger().Add(serialize)

	manager.SetRequestQ() <- entry
	<-manager.SetResponseQ()
	return nil
}

func Delete(k string) error {
	manager := memory.GetMemoryManager()
	entry := dao.InitTombEntry(util.StringToBytes(k))
	entryB, err := dao.Serialize(entry)
	if err != nil {
		panic(err)
	}
	transaction.GetCommitLogger().Add(entryB)
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

package segment

import (
	"bufio"
	"fmt"
	"os"
	"rebitcask/internal/dao"
	"rebitcask/internal/settings"
)

func getSegmentFilePath(segId string) string {
	return fmt.Sprintf("%v%v%v%v", settings.ENV.DataPath, settings.SEGMENT_FILE_FOLDER, segId, settings.SEGMENT_FILE_EXT)
}

func getSegmentIndexFilePath(segId string) string {
	return fmt.Sprintf("%v%v%v%v", settings.ENV.DataPath, settings.INDEX_FILE_FOLDER, segId, settings.SEGMENT_KEY_OFFSET_FILE_EXT)
}

func getSegmentMetaDataFilePath(segId string) string {
	return fmt.Sprintf("%v%v%v%v", settings.ENV.DataPath, settings.SEGMENT_FILE_FOLDER, segId, settings.SEGMENT_FILE_METADATA_EXT)
}

func segmentToFile(s *Segment, pairs []dao.Pair) {
	/**
	 * Note, assuming that key in pairs are sorted in ascending order
	 */
	filePath := getSegmentFilePath(s.Id)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777) //TODO: optimize the mode
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	curroffset := 0
	s.smallestKey = pairs[0].Key.Val // the first key is the smallest value
	// TODO: convert this pair to generator pattern, hide inside segment, we don't need to know if the data needs to be serialized
	for _, p := range pairs {
		data, err := dao.Serialize(p)
		if err != nil {
			panic("Error while serializing data")
		}
		offset, err := writer.WriteString(data + settings.DATASAPARATER)
		if err != nil {
			panic("something went wrong while writing to segment")
		}
		// offset minus data saparater = the length of the data
		s.pIndex.Set(p.Key, curroffset, offset-len([]byte(settings.DATASAPARATER)))
		curroffset += offset
	}
	writer.Flush()
	file.Sync()
	s.smallestKey = pairs[0].Key.GetVal().(string)
	s.keyCount = len(pairs)
}

func segmentToMetadata(s *Segment) {
	filePath := getSegmentMetaDataFilePath(s.Id)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777) //TODO: optimize the mode
	if err != nil {
		panic(err)
	}
	defer file.Close()

	/**
	 * Currently only store level information for segment manager to backup
	 */
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(fmt.Sprintf("level::%v", s.Level))
	if err != nil {
		panic("something went wrong while writing segment metadata")
	}
	writer.Flush()
	// We don't need to fd.Sync() metadata, since the read is not necessarily to do
	// immediately read, like Get operation
}

func segmentIndexToFile(segment *Segment) {
	filePath := getSegmentIndexFilePath(segment.Id)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777) //TODO: optimize the mode
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	offsetMap := segment.pIndex.OffsetMap

	for key, val := range offsetMap {
		data := segmentIndexSerialize(key.Format(), val.Format())
		_, err := writer.WriteString(data + settings.DATASAPARATER)
		if err != nil {
			panic("something went wrong while writing to segment")
		}
	}

	writer.Flush()
	// We don't need to fd.Sync() metadata, since the read is not necessarily to do
	// immediately read, like Get operation, since this index is mainly for crash recovery
}

// TODO: refactor this
func segmentIndexSerialize(key string, val string) string {
	// format -> KeyDataType::KeyLen::Key::offset::length
	return fmt.Sprintf("%v::%v", key, val)
}

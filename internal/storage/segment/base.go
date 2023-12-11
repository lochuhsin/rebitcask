package segment

import (
	"bufio"
	"os"
	"rebitcask/internal/storage/dao"
	"time"

	"github.com/google/uuid"
)

/**
 * I'm going to use SSTable as segment implementation
 * The key point of SSTable is that the keys were sorted
 * in ascending order. Therefore the head of the file (usually it's the end of the file)
 * is the smallest key. This is helpful, since we store the smallest key of the segment
 * in memory. When we are looking up to see if key exists,
 * we only need to start looking at files that Segkeies who were smaller.
 * This increases the performance of lookup.
 *
 * Each segment accompanies a segment index
 * which contains all the key and offset to the segment
 *
 * Then we have another
 */

// Design a segment index structure that match the upper condition
// implemented using binary search tree
type Segment struct {
	id          string
	level       int    // reference from levelDB, using level indicate the compaction process
	smallestKey string // indicates the smallest key in current segment
	timestamp   int64  // the time that segment was created
	keyCount    int
}

func InitSegment() Segment {
	return Segment{id: uuid.New().String(), level: 0, smallestKey: "", timestamp: time.Now().UnixMicro()}
}

func (s *Segment) Get(k dao.NilString) (dao.Base, bool) {

	filePath := getSegmentFilePath(s.id)
	fd, err := os.Open(filePath)
	if err != nil {
		panic("Cannot open segment file")
	}
	defer fd.Close()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		line := scanner.Text()
		pair, err := dao.DeSerialize(line) // Figure out a better way to split between keys
		if err != nil {
			panic("Something went wrong while deserializing data")
		}

		if pair.Key.IsEqual(k) {
			return pair.Val, true
		}
	}

	return nil, false
}

func (s *Segment) WriteFile(pairs []dao.Pair) {
	/**
	 * Note, assuming that key in pairs are sorted in ascending order
	 */
	filePath := getSegmentFilePath(s.id)
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	curroffset := 0
	s.smallestKey = pairs[0].Key.Val // the first key is the smallest value
	for _, p := range pairs {
		data, err := dao.Serialize(p)
		if err != nil {
			panic("Error while serializing data")
		}
		offset, err := writer.WriteString(data + "\n") // Figure out a better way to split between keys
		if err != nil {
			panic("something went wrong while writing to segment")
		}
		curroffset += offset
	}
	writer.Flush()

	s.smallestKey = pairs[0].Key.GetVal().(string)
	s.keyCount = len(pairs)
}

type SegmentStack struct {
	stack []Segment
}

func InitSegmentStack() SegmentStack {
	return SegmentStack{stack: []Segment{}}
}

func (s *SegmentStack) Add(seg Segment) {
	s.stack = append(s.stack, seg)
}

func (s *SegmentStack) Pop() (Segment, bool) {
	for len(s.stack) > 0 {
		seg := s.stack[len(s.stack)-1]
		s.stack = s.stack[:len(s.stack)-1]
		return seg, true
	}
	return *new(Segment), false
}

func (s *SegmentStack) Size() int { return len(s.stack) }

func (s *SegmentStack) list() *[]Segment {

	newSeg := []Segment{}

	for i := len(s.stack) - 1; i >= 0; i-- {
		newSeg = append(newSeg, s.stack[i])
	}
	return &newSeg
}

// order_by timestamp
type SegmentCollection struct {
	/**
	 * We are using stack to get the native characteristics
	 * of first in last out, which meets the requirements of
	 * order by timestamp
	 */
	zeroLevelSeg SegmentStack
	level        map[int][]Segment
	maxLevel     int // whenever a compaction starts, adjust this maxLevel
}

func InitSegmentCollection() SegmentCollection {
	stack := InitSegmentStack()
	return SegmentCollection{level: map[int][]Segment{}, zeroLevelSeg: stack, maxLevel: 0}
}

func (s *SegmentCollection) AddSegment(seg Segment) {
	s.zeroLevelSeg.Add(seg)
}

func (s *SegmentCollection) CompactionCondition() bool {
	/**
	 * Implement the compaction condtion for manager to determine
	 * When we are starts to compact
	 */
	return false
}

func (s *SegmentCollection) Compaction() {
	panic("not implemented yet")
}

type SegmentIndex struct {
}

type SegmentIndexCollection struct {
}

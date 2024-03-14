package memory

import (
	"rebitcask/internal/memory/models"
	"testing"
)

func TestInitMemoryStorage(t *testing.T) {
	mBlockStorage := NewMemoryStorage()
	currBlock := mBlockStorage.getCurrentBlock()

	if currBlock != nil {
		t.Error("initial state of block should be empty")
	}
}

func TestMemoryStorageBlockCount(t *testing.T) {
	mBlockStorage := NewMemoryStorage()
	count := 5
	for i := 0; i < 5; i++ {
		mBlockStorage.createNewBlock(models.ModelType("bst"))
	}
	if mBlockStorage.getBlockCount() != count {
		t.Error("block count inconsistent")
	}
}

func TestMemoryStorageBlockOrder(t *testing.T) {
	mBlockStorage := NewMemoryStorage()
	count := 100
	blockIds := []BlockId{}
	for i := 0; i < count; i++ {
		mBlockStorage.createNewBlock(models.ModelType("bst"))
		blockId := mBlockStorage.getCurrentBlockId()
		blockIds = append(blockIds, blockId)
	}
	iterateBlocks := mBlockStorage.iterateExistingBlocks()
	for i := 0; i < count; i++ {
		if blockIds[i] != iterateBlocks[count-i-1].Id {
			t.Error("iterate blocks id mismatch")
		}
	}
}

func TestMemoryStorageRemoveBlock(t *testing.T) {
	mBlockStorage := NewMemoryStorage()
	count := 100
	removedBlocks := []BlockId{}
	for i := 0; i < count; i++ {
		mBlockStorage.createNewBlock(models.ModelType("bst"))
		blockId := mBlockStorage.getCurrentBlockId()
		removedBlocks = append(removedBlocks, blockId)
	}

	for _, blockId := range removedBlocks {
		mBlockStorage.removeMemoryBlock(blockId)
	}

	if mBlockStorage.getBlockCount() != 0 {
		t.Error("storage should be empty")
	}

	if mBlockStorage.getCurrentBlock() != nil {
		t.Error("current block should be empty")
	}
}

package rebitcask

import (
	"fmt"
	"os"
	"rebitcask/internal/memory"
	"rebitcask/internal/scheduler"
	"rebitcask/internal/segment"
	"rebitcask/internal/settings"
)

func Init(envPaths ...string) {
	/**
	 * Should call this, whenever the server is up
	 */

	settings.InitConfig(envPaths...)
	config := settings.Config

	segDir := fmt.Sprintf("%s%s", config.DataFolderPath, settings.SEGMENT_FILE_FOLDER)
	indexDir := fmt.Sprintf("%s%s", config.DataFolderPath, settings.INDEX_FILE_FOLDER)
	os.MkdirAll(segDir, os.ModePerm)
	os.MkdirAll(indexDir, os.ModePerm)

	memory.InitMemoryManager(memory.ModelType(settings.Config.MemoryModel))
	segment.InitSegmentManager()
	scheduler.InitScheduler()
}

package test

import (
	"fmt"
	"os"
	"rebitcask"
	"rebitcask/internal/setting"
	"testing"
	"time"
)

func setup() {
	rebitcask.Setup(".env.test")
}

func teardown() {
	time.Sleep(time.Second * 3)
	os.RemoveAll(setting.Config.DATA_FOLDER_PATH)
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func TestStorageSet(t *testing.T) {
	dataCount := setting.Config.MEMORY_COUNT_LIMIT*3 + 1

	keys, vals := GenerateLowDuplicateRandom(dataCount)
	for i, k := range keys {
		err := rebitcask.Set(k, vals[i])
		if err != nil {
			t.Error("Something went wrong while setting")
		}
	}
}

func TestStorageGet(t *testing.T) {

	dataCount := setting.Config.MEMORY_COUNT_LIMIT*3 + 1

	keys, _ := GenerateLowDuplicateRandom(dataCount)
	for _, k := range keys {
		_, _ = rebitcask.Get(k)
	}
}

func TestStorageDelete(t *testing.T) {

	config := setting.Config
	dataCount := config.MEMORY_COUNT_LIMIT*3 + 1
	keys, _ := GenerateLowDuplicateRandom(dataCount)
	for _, k := range keys {
		err := rebitcask.Delete(k)
		if err != nil {
			t.Error("Something went wrong while setting")
		}
	}
}

func TestStorageSetGet(t *testing.T) {
	env := setting.Config
	dataCount := env.MEMORY_COUNT_LIMIT*3 + 1

	keys, values := GenerateLowDuplicateRandom(dataCount)
	for i, k := range keys {
		err := rebitcask.Set(k, values[i])
		if err != nil {
			t.Error("Something went wrong while setting")
		}
	}

	for i, k := range keys {
		val, status := rebitcask.Get(k)
		if !status {
			t.Error("the key should exist", k)
			break
		}

		if val != values[i] {
			t.Error("the value should be equal to the generated value")
		}
	}
}

func TestStorageSetDelete(t *testing.T) {
	dataCount := setting.Config.MEMORY_COUNT_LIMIT*10 + 1

	keys, vals := GenerateLowDuplicateRandom(dataCount)
	for i, k := range keys {
		err := rebitcask.Set(k, vals[i])
		if err != nil {
			t.Error("Something went wrong while setting")
		}
	}

	for _, k := range keys {
		err := rebitcask.Delete(k)
		if err != nil {
			t.Error("Delete operation should work")
		}
	}

	for _, k := range keys {
		val, status := rebitcask.Get(k)
		if status {
			str := fmt.Sprintf("the key should not exist: %v", val)
			t.Error(str)
		}
	}
}

func TestEmptyGet(t *testing.T) {
	keys, _ := GenerateLowDuplicateRandom(100)

	for _, k := range keys {
		_, status := rebitcask.Get(k)
		if status {
			t.Error("the key should not exist")
		}
	}
}

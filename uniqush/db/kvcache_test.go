package db

import (
    "testing"
    "fmt"
    "os"
)

type FakeFlusher struct {
}

func (f *FakeFlusher) Set(key string, v interface{}) os.Error {
    fmt.Print("Flush: ", key)
    return nil
}

func (f *FakeFlusher) Remove(key string, v interface{}) os.Error {
    fmt.Print("Remove: ", key)
    return nil
}

func (f *FakeFlusher) Flush() os.Error {
    return nil
}

func TestLRUCache(t *testing.T) {
    strategy := NewLRUPeriodFlushStrategy(3, 100)
    storage := NewInMemoryKeyValueStorage(10)
    flusher := &FakeFlusher{}

    cache := NewKeyValueCache(storage, strategy, flusher)
    fmt.Print("Start LRU cache test ...\t")

    for i := 0; i < 10; i++ {
        str := fmt.Sprint(i)
        cache.Show(str, str)
    }

    keys, _ := cache.Keys()
    if !same(convert2string([]int{7,9,8}), keys) {
        t.Errorf("should be [7 8 9], but %v", keys)
    }

    if v, _ := cache.Get("1"); v != nil {
        t.Errorf("%v should not be in cache", v)
    }

    cache.Get("7")
    cache.Show("1", "1")

    if v, _ := cache.Get("8"); v != nil {
        keys, _ := cache.Keys()
        t.Errorf("%v should not be in cache; cache content: %v", v, keys)
    }
    fmt.Print("OK\n")
}


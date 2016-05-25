//
// Author: ikether
// Email: ikether@126.com
//
// Copyright 2016 ikether. All Right reserved.

package cache

import (
	"bytes"
	"encoding/binary"

	"github.com/zjwdmlmx/freecache"
)

type cacheProxy struct {
	cache *freecache.Cache
}

func newCache(size int) (cache *cacheProxy) {
	cache = new(cacheProxy)
	cache.cache = freecache.NewCache(size)
	return
}

func (cache *cacheProxy) Set(key, value []byte, expire int) (err error) {
	err = cache.cache.Set(key, value, expire)
	return
}

func (cache *cacheProxy) Get(key []byte) (value []byte, err error) {
	value, err = cache.cache.Get(key)
	return
}

func (cache *cacheProxy) SSetString(key, value string, expire int) (err error) {
	err = cache.cache.Set([]byte(key), []byte(value), expire)
	return
}

func (cache *cacheProxy) SGetString(key string) (value string, err error) {
	var (
		valueByte []byte
	)
	valueByte, err = cache.cache.Get([]byte(key))

	if err != nil {
		return
	}

	value = string(valueByte)
	return
}

func (cache *cacheProxy) SSetInt64(key string, value int64, expire int) (err error) {
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, value)

	if err != nil {
		return
	}

	err = cache.cache.Set([]byte(key), buf.Bytes(), expire)
	return
}

func (cache *cacheProxy) SGetInt64(key string) (value int64, err error) {
	var (
		byteValue []byte
	)

	byteValue, err = cache.cache.Get([]byte(key))

	if err != nil {
		return
	}

	binary.Read(bytes.NewReader(byteValue), binary.BigEndian, &value)

	return
}

func (cache *cacheProxy) SSetUint64(key string, value uint64, expire int) (err error) {
	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, value)

	if err != nil {
		return
	}

	err = cache.cache.Set([]byte(key), buf.Bytes(), expire)
	return
}

func (cache *cacheProxy) SGetUint64(key string) (value uint64, err error) {
	var (
		byteValue []byte
	)

	byteValue, err = cache.cache.Get([]byte(key))

	if err != nil {
		return
	}

	binary.Read(bytes.NewReader(byteValue), binary.BigEndian, &value)

	return
}

var Cached *cacheProxy

func init() {
	Cached = newCache(40 * 1024 * 1024)
}

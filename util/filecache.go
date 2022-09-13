package util

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
	"time"
)

type FileCache struct {
	CacheDir string
}

type cacheData struct {
	Value       string    `json:"value"`
	ExpiresTime time.Time `json:"expires_time"`
}

var IsNotExist = errors.New("未找到当前key的值")
var IsExpires = errors.New("该值已经过期")

func NewFileCache(cacheDir string) (*FileCache, error) {
	if cacheDir == "" {
		cacheDir = "./wechat_cache/"
	}

	//判断缓存目录不存在则创建
	_, err := os.Stat(cacheDir)
	if err != nil && os.IsNotExist(err) {
		err = os.Mkdir(cacheDir, os.FileMode(0666))
		if err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	return &FileCache{CacheDir: cacheDir}, nil
}

// SetData 设置值
func (f FileCache) SetData(ctx context.Context, key string, val string, expiresSeconds int64) error {
	fh, err := os.OpenFile(f.CacheDir+key, os.O_CREATE|os.O_WRONLY, os.FileMode(0666))
	if err != nil {
		return err
	}
	defer func(fh *os.File) {
		err = fh.Close()
		if err != nil {
			//TODO 记录日志
		}
	}(fh)

	data := cacheData{
		Value:       val,
		ExpiresTime: time.Now().Add(time.Duration(expiresSeconds) * time.Second),
	}
	dataJson, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = fh.Write(dataJson)
	if err != nil {
		return err
	}

	err = fh.Sync()
	if err != nil {
		return err
	}

	return nil
}

func (f FileCache) GetData(ctx context.Context, key string) (string, error) {
	fh, err := os.OpenFile(f.CacheDir+key, os.O_RDONLY, os.FileMode(0666))
	if err != nil && errors.Is(err, os.ErrNotExist) {
		return "", IsNotExist
	}
	defer func(fh *os.File) {
		err = fh.Close()
		if err != nil {
			//TODO 记录日志
		}
	}(fh)

	content, err := io.ReadAll(fh)
	if err != nil {
		return "", err
	}
	cache := new(cacheData)
	if err = json.Unmarshal(content, cache); err != nil {
		return "", err
	}

	if !cache.ExpiresTime.After(time.Now()) {
		return "", IsExpires
	}

	return cache.Value, nil
}

func (f FileCache) DelData(ctx context.Context, key string) error {
	err := os.Remove(f.CacheDir + key)
	if err != nil && os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	} else {
		return nil
	}
}

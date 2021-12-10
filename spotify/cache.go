package spotify

import (
	"encoding/gob"
	"fmt"
	"go-music/youtube"
	"os"
	"path/filepath"
	"sync"
)

type Cache struct {
	stored map[string]youtube.Result
	lock   *sync.Mutex
}

func LoadCache() (Cache, error) {
	dir, err := constructDataDirPath()
	if err != nil {
		return Cache{}, err
	}

	f, err := os.Open(filepath.Join(dir, "matchcache"))
	if os.IsNotExist(err) {
		f, err = os.OpenFile(filepath.Join(dir, "matchcache"), os.O_CREATE|os.O_RDWR, 0755)
		if err != nil {
			return Cache{}, err
		}
		defer f.Close()

		return Cache{stored: map[string]youtube.Result{}, lock: &sync.Mutex{}}, nil
	} else if err != nil {
		return Cache{}, fmt.Errorf("error while opening cache file: %v", err)
	}

	items := map[string]youtube.Result{}
	if err = gob.NewDecoder(f).Decode(&items); err != nil {
		f.Close()

		if err := os.Remove(filepath.Join(dir, "matchcache")); err != nil && !os.IsNotExist(err) {
			return Cache{}, fmt.Errorf("error while removing cache file: %v", err)
		}
		f, err := os.OpenFile(filepath.Join(dir, "matchcache"), os.O_CREATE|os.O_RDWR, 0755)
		if err != nil {
			return Cache{}, fmt.Errorf("error while writing cache file: %v", err)
		}
		defer f.Close()

		return Cache{stored: map[string]youtube.Result{}, lock: &sync.Mutex{}}, nil
	}
	return Cache{stored: items, lock: &sync.Mutex{}}, nil
}

func (c *Cache) AddToCache(id string, result youtube.Result) error {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.stored[id] = result

	dir, err := constructDataDirPath()
	if err != nil {
		return err
	}

	if err := os.Remove(filepath.Join(dir, "matchcache")); err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("error while removing cache file: %v", err)
	}
	f, err := os.OpenFile(filepath.Join(dir, "matchcache"), os.O_CREATE|os.O_RDWR, 0755)
	if err != nil {
		return fmt.Errorf("error while writing cache file: %v", err)
	}
	defer f.Close()

	if err = gob.NewEncoder(f).Encode(c.stored); err != nil {
		return fmt.Errorf("error while encoding cache file: %v", err)
	}

	return nil
}

func (c *Cache) Result(id string) *youtube.Result {
	r, ok := c.stored[id]
	if !ok {
		return nil
	}
	return &r
}

func constructDataDirPath() (string, error) {
	dir, err := os.UserCacheDir()
	if err != nil {
		return "", fmt.Errorf("error while getting user cache directory: %v", err)
	}
	dir = filepath.ToSlash(filepath.Join(dir, "go-music"))
	if _, err = os.Stat(dir); os.IsNotExist(err) {
		if err = os.Mkdir(dir, 0755); err != nil {
			return "", fmt.Errorf("error while creating go-music data at %s directory: %v", dir, err)
		}
	} else if err != nil {
		return "", fmt.Errorf("error while retrieving go-music data directory info: %v", err)
	}
	return dir, nil
}

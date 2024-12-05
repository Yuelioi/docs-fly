package cache

import (
	"docsfly/internal/common/constants"
	"docsfly/internal/types"
	"errors"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"gopkg.in/yaml.v3"
)

/*
储存本地必要数据
*/

var LocalCacheID = "localCache"

type localCache struct {
	storage map[string]*types.Group // eg:[cg/other][collections]
}

func NewLocalCache() *localCache {
	return &localCache{
		storage: map[string]*types.Group{},
	}
}

func (i *localCache) ID() string {
	return LocalCacheID
}

// id: cg/other
func (i *localCache) Add(id string, item any) error {
	collection, ok := item.(*types.Document)
	if !ok {
		return errors.New("item is not a collection")
	}

	// 会preload数据, 因此不会出现没有的情况(仅在目录不变动的情况下)
	if group, ok := i.storage[id]; ok {
		group.Documents = append(group.Documents, collection)
	} else {
		group := &types.Group{
			Documents: []*types.Document{},
		}
		group.Documents = append(group.Documents, collection)
	}

	return i.AfterModify(id)
}
func (i *localCache) Update(id string, item any) error {
	collection, ok := item.(*types.Document)
	if !ok {
		return errors.New("item is not a collection")
	}

	if group, ok := i.storage[id]; ok {
		for index, c := range group.Documents {
			if c.Link == collection.Link {
				group.Documents[index] = collection
			}
		}
	}
	return i.AfterModify(id)
}

// 增/删/改后 保存本地数据
func (i *localCache) AfterModify(id string) error {
	if group, ok := i.storage[id]; ok {
		data, err := yaml.Marshal(group)
		if err != nil {
			return err
		}

		folders := strings.Split(id, string(os.PathSeparator))
		folders = slices.Concat([]string{constants.ConfInst.Resource.Documents}, folders, []string{constants.ConfInst.Resource.MetaFile})
		metaFile := filepath.Join(folders...)
		// 文件不存在
		if _, err := os.Stat(metaFile); err == os.ErrNotExist {
			return err
		}

		return os.WriteFile(metaFile, data, os.ModePerm)
	}

	return nil

}

func (i *localCache) Exists(id string) bool {
	if _, ok := i.storage[id]; ok {
		return true
	}

	return false
}

func (i *localCache) Get(id string) (any, error) {
	if item, ok := i.storage[id]; ok {
		return item, nil
	}

	return nil, errors.New("not found")
}

// Preload implements Manager.
func (i *localCache) Preload(data any) error {
	if storage, ok := data.(map[string]*types.Group); ok {
		i.storage = storage
	}
	return nil
}

// Remove implements Manager.
func (i *localCache) Remove(id string, item any) error {
	collection, ok := item.(*types.Document)
	if !ok {
		return errors.New("item is not a collection")
	}

	if group, ok := i.storage[id]; ok {
		for index, c := range group.Documents {
			if c.Link == collection.Link {
				group.Documents = append(group.Documents[:index], group.Documents[index+1:]...)
				i.storage[id] = group
				break
			}
		}
	}

	return i.AfterModify(id)
}

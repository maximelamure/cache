package cache

import (
	"bytes"
	"context"
	"encoding/gob"

	"github.com/bradfitz/gomemcache/memcache"
)

type client struct {
	mc *memcache.Client
}

var (
	c *client
)

// Init initiates a memcache client using the provided server(s)
// with equal weight. If a server is listed multiple times,
// it gets a proportional amount of weight.
func Init(servers ...string) {
	mc := memcache.New(servers...)
	c = &client{mc: mc}
}

// Get gets the item for the given key. The key must be
// at most 250 bytes in length.
func Get(key string, obj interface{}) bool {
	if len(key) > 0 && c != nil {
		item, err := c.mc.Get(key)
		if err != nil {
			return false
		}

		if err := objFromBytes(item.Value, obj); err != nil {
			return false
		}

		return true
	}

	return false
}

func objFromBytes(b []byte, obj interface{}) error {
	dec := gob.NewDecoder(bytes.NewReader(b))
	return dec.Decode(obj)
}

func bytesFromObj(obj interface{}) (bytes.Buffer, error) {
	var b bytes.Buffer
	enc := gob.NewEncoder(&b)
	return b, enc.Encode(obj)
}

// Set writes the given item, unconditionally.
// expiration is the cache expiration time, in seconds: either a relative
// time from now (up to 1 month), or an absolute Unix epoch time.
// Zero means the Item has no expiration time.
func Set(key string, object interface{}, expiration int32) error {
	if c != nil && len(key) > 0 {
		b, err := bytesFromObj(object)
		if err != nil {
			return err
		}

		item := &memcache.Item{
			Key:        key,
			Value:      b.Bytes(),
			Expiration: expiration,
		}

		return c.mc.Set(item)
	}

	return nil
}

func Delete(ctx context.Context, key string) error {
	return c.mc.Delete(key)
}

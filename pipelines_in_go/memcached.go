import (
	"bytes"
	"encoding/gob"
	"os"
	"time"
	"github.com/bradfitz/gomemcache/memcache"
)


type Client struct {
	client *memcache.Client
}

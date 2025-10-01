package interfaces

import (
	"time"
)

type Redis interface {
	Get(key string) (res string, err error)
	Set(key string, value interface{}, expires time.Duration) (err error)
}

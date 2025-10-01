package redis

import "fmt"

func keyGenerator(key string) string {
	return fmt.Sprintf("coin:%s", key)
}

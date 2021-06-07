package services

import (
	"fmt"
)

const (
	defaultCachedDuration  int    = 5000
	defaultCachedNamespace string = "goCms."
)

func Store(key string, obj interface{}) {
	fmt.Println("defs:", defaultCachedDuration, defaultCachedNamespace)
}

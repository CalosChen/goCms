package services

import (
	"fmt"
)

const (
	defaultCachedDuration  = 5000
	defaultCachedNamespace = "goCms."
)

func Store(key string, obj interface{}) {
	fmt.Println("defs:", defaultCachedDuration, defaultCachedNamespace)
}

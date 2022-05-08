package utils

import "hash/fnv"

func hash(s string) uint64 {
	h := fnv.New64()
	h.Write([]byte(s))
	return h.Sum64()
}

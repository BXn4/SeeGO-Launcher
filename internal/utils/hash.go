package utils

import (
	"crypto/md5"
	"encoding/hex"
)

// Returns the MD5 hash of the string (v = value)
func GetHash(v string) string {
	hash := md5.Sum([]byte(v))
	return hex.EncodeToString(hash[:])
}

// Compares the hashes of two string (r = remote, l = local)
func CompareHashes(r, l string) bool {
	remoteHash := GetHash(r)
	localHash := GetHash(l)

	return remoteHash == localHash
}

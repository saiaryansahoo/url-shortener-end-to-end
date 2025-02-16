package utils

import (
    "crypto/md5"
    "encoding/hex"
)

func GenerateShortURL(originalURL string) string {
    hash := md5.Sum([]byte(originalURL))
    return hex.EncodeToString(hash[:])[:8]
}
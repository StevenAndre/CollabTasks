package utils
import (
    "math/rand"
)
const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"


func GeneradorIds(length int) string {
	byteArray := make([]byte, length)
    for i := range byteArray {
        byteArray[i] = charset[rand.Intn(len(charset))]
    }
    return string(byteArray)
}
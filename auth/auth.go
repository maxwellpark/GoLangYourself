package auth

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"
)

func Run() {
	dateStr := getDateStr()
	fmt.Printf("Date string: %s", dateStr)
}

func GetToken(pubKey string, privKey string) {
	dateStr := getDateStr()
	sum := sha256.Sum256([]byte(pubKey + dateStr))
	fmt.Printf("Sum: %s", sum)
}

func getDateStr() string {
	time := time.Now()
	fmt.Printf("\nTime now: %s", time.String())
	formatted := time.Format("2006-01-02 15:04")
	formatted = strings.ReplaceAll(formatted, "-", "")
	formatted = strings.ReplaceAll(formatted, " ", "")
	formatted = strings.ReplaceAll(formatted, ":", "")
	fmt.Printf("\nFormatted time: %s", formatted)
	return formatted
}

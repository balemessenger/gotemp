package pkg

import (
	"encoding/hex"
	"regexp"
)

func ConvertToHex(token []byte) string {
	str := hex.EncodeToString(token)
	var re = regexp.MustCompile("[^a-fA-F0-9]")
	return re.ReplaceAllString(str, "")
}

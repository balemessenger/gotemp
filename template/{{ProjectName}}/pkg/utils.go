package pkg

import (
	"encoding/hex"
	"regexp"
	"sync"
)

var (
	utilsOnce sync.Once
	utils     *Utils
)

type Utils struct {
}

func NewUtils() *Utils {
	return &Utils{}
}

func GetUtils() *Utils {
	utilsOnce.Do(func() {
		utils = NewUtils()
	})
	return utils
}

func (Utils) ConvertToHex(token []byte) string {
	str := hex.EncodeToString(token)
	var re = regexp.MustCompile("[^a-fA-F0-9]")
	return re.ReplaceAllString(str, "")
}

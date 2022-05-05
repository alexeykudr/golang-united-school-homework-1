package solution

import (
	"github.com/kyokomi/emoji/v2"
)

func GetMessage() string {

	heyStr := emoji.Sprint("Hello :world_map:!")
	return heyStr
}

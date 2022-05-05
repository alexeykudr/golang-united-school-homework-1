package solution

import "github.com/kyokomi/emoji/v2"

func GetMessage() string {
	hey_str := emoji.Sprintf("Hello :world_map:!")
	return hey_str
}

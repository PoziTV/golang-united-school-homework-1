package solution

import (
	"github.com/kyokomi/emoji"
)

func GetMessage() string {
	const worldMap = ":world_map:"
	out := emoji.Sprint(worldMap)

	return out
}

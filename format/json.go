package format

import (
	"encoding/json"
	"fmt"

	"github.com/whywaita/gigani/parse"
)

func JSON(animes []parse.Anime) string {
	j, err := json.Marshal(animes)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return string(j)
}

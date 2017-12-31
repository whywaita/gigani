package output

import (
	"encoding/json"
	"fmt"

	"github.com/whywaita/gigani/parse"
)

func JSON(animes []parse.Anime) {
	j, err := json.Marshal(animes)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(j))
}

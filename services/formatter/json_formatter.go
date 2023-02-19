package formatter

import (
	"encoding/json"
	"fmt"
	"github.com/roadsigns/httpct/services"
)

const JSON_FORMATTER = "json"

type JsonFormatter struct {
}

func (j JsonFormatter) Format(results services.HttpHeaderScanResults) {
	b, err := json.MarshalIndent(results, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}

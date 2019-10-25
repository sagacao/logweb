package elastic

import (
	"encoding/json"

	"github.com/henrylee2cn/faygo"
)

func dumpQueryByJson(query map[string]interface{}) {
	faygo.Debugf("========== dumpQueryByJson ============")
	data, err := json.Marshal(query)
	if err != nil {
		faygo.Warningf("err: %v", err)
	} else {
		faygo.Debugf("data: %v", string(data))
	}
	faygo.Debugf("========== dumpQueryByJson ============")
}

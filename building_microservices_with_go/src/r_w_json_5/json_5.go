package r_w_json_5

import (
	"net/http"
	"encoding/json"
)

// 性能问题　
// Encoder Decoder 减少中间变量，减少代码　直接使用流处理

func helloWorldHandler(w http.ResponseWriter,r *http.Request){
	var request helloWorldRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil{
		http.Error(w,"Bad request " ,http.StatusBadRequest)
		return
	}

	response := helloWorldResponse{Message:"Hello " + request.Name}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
}

type helloWorldRequest struct{
	Name string `json:"name"`
}
type helloWorldResponse struct{
	Message string `json:"message"`
}
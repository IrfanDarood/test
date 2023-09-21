package web

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
)

// Query handles chaincode query requests.
func (setup OrgSetup) Query(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received Query request")
	queryParams := r.URL.Query()
	chainCodeName := queryParams.Get("chaincodeid")
	channelID := queryParams.Get("channelid")
	function := queryParams.Get("function")
	// args := r.URL.Query()["args"]


	var reqBody = new(ReqBody)
	b,err:=io.ReadAll(r.Body);
	// fmt.Println(string(b))

	json.Unmarshal(b,reqBody)
	// fmt.Println(reqBody)
	// fmt.Printf("channel: %s, chaincode: %s, function: %s, args: %s\n", channelID, chainCodeName, function, args)
	network := setup.Gateway.GetNetwork(channelID)
	contract := network.GetContract(chainCodeName)
	evaluateResponse, err := contract.EvaluateTransaction(function, reqBody.Args...)
	// fmt.Println(string(evaluateResponse), " Byte " ,  evaluateResponse)
	var resBody = new(Response)
	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println("err.Error() ",err.Error())
		// resBody.Status = 500
		// resBody.Message = err.Error()
		json.NewEncoder(w).Encode(resBody)
		return
	}
	if len(evaluateResponse) == 0 {
		fmt.Println(string(b))
		// resBody.Status = 404
		// resBody.Message = "Not found"
		json.NewEncoder(w).Encode(resBody)
		return 
	}

	// resBody.Status = 200
	// resBody.Message = string(evaluateResponse)
	json.NewEncoder(w).Encode(resBody)
	return
}


type ReqBody struct{
	Args []string
	Transient map[string][]byte//interface{}
}
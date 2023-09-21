package web

import (
	"fmt"
	"net/http"
	"encoding/json"
	// "context"
	//"io"
	// "math"
	"time"
	// "strconv"
	"log"
	//"encoding/base64"
//	"github.com/google/uuid"

	"github.com/hyperledger/fabric-gateway/pkg/client"
)

// Invoke handles chaincode invoke requests.
func (setup *OrgSetup) Invoke(w http.ResponseWriter, r *http.Request) {
	timeStamp := make(map[string]time.Duration)
	var asset Asset
	err:=json.NewDecoder(r.Body).Decode(&asset)
	// var assetId = uuid.New().String()
	// asset.AssetID=assetId
	fmt.Println("AssetId is", asset.AssetID)
	
	// jsonBytes, err := json.Marshal(asset)
	// var asset_Data = map[string][]byte{
	// 	"asset_properties": jsonBytes,
	// }

	fmt.Println("Received Invoke request")
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %s", err)
		return
	}
	chainCodeName := r.FormValue("chaincodeid")
	channelID := r.FormValue("channelid")
	function := r.FormValue("function")
	args := r.Form["args"]
	
	fmt.Printf("channel: %s, chaincode: %s, function: %s, args: %s\n", channelID, chainCodeName, function, args)
	network := setup.Gateway.GetNetwork(channelID)
	contract := network.GetContract(chainCodeName)
	
	txn_proposal, err := contract.NewProposal(function,client.WithArguments("one", "two"))
	if err != nil {
		fmt.Fprintf(w, "Error creating txn proposal: %s", err)
		return
	}
	txn_endorsed, err := txn_proposal.Endorse()
	if err != nil {
		fmt.Fprintf(w, "Error endorsing txn: %s", err)
		return
	}
	startTime:=time.Now()
	txn_committed, err := txn_endorsed.Submit()
	t:=time.Now()
	if err != nil {
		fmt.Fprintf(w, "Error submitting transaction: %s", err)
		return
	}
	
	
	endTime:=t.Sub(startTime)
	// fmt.Println("Start time is ",startTime)
	// fmt.Println("end` time is ",t)
	fmt.Println("elapsed time",endTime)
	
	// fmt.Println("elapsed time ",float64(endTime))
	
	// rounded := math.Round(endTime*100) / 100
	// fmt.Println("Rounded",rounded)
	timeStamp[txn_endorsed.TransactionID()]=endTime
	fmt.Print(txn_committed);
	
	var sum time.Duration
	count:=0
	for _,j := range timeStamp{
		count++
		// slice := j[:len(j)-2]
		
		// t, err := strconv.ParseFloat(slice,64)
		if err != nil {
			// ... handle error
				fmt.Print("error")
				panic(err)
		}

		sum+=j
	


	}
	
	if err != nil {
		fmt.Println("err.Error() ",err.Error())
		
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]time.Duration)
	resp["sum"] = sum
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return 




}



type Asset struct {
	ObjectType  string  `json:"objectType"`
	AssetKey    string  `json:"assetKey"`
	AssetID     string  `json:"assetID"`
	PrevAssetID string  `json:"prevAssetID"`
	Asset       string  `json:"asset"`
	Qty         float64 `json:"qty"`
	Owner       string  `json:"owner"`
	Active      string  `json:"active"`
	Version     int     `json:"version"`
	NewOwner  string `json:"new_owner"`
}



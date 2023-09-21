package web

import (
	"fmt"
	"net/http"


	"github.com/hyperledger/fabric-gateway/pkg/client"
)

// OrgSetup contains organization's config to interact with the network.
type OrgSetup struct {
	OrgName      string
	MSPID        string
	CryptoPath   string
	CertPath     string
	KeyPath      string
	TLSCertPath  string
	PeerEndpoint string
	GatewayPeer  string
	Gateway      client.Gateway
	Port 		string
}

// Serve starts http web server.
func Serve(setups OrgSetup) {
/*c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3001"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type"},
	})

*/
	http.HandleFunc("/query", setups.Query)
	http.HandleFunc("/invoke", setups.Invoke)
	http.HandleFunc("/transfer",setups.Transfer)
	
	// handler := c.Handler(http.DefaultServeMux)

	
	fmt.Println("Listening (http://localhost:3000/)...peer0 org")
	if err := http.ListenAndServe(fmt.Sprintf(":%s",setups.Port),nil); err != nil {
		fmt.Println(err)
	}
}

package main

import (
	"fmt"
	"rest-api-go/web"
	"os"
)

var Peer string= os.Getenv("PEER")
var Org string=os.Getenv("ORG")
var Port string=os.Getenv("PORT")
func main() {
	fmt.Print(Peer,Org)

	//Initialize setup for Org1
	// cryptoPath := fmt.Sprintf("/data/deployment/crypto-config/peerOrganizations/%s",Org)
	//Initialize setup for Org1
	cryptoPath := fmt.Sprintf("/data/deployment/crypto-config/peerOrganizations/%s",Org) //@@change path of peerOrgs
	//cryptoPath := "../../Change_HyoerLedger_peer/f-sample/vanilla-test-network/organizations/peerOrganizations/org2.example.com"
	orgConfig := web.OrgSetup{
		OrgName:      "manufacturer",
		MSPID:        "manufacturerMSP",
		CertPath:     cryptoPath + fmt.Sprintf("/users/Admin@%s/msp/signcerts/cert.pem",Org),
		KeyPath:      cryptoPath + fmt.Sprintf("/users/Admin@%s/msp/keystore/",Org),
		TLSCertPath:  cryptoPath + fmt.Sprintf("/peers/%s-%s/tls/ca.crt",Peer,Org),
		PeerEndpoint: fmt.Sprintf("%s-%s:7051",Peer,Org),
		GatewayPeer:  fmt.Sprintf("%s-%s",Peer,Org),
		Port:Port,
	}

	orgSetup, err := web.Initialize(orgConfig)
	if err != nil {
		fmt.Sprintf("Error initializing setup for Peer0-icici: ", err)
	}
	web.Serve(web.OrgSetup(*orgSetup))
}

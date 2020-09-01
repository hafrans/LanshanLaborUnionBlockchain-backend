package smsrpc

import (
	"context"
	"crypto/x509"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"time"
)
import pb "RizhaoLanshanLabourUnion/services/smsqueue/proto"

var rpcServer = "rpc.node-1.hafrans.com:65001"

var conn *grpc.ClientConn

func init() {
	var err error
	conn, err = initializeConnection()
	if err != nil {
		log.Panic("rpc server failed!")
	}
}

func initializeConnection() (*grpc.ClientConn, error) {

	cp := x509.NewCertPool()
	cp.AppendCertsFromPEM([]byte(`
-----BEGIN CERTIFICATE-----
MIIE8zCCA9ugAwIBAgIBATANBgkqhkiG9w0BAQsFADCBqDELMAkGA1UEBhMCQ04x
EDAOBgNVBAgTB0JlaWppbmcxGTAXBgNVBAcTEEhhaWRpYW4gRGlzdHJpY3QxEjAQ
BgNVBAoTCUNodXVrYSBSbzEhMB8GA1UECxMYU2VjdXJpdHkgTWFuYWdlbWVudCBV
bml0MRUwEwYDVQQDEwxDaHV1a2EgUm8gQ0ExHjAcBgkqhkiG9w0BCQEWD2tleUBo
YWZyYW5zLmNvbTAgFw0xOTEyMTYwODM0MDBaGA8yMDU1MTIxNjA4MzQwMFowgagx
CzAJBgNVBAYTAkNOMRAwDgYDVQQIEwdCZWlqaW5nMRkwFwYDVQQHExBIYWlkaWFu
IERpc3RyaWN0MRIwEAYDVQQKEwlDaHV1a2EgUm8xITAfBgNVBAsTGFNlY3VyaXR5
IE1hbmFnZW1lbnQgVW5pdDEVMBMGA1UEAxMMQ2h1dWthIFJvIENBMR4wHAYJKoZI
hvcNAQkBFg9rZXlAaGFmcmFucy5jb20wggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAw
ggEKAoIBAQC8P+mqEKJ+hxhBaxSIBh6sG9RgK+U+22scxsd9PF+hi1TRmkD9hOVU
+vgHarbw1upUCiWDuEECZ1ORXcgcZrTRIeqUzwBaPGoW8HNl/PTAoMDVCEAu2HnZ
J2pZyoDorknlugD0KmN77Tr1Mmwzygab+/ppzgf0OVhXuUgkvaIj35fyvl/7JoGN
8/ss9GZ3YswUGuoDaFdNSpU76u071Tn2UU4hWEgx8xlQRNRJI6DIu8Phy+7mi9+E
u5sBRrBkGY/Xo1ZXG2VKwINm/m4tB5bFepcBVrs1LMkVnf4MfOKfXqhQXUxoQRxB
LcuOcur2Npz49XUUyLJZmZlKVSUjPkhNAgMBAAGjggEiMIIBHjAMBgNVHRMEBTAD
AQH/MAwGA1UdDwQFAwMH/4AwgewGA1UdJQSB5DCB4QYIKwYBBQUHAwEGCCsGAQUF
BwMCBggrBgEFBQcDAwYIKwYBBQUHAwQGCCsGAQUFBwMIBgorBgEEAYI3AgEVBgor
BgEEAYI3AgEWBgorBgEEAYI3CgMBBgorBgEEAYI3CgMDBgorBgEEAYI3CgMEBglg
hkgBhvhCBAEGCysGAQQBgjcKAwQBBggrBgEFBQcDBQYIKwYBBQUHAwYGCCsGAQUF
BwMHBggrBgEFBQgCAgYKKwYBBAGCNxQCAgYIKwYBBQUHAwkGCCsGAQUFBwMNBggr
BgEFBQcDDgYHKwYBBQIDBTARBglghkgBhvhCAQEEBAMCAPcwDQYJKoZIhvcNAQEL
BQADggEBAKPSxudslJ5yX/+A56iiCewd1NsFaqCwa1xOD3JQpW1TER+BlulJC3iK
dYI4mypYBY5qX1oVm8IQ9oZu+DRAZroW2VJNGOduZgp57Nul7M7vjQaTyEUB7Sww
wlWixGSEZCnWz6M2EEZf2q6CdYnVJn83gHm4PQiDVfF1twM8BlEJOxpTWhsJo+el
5QKOV/uv06XwbVNFNQ+tsFR3TAjKK35/uscRbWcPNF8qz0EDJK5mNyXQCG1wyIoQ
ZBpWofkVBwACFhC2SO5tnibQWO46Nps8cOmItjkJFGNL9qePtXi9APpU3cTCuggB
+F0EEE9ix5uJXkJBaq6XzMxUG1S9lVY=
-----END CERTIFICATE-----
`))
	cred := credentials.NewClientTLSFromCert(cp, "rpc.node-1.hafrans.com")

	conn, err := grpc.Dial(rpcServer, grpc.WithTransportCredentials(cred), grpc.WithBlock())

	if err != nil {
		return nil, err
	} else {
		return conn, nil
	}
}

func GetProtoBuffClient(conn *grpc.ClientConn) pb.UnicomMessagePushClient {
	return pb.NewUnicomMessagePushClient(conn)
}

func CloseConn() error {
	if conn != nil {
		if err := conn.Close(); err != nil {
			return err
		}
	}
	return nil
}

func sendRequest(client pb.UnicomMessagePushClient, request *pb.SendRequest) (*pb.SendResponse, error) {

	if client == nil {
		return nil, nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	resp, err := client.SendMessage(ctx, request)
	if err != nil {
		return nil, err
	} else {
		return resp, err
	}

}

func SendMessage(account, password, mobile, content string) {
	sendRequest(GetProtoBuffClient(conn), &pb.SendRequest{Phone: mobile, Content: content, Password: password, Account: account})
}

package sdkop

import (
	"chainmaker.org/chainmaker/common/v2/crypto"
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"chainmaker.org/chainmaker/sdk-go/v2/examples"
	"encoding/json"
	"fmt"
	"strings"
	"unsafe"

	//"chainmaker.org/chainmaker-sdk-go/pb/protogo/common"
	"log"
)

var (
	Client1 *sdk.ChainClient
	Client2 *sdk.ChainClient
) 

// func init(){
// 	Client1, err :=createClientWithConfig(sdkConfigPath)
// 	if err!=nil{
// 		fmt.Println(err)
// 		return
// 	}
// 	Client2, err := createClientWithConfig(sdkConfigPath2)
// 	if err!=nil{
// 		fmt.Println(err)
// 		return
// 	}
// }


func createNode(nodeAddr string, connCnt int) *sdk.NodeConfig {
	node := sdk.NewNodeConfig(
		// 节点地址，格式：127.0.0.1:12301
		sdk.WithNodeAddr(nodeAddr),
		// 节点连接数
		sdk.WithNodeConnCnt(connCnt),
		// 节点是否启用TLS认证
		sdk.WithNodeUseTLS(true),
		// 根证书路径，支持多个
		sdk.WithNodeCAPaths(caPaths),
		// TLS Hostname
		sdk.WithNodeTLSHostName(tlsHostName),
	)

	return node
}

/*
用配置文件的方式创建连接
*/
func createClientWithConfig(sdk_conf_path string) (*sdk.ChainClient, error) {
	/*
	, sdk.WithUserKeyFilePath(clientKeyFilePaths),
		sdk.WithUserCrtFilePath(clientCrtFilePaths), sdk.WithChainClientOrgId(orgId), sdk.WithChainClientChainId(chainId)
	*/
	chainClient, err := sdk.NewChainClient(sdk.WithConfPath(sdk_conf_path), sdk.WithChainClientLogger(getDefaultLogger()))
	if err != nil {
		return nil, err
	}

	return chainClient, nil
}


func createAdminWithConfig(orgId string) (*sdk.ChainClient, error) {
	//fmt.Println(adminKeyPath, orgId)
	//fmt.Println(adminCrtPath, orgId)
	adminClient, err := sdk.NewChainClient(
		sdk.WithConfPath(sdkConfigPath),
		sdk.WithUserKeyFilePath(fmt.Sprintf(adminKeyPath, orgId)),
		sdk.WithUserCrtFilePath(fmt.Sprintf(adminCrtPath, orgId)),
	)
	if err != nil {
		return nil, err
	}

	//启用证书压缩（开启证书压缩可以减小交易包大小，提升处理性能）
	err = adminClient.EnableCertHash()
	if err != nil {
		return nil, err
	}

	return adminClient, nil
}

//func createUserContract(client *sdk.ChainClient, admin1, admin2, admin3, admin4 *sdk.ChainClient,
//	contractName, version, byteCodePath string, runtime common.RuntimeType, kvs []*common.KeyValuePair, withSyncResult bool) (*common.TxResponse, error) {
//	//fmt.Printf("h\n");
//	payloadBytes, err := client.CreateContractCreatePayload(contractName, version, byteCodePath, runtime, kvs)
//	if err != nil {
//		return nil, err
//	}
//	//fmt.Printf("PayloadBytes: %+v\n",payloadBytes)
//
//	// 各组织Admin权限用户签名
//	signedPayloadBytes1, err := admin1.SignContractManagePayload(payloadBytes)
//	if err != nil {
//		return nil, err
//	}
//
//	signedPayloadBytes2, err := admin2.SignContractManagePayload(payloadBytes)
//	if err != nil {
//		return nil, err
//	}
//
//	signedPayloadBytes3, err := admin3.SignContractManagePayload(payloadBytes)
//	if err != nil {
//		return nil, err
//	}
//
//	signedPayloadBytes4, err := admin4.SignContractManagePayload(payloadBytes)
//	if err != nil {
//		return nil, err
//	}
//
//	// 收集并合并签名
//	mergeSignedPayloadBytes, err := client.MergeContractManageSignedPayload([][]byte{signedPayloadBytes1,
//		signedPayloadBytes2, signedPayloadBytes3, signedPayloadBytes4})
//	if err != nil {
//		return nil, err
//	}
//
//	// 发送创建合约请求
//	resp, err := client.SendContractManageRequest(mergeSignedPayloadBytes, createContractTimeout, withSyncResult)
//	if err != nil {
//		return nil, err
//	}
//
//	err = sdk.CheckProposalRequestResp_ext(resp, true)
//	if err != nil {
//		return nil, err
//	}
//
//	return resp, nil
//}


func createUserContract(client *sdk.ChainClient, admin1, admin2, admin3, admin4 string,
	contractName, version, byteCodePath string, runtime common.RuntimeType, kvs []*common.KeyValuePair,
	withSyncResult bool) (*common.TxResponse, error) {

	payload, err := client.CreateContractCreatePayload(contractName, version, byteCodePath, runtime, kvs)
	if err != nil {
		return nil, err
	}

	//endorsers, err := examples.GetEndorsers(payload, admin1, admin2, admin3, admin4)
	endorsers, err := examples.GetEndorsersWithAuthType(crypto.HashAlgoMap[client.GetHashType()],
		client.GetAuthType(), payload, admin1, admin2, admin3, admin4)
	if err != nil {
		return nil, err
	}

	resp, err := client.SendContractManageRequest(payload, endorsers, createContractTimeout, withSyncResult)
	if err != nil {
		return nil, err
	}

	// TODO: ??
	//err = examples.CheckProposalRequestResp(resp, true)
	//if err != nil {
	//	return nil, err
	//}

	return resp, nil
}


//func testUserContractAssetCreate(client *sdk.ChainClient, kvs []*common.KeyValuePair, withSyncResult bool,
//	isIgnoreSameContract bool, usernames ...string) {
//
//	resp, err := createUserContract(client, assetContractName, assetVersion, assetByteCodePath,
//		common.RuntimeType_WASMER, kvs, withSyncResult, usernames...)
//	if !isIgnoreSameContract {
//		if err != nil {
//			log.Fatalln(err)
//		}
//	}
//
//	fmt.Printf("CREATE asset contract resp: %+v\n", resp)
//}


func userContractAssetCreate(client *sdk.ChainClient,
	admin1, admin2, admin3, admin4 string, kvs []*common.KeyValuePair, withSyncResult bool, isIgnoreSameContract bool) {

	resp, err := createUserContract(client, admin1, admin2, admin3, admin4,
		assetContractName, assetVersion, assetByteCodePath, common.RuntimeType_WASMER, kvs, withSyncResult)
	if !(isIgnoreSameContract && err!=nil&& strings.Contains(err.Error(),"Already")) {
		if nil!=err{
			fmt.Printf("%+v %T\n",err,err)
		}
	}

	fmt.Printf("CREATE asset contract resp: %+v\n", resp)
}

/*
调用智能合约
*/
//func invokeUserContract(client *sdk.ChainClient, contractName, method, txId string, params map[string]string, withSyncResult bool) error {
//
//	resp, err := client.InvokeContract(contractName, method, txId, params, -1, withSyncResult)
//	if err != nil {
//		return err
//	}
//	//fmt.Printf("调用完智能合约： %+v\n",resp)
//	if resp.Code != common.TxStatusCode_SUCCESS {
//		return fmt.Errorf("invoke contract failed, [code:%d]/[msg:%s]\n", resp.Code, resp.Message)
//	}
//
//	if !withSyncResult {
//		//fmt.Printf("invoke contract success, resp: [code:%d]/[msg:%s]/[txId:%s]\n", resp.Code, resp.Message, resp.ContractResult.Result)
//	} else {
//		//fmt.Printf("invoke contract success, resp: [code:%d]/[msg:%s]/[contractResult:%s]\n", resp.Code, resp.Message, resp.ContractResult)
//	}
//
//	return nil
//}

func invokeUserContract(client *sdk.ChainClient, contractName, method, txId string,
	kvs []*common.KeyValuePair, withSyncResult bool) (string, error) {

	resp, err := client.InvokeContract(contractName, method, txId, kvs, 10, withSyncResult)
	if err != nil {
		return "", err
	}

	if resp.Code != common.TxStatusCode_SUCCESS {
		return "", fmt.Errorf("invoke contract failed, [code:%d]/[msg:%s]\n", resp.Code, resp.Message)
	}

	if !withSyncResult {
		fmt.Printf("invoke contract success, resp: [code:%d]/[msg:%s]/[txId:%s]\n", resp.Code, resp.Message, resp.TxId)
	} else {
		fmt.Printf("invoke contract success, resp: [code:%d]/[msg:%s]/[contractResult:%s]\n", resp.Code, resp.Message, resp.ContractResult)
	}

	return resp.TxId, err
}

func userContractAssetInvokeRegister(client *sdk.ChainClient, method string, withSyncResult bool) (string, error) {
	txid, err := invokeUserContract(client, assetContractName, method, "", nil, withSyncResult)
	return txid, err
}

func userContractAssetInvoke(client *sdk.ChainClient, name, method, args, amount, addr string, withSyncResult bool) (string, error) {

	m := make(map[string]string)
	err := json.Unmarshal([]byte(args), &m)
	if err != nil {
		return "", err
	}
	kvs := []*common.KeyValuePair{}

	for k,v := range m {
		kvs = append(kvs, &common.KeyValuePair{
			Key: k,
			Value: *(*[]byte)(unsafe.Pointer(&v)),
		})
	}

	txid, err := invokeUserContract(client, name, method, "", kvs, withSyncResult)

	return txid, err
}

func testUserContractAssetInvoke(client *sdk.ChainClient, method string, amount, addr string, withSyncResult bool) {
	kvs := []*common.KeyValuePair{
		{
			Key:   "amount",
			Value: []byte(amount),
		},
		{
			Key:   "to",
			Value: []byte(addr),
		},
	}

	txid, err := invokeUserContract(client, assetContractName, method, "", kvs, withSyncResult)
	fmt.Println(txid)
	if err != nil {
		log.Fatalln(err)
	}
}

func userContractAssetQuery(method string, kvs []*common.KeyValuePair) string {
	resp, err := Client1.QueryContract(assetContractName, method, kvs, -1)
	
	if err!=nil{
		fmt.Printf("error: %v\n",err)
		return ""
	}
	fmt.Printf("QUERY asset contract [%s] resp: %+v\n", method, resp)

	err = examples.CheckProposalRequestResp(resp, true)
	if err!=nil{
		fmt.Printf("error: %v\n",err)
		return ""
	}
	return string(resp.ContractResult.Result)
}


//func getBalance(addr string) int32{
//	params := map[string]string{
//		"owner": addr,
//	}
//	balance := userContractAssetQuery("balance_of", params)
//	val, err := sdk.BytesToInt([]byte(balance), binary.LittleEndian)
//	if err!=nil{
//		fmt.Printf("error: %v\n",err)
//		return -1
//	}
//	fmt.Printf("client [%s] balance: %d\n", addr, val)
//	return val
//}

func getBalance(client *sdk.ChainClient, addr string) {
	kvs := []*common.KeyValuePair{
		{
			Key:   "owner",
			Value: []byte(addr),
		},
	}

	balance := testUserContractAssetQuery(client, "balance_of", kvs)

	fmt.Printf("client [%s] balance: %s\n", addr, balance)
}

func testUserContractAssetQuery(client *sdk.ChainClient, method string, kvs []*common.KeyValuePair) string {
	resp, err := client.QueryContract(assetContractName, method, kvs, -1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("QUERY asset contract [%s] resp: %+v\n", method, resp)

	err = examples.CheckProposalRequestResp(resp, true)
	if err != nil {
		log.Fatalln(err)
	}
	return string(resp.ContractResult.Result)
}

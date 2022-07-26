package sdkop

import (
	"chainmaker.org/chainmaker/pb-go/v2/common"
	sdk "chainmaker.org/chainmaker/sdk-go/v2"
	"chainmaker.org/chainmaker/sdk-go/v2/examples"
	"encoding/json"
	"fmt"
	"unsafe"
)

func Connect_chain(tp int, sdkConfigPath string) *sdk.ChainClient{
	
	chainClient, err :=createClientWithConfig(sdkConfigPath)
	//if 2==tp{
	//	chainClient, err =createClientWithConfig(sdkConfigPath2)
	//}
	if err!=nil{
		fmt.Println(err)
		return nil
	}
	return chainClient
}

/*
实例化代币合约
*/
func ContractInstance(){
	chainClient, err :=createClientWithConfig(sdkConfigPath)
	if err!=nil{
		fmt.Println(err)
		return
	}
	Client1=chainClient
	//fmt.Printf("%v\n",chainClient)
	//fmt.Println("====================== create admin1 ======================")
	//admin1, err := createAdminWithConfig(orgNodeId1)
	//if err!=nil{
	//	fmt.Println(err)
	//	return
	//}
	////fmt.Println("====================== create admin2 ======================")
	//admin2, err := createAdminWithConfig(orgNodeId2)
	//if err!=nil{
	//	fmt.Println(err)
	//	return
	//}
	////fmt.Println("====================== create admin3 ======================")
	//admin3, err := createAdminWithConfig(orgNodeId3)
	//if err!=nil{
	//	fmt.Println(err)
	//	return
	//}
	////fmt.Println("====================== create admin4 ======================")
	//admin4, err := createAdminWithConfig(orgNodeId4)
	//if err!=nil{
	//	fmt.Println(err)
	//	return
	//}

	fmt.Println("====================== 1)安装钱包合约 ======================")
	pairs := []*common.KeyValuePair{
		{
			Key:   "issue_limit",
			Value: []byte("100000000"),
		},
		{
			Key:   "total_supply",
			Value: []byte("100000000"),
		},
	}
	userContractAssetCreate(chainClient, examples.UserNameOrg1Admin1, examples.UserNameOrg2Admin1, examples.UserNameOrg3Admin1, examples.UserNameOrg4Admin1, pairs, true, true)
	fmt.Println("====================== 安装钱包合约成功 ======================")
}

/*---------------上边是通用的，下边和智能合约相关----------------*/
func RegisterUser(){
	client2, err := createClientWithConfig(sdkConfigPath2)
	if err!=nil{
		fmt.Println(err)
		return
	}
	Client2=client2
	fmt.Println("====================== 2)注册另一个用户 ======================")
	userContractAssetInvokeRegister(client2, "register", true)
   	fmt.Println("====================== 注册另一个用户成功 ======================")
   
}	
//func UserContractAssetQuery(client1,client2 *sdk.ChainClient, id bool,  name, method, args string) string {
func UserContractAssetQuery(client *sdk.ChainClient, id bool,  name, method, args string) string {
	/*
	client *sdk.ChainClient, method string, params map[string]string
	*/
	//method:="query_address"
	//var params map[string]string
	m := make(map[string]string)
	json.Unmarshal([]byte(args), &m)
	kvs := []*common.KeyValuePair{}

	for k,v := range m {
		//if reflect.TypeOf(v) == bool
		//switch v

		//if !unicode.IsDigit(*(*rune)(unsafe.Pointer(&v))) {
		//num, ok := strconv.ParseInt(v,10, 64)
		//if ok == nil {
		//	kvs = append(kvs, &common.KeyValuePair{
		//		Key: k,
		//		Value: *(*[]byte)(unsafe.Pointer(&num)),
		//	})
		//} else {
		//	kvs = append(kvs, &common.KeyValuePair{
		//		Key: k,
		//		Value: *(*[]byte)(unsafe.Pointer(&v)),
		//	})
		//}
		//}
		kvs = append(kvs, &common.KeyValuePair{
			Key: k,
			Value: *(*[]byte)(unsafe.Pointer(&v)),
		})

	}
	resp, err := client.QueryContract(name, method, kvs, 10)
	//if id {
	//	resp, err = client1.QueryContract(name, method, kvs, 10)
	//}

	if err!=nil{
		fmt.Printf("get error: %+v\n",err)
		return ""
	}
	//fmt.Printf("QUERY asset contract [%s] resp: %+v\n", method, resp)

	err = examples.CheckProposalRequestResp(resp, true)
	if err!=nil{
		fmt.Printf("check get error: %+v\n",err)
		return ""
	}
	return string(resp.ContractResult.Result)
}


func UserContractAssetInvoke(client *sdk.ChainClient, name, method, args, amount, addr string, withSyncResult bool){
	err:=userContractAssetInvoke(client, name, method, args, amount,addr,withSyncResult)
	if err!=nil{
		fmt.Printf("invoke error : %v\n",err);
	}
}

func GetBalance(client *sdk.ChainClient, addr string) {
	getBalance(client, addr)
}
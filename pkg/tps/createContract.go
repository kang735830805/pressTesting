package tps

import "chainpress/pkg/sdkop"

func Init(){
	//1. 注册转账智能合约
	sdkop.ContractInstance()
	//2. 注册新用户
	sdkop.RegisterUser()
}

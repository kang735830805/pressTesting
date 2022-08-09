module chainpress

go 1.16

require (
	chainmaker.org/chainmaker-sdk-go v1.2.6
	chainmaker.org/chainmaker/common/v2 v2.2.1
	//chainmaker.org/chainmaker-sdk-go v1.2.3
	github.com/panjf2000/ants/v2 v2.5.0
	github.com/spf13/cobra v1.1.1
	github.com/stretchr/testify v1.7.2 // indirect
	go.uber.org/zap v1.17.0
)

//chainmaker.org/chainmaker-sdk-go => ./chainmaker-sdk-go
replace chainmaker.org/chainmaker-sdk-go => ../chainmaker-sdk-go

replace chainmaker.org/chainmaker-go/common => ../chainmaker-sdk-go/common

replace chainmaker.org/chainmaker-sdk-go/common/log => ../chainmaker-sdk-go/common/log

//replace chainmaker.org/chainmaker/sdk-go/v2 => ./chainmaker-sdk-go

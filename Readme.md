
##TPS
|  longname   |  shortname   | description  | examples  |
|  ----  | ----  | ----  | ----  |
| threadNum  | t | 进程数量 | 1000 |
| loopNum  | c | 循环数量 | 100 |
| name  | n | 合约名称 | test |
| method  | m | 合约内方法 | save |
| parameter  | p | 合约参数 | "{\"file_name\":\"name007\"}" |
| sdkPath  | s | SdkConfig路径 | ./sdk_config.yml |

####tps命令行Examples：
```shell
./main tps -l 1 -t 1 -n fact -m save -p "{\"file_name\":\"name007\",\"file_hash\":\"ab3456df5799b87c77e7f88\",\"time\":\"65432ç34\"}" -s ./sdk_config.yml
```


##QPS
|  longname   |  shortname   | description  | examples  |
|  ----  | ----  | ----  | ----  |
| threadNum  | t | 进程数量 | 100 |
| loopNum  | c | 循环数量 | 100 |
| txId  | t | 长安链内的交易txId | 1705f56583e6cd78ca18426f40000b933b8842529c9346b49f4bbcae4b57a57e |
| sdkPath  | s | SdkConfig路径 | ./sdk_config.yml |

####qps命令行Examples：
```shell
#qps 测试
./main qps -l 1 -t 1 -i 'f17c89e559e44f99889be0448e6a0704c2e7b505ce644c2ca5022942a9f9ac31' -s "./sdk_config.yml"
```

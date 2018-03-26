基于点融提供的Node部署版本中迁移至Go。

由于点融未提供智能合约源码，只能基于bin和abi来生成。

## eth:

* 直接引用即可

## bcos:

* 需设置项目的vendor：./vendor/github.com/ethereum/go-ethereum 引用 [gitlab.chainedfinance.com/cf\_bcos/go-ethereum](http://gitlab.chainedfinance.com/cf_bcos/go-ethereum)
* *推荐方式*

`git submodule add -b bcos_v1.8.1 git@gitlab.chainedfinance.com:cf_bcos/go-ethereum.git ./vendor/github.com/ethereum/go-ethereum`[](http://gitlab.chainedfinance.com、cf_bcos/go-ethereum)
# consensus
意在集成各种共识模块，初始版本fork Tendermint。
Tendermint 的模块化使得实现区块链各种可插拔，清晰明朗。
该版本会在Tendermint的基础上去平行的一些共识算法。
比如： 
1. 无共识模块，只有p2p的节点间互联。
2. pow的共识。
3. raft的共识。
4. 原有的Tendermint共识。

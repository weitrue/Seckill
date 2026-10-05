源包地址：
github.com/cbergoon/merkletree

基本符合生成merkle的功能，主要的问题是合约中的hash需要排序，所以需要修改。将该包内的代码复制到了basetree.go中，并进行修改
package core

import (
	"fmt"
	"github.com/urfave/cli"
	"strconv"
)

type BlockChain struct {
	chain []*Block
	txpool []*Transaction
}

var bc BlockChain

//初始化区块链的创世纪块
func init(){
	bc.chain = append(bc.chain,&Block{1})
	bc.chain = append(bc.chain,&Block{2})
	bc.chain = append(bc.chain,&Block{3})
	bc.chain = append(bc.chain,&Block{4})
	bc.chain = append(bc.chain,&Block{5})
}


//展示区块链
func ShowBlockchain(ctx *cli.Context){
	bc.String()
}

//根据区块number展示区块
func ShowBlockByNumber(ctx *cli.Context){
	n,err := strconv.Atoi(ctx.Args().First())
	if err != nil{
		fmt.Println(err)
	}

	//检测是否超过区块链的个数
	if n>len(bc.chain)-1 || n<0{
		fmt.Printf("there is not a block by number %d\n",n)
	}else {
		block := bc.chain[n]
		fmt.Printf("chain[%d]-->"+block.String()+"\n",n)
	}

}

//String化
func (bc *BlockChain) String(){
	fmt.Println("this is a blockchain")
}

func (bc *BlockChain) AddBlock(b *Block){
	bc.chain = append(bc.chain, b)
}

func (bc *BlockChain) AddTransaction(t *Transaction){
	bc.txpool = append(bc.txpool,t)
}




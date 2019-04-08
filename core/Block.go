package core

import "fmt"

type Block struct {
	id int
}


//将区块的数据String化
func (b *Block) String() string {
	return fmt.Sprintf("this is a block")
}

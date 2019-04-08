package core

import "flag"

type BtcdClient struct {


}

func NewBtcdClient() *BtcdClient {
	bc := &BtcdClient{}
	return bc
}

func init(){

}

func (*BtcdClient) run(){
	flag.Parse()

}
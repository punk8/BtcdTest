package main

import (
	"flag"
	"fmt"
	"strings"
)

/*func main(){
	btcd := core.NewBtcdClient()
	fmt.Printf("Now run a BtcdClient..")
	btcd.run()

}*/

var name = flag.String("name","张三","学生姓名")

var major string

func init(){
	const (
		defaultMajor = "computer"
		usage = "their major"
	)
	flag.StringVar(&major,"major",defaultMajor,usage)
	flag.StringVar(&major,"m",defaultMajor,usage+"(简写)")
}

type Classmates []string


func (i *Classmates) String() string{
	return fmt.Sprint(*i)
}

func (i *Classmates) Set(value string) error{
	for _,dt := range strings.Split(value,","){
		*i = append(*i,dt)
	}
	return nil
}

var mates Classmates

func init(){
	flag.Var(&mates,"class","逗号分割到同学列表")
}


func main(){
	flag.Parse()
	fmt.Println("name",*name)
	fmt.Println("major",major)
	fmt.Println("classmates",mates)
}
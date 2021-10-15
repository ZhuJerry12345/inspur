package main

//tset
import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type BlockMessage struct {
}

type Transaction struct {
	Hash []byte
	From int8
}

func (t *Transaction) IsEmpty() bool {
	return reflect.DeepEqual(t, Transaction{})
}

type Line struct {
	InstituteLine [4][]Transaction //机构交易队列
	instituteNum  [4]int           //队列长度排序
	total         int              //剩余机构数量
	lotalTransNum int              //总交易数量
	lastOperation [2]int           //上次出队位置
}

type Quintet struct {
	Trans   [3]Transaction
	Compute int8
	Storage []int8
}

func (l *Line) Push(trans Transaction) { //入队
	l.InstituteLine[trans.From] = append(l.InstituteLine[trans.From], trans)

	l.lotalTransNum++
}

func (l *Line) Init() { //排序
	l.total = 4
	temp := -1
	location := -1
	length := [4]int{len(l.InstituteLine[0]), len(l.InstituteLine[1]), len(l.InstituteLine[2]), len(l.InstituteLine[3])}
	for i := 0; i < 4; i++ {
		for i := 0; i < 4; i++ { //选出最大的
			if temp < length[i] {
				temp = length[i]
				location = i
			}
		}
		if temp == 0 {
			l.instituteNum[i] = -1 //该机构没有交易
			l.total--
			length[location] = -1 //确定位置后不参与排序
			temp = -1
			location = -1
		} else {
			l.instituteNum[i] = location
			length[location] = -1
			temp = -1
			location = -1
		}
	}
}

func (l *Line) Pull() (trans [3]Transaction) { //出队
	if l.total == 4 || l.total == 3 {
		trans[0] = l.InstituteLine[l.instituteNum[0]][0] //从最长的前三个里面取
		trans[1] = l.InstituteLine[l.instituteNum[1]][0]
		trans[2] = l.InstituteLine[l.instituteNum[2]][0]
		l.InstituteLine[l.instituteNum[0]] = l.InstituteLine[l.instituteNum[0]][1:len(l.InstituteLine[l.instituteNum[0]])] //取完删掉
		l.InstituteLine[l.instituteNum[1]] = l.InstituteLine[l.instituteNum[1]][1:len(l.InstituteLine[l.instituteNum[1]])]
		l.InstituteLine[l.instituteNum[2]] = l.InstituteLine[l.instituteNum[2]][1:len(l.InstituteLine[l.instituteNum[2]])]
	} else if l.total == 2 {
		trans[0] = l.InstituteLine[l.instituteNum[0]][0]
		trans[1] = l.InstituteLine[l.instituteNum[0]][1]
		trans[2] = l.InstituteLine[l.instituteNum[1]][0]
		l.InstituteLine[l.instituteNum[0]] = l.InstituteLine[l.instituteNum[0]][2:len(l.InstituteLine[l.instituteNum[0]])]
		l.InstituteLine[l.instituteNum[1]] = l.InstituteLine[l.instituteNum[0]][1:len(l.InstituteLine[l.instituteNum[1]])]
	} else if l.total == 1 {
		trans[0] = l.InstituteLine[l.instituteNum[0]][0]
		trans[1] = l.InstituteLine[l.instituteNum[0]][1]
		trans[2] = l.InstituteLine[l.instituteNum[0]][2]
		l.InstituteLine[l.instituteNum[0]] = l.InstituteLine[l.instituteNum[0]][3:len(l.InstituteLine[l.instituteNum[0]])]
	}
	l.Init()
	return trans
}

func (q *Quintet) Distribute() {
	if !q.Trans[1].IsEmpty() && !q.Trans[2].IsEmpty() {
		A := q.Trans[0].From
		B := q.Trans[1].From
		C := q.Trans[2].From
		if A != B && A != C {
			temp := int8(uint8(q.Trans[0].Hash[0]) % 3)
			if temp == 0 {
				q.Compute = A
			} else if temp == 1 {
				q.Compute = B
			} else {
				q.Compute = C
			}
			q.Storage = append(q.Storage, 6-A-B-C)
		}
		if A != B && A == C {
			q.Compute = q.Trans[0].From
			a := 6 - A - B
			b := 14 - A ^ 2 - B ^ 2
			q.Storage = append(q.Storage, (a+(2*b-a^2)^(1/2))/2)
			q.Storage = append(q.Storage, (a-(2*b-a^2)^(1/2))/2)
		}
		if A == B && A == C {
			q.Compute = q.Trans[0].From
			q.Storage = append(q.Storage, (A+1)%4)
			q.Storage = append(q.Storage, (A+2)%4)
			q.Storage = append(q.Storage, (A+3)%4)
		}
	} else {
		if q.Trans[1].IsEmpty() {
			q.Compute = q.Trans[0].From
			q.Storage = append(q.Storage, (q.Trans[0].From+int8(uint8(q.Trans[0].Hash[0])%3+1))%4)
			/*} else {                           直接拆成一个存副本算了 或者挪到下一个区块？
			if q.Trans[0].From == q.Trans[1].From{
				q.Compute = q.Trans[0].From
				q.Storage = append(q.Storage, int8(uint8(q.Trans[0].Hash[0]) % 3))
				q.Storage = append(q.Storage, int8(uint8(q.Trans[1].Hash[0]) % 3))
			}else{
				q.Compute = int8(uint8(q.Trans[0].Hash[0]) % 3)
			}*/
		}
	}

}

func GenerateBlocks() (transses []Transaction) { //随机生成一个区块
	var from int8
	var data []byte
	var hash []byte
	sha := sha256.New()
	var trans Transaction
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 62; i++ {
		//产生机构编号
		from = int8(rand.Intn(4))
		//产生随机哈希
		data = append(data, byte(rand.Uint64()))
		_, err := sha.Write(data)
		if err != nil {
			err.Error()
			return
		}
		hash = sha.Sum(nil)
		trans.From = from
		trans.Hash = hash
		transses = append(transses, trans)
		//fmt.Println(trans.From, trans.Hash)
	}
	return transses
}
func GetBlockMessage() {

}

func GenerateDateDistribute(transses []Transaction) {
	var line Line
	var quintetlinte []Quintet
	for _, trans := range transses { //交易入队
		line.Push(trans)
	}

	fmt.Println(len(line.InstituteLine[0]), len(line.InstituteLine[1]), len(line.InstituteLine[2]), len(line.InstituteLine[3]))

	line.Init()
	for i := 0; i < len(transses)/3; i++ { //交易出队
		var quintet Quintet
		quintet.Trans = line.Pull()
		quintet.Distribute() //确定冗余方式
		fmt.Println(quintet.Trans[0].From, quintet.Trans[1].From, quintet.Trans[2].From, quintet.Compute, quintet.Storage)
		quintetlinte = append(quintetlinte, quintet)
	}
	if len(transses)%3 != 0 {
		if line.total == 1 {
			for i := 0; i < len(transses)%3; i++ { //交易出队
				var quintet Quintet
				quintet.Trans[0] = line.InstituteLine[line.instituteNum[0]][i]
				quintet.Distribute() //确定冗余方式
				fmt.Println(quintet.Trans[0].From, quintet.Trans[1].From, quintet.Trans[2].From, quintet.Compute, quintet.Storage)
				quintetlinte = append(quintetlinte, quintet)
			}
		} else {
			for i := 0; i < len(transses)%3; i++ { //交易出队
				var quintet Quintet
				quintet.Trans[0] = line.InstituteLine[line.instituteNum[i]][0]
				quintet.Distribute() //确定冗余方式
				fmt.Println(quintet.Trans[0].From, quintet.Trans[1].From, quintet.Trans[2].From, quintet.Compute, quintet.Storage)
				quintetlinte = append(quintetlinte, quintet)
			}
		}
	}
}
func GetBlockData() {

}
func Calculate() {

}
func main() {
	//var AData [4096]byte
	//var BData [4096]byte
	//var CData [4096]byte
	//var DData [4096]byte
	block := GenerateBlocks()
	for _, trans := range block {
		fmt.Println(trans.From, trans.Hash)
	}

	GetBlockMessage()
	GenerateDateDistribute(block)
	GetBlockData()
	Calculate()
}

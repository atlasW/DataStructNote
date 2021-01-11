package search

import (
	"time"
)

/*
顺序搜索   遍历
*/
const N = 17186792

type Location struct {
	ID string    `json:"_id"`
	V  int       `json:"v"`
	P  int       `json:"p"`
	I  string    `json:"i"`
	O  float64   `json:"o"`
	S  int       `json:"s"`
	L  []float64 `json:"l"`
	T  struct {
		Date time.Time `json:"$date"`
	} `json:"t"`
	D int `json:"d"`
	B int `json:"b"`
	C struct {
		Date time.Time `json:"$date"`
	} `json:"c"`
}

//
//func main() {
//	starttime := time.Now()
//	path := "../example/1.json"
//	jsonfile, _ := os.Open(path)
//	i := 0 //行数
//	br := bufio.NewReader(jsonfile)
//	StructArry := make([]Location, N, N)
//	var line []byte
//	var err error
//	var value Location
//	for {
//		line, _, err = br.ReadLine()
//		if err == io.EOF {
//			break
//		}
//		json.Unmarshal(line, &value)
//		StructArry[i] = value
//		i++
//	}
//	fmt.Println("本次查询用了：", time.Since(starttime))
//	//
//	for {
//		var a int
//		var b []int
//		fmt.Printf("请输入需要查询的i:")
//		fmt.Scanf("%d", &a)
//		starttime := time.Now()
//		for i := 0; i < N; i++ {
//			if StructArry[i].S == a {
//				//fmt.Println(StructArry[i])
//				b = append(b, i)
//			}
//		}
//		fmt.Println("本次查询用了：", time.Since(starttime))
//		fmt.Println(len(b))
//		b = []int{}
//	}
//}

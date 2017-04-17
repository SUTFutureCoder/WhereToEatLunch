package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	rr = rand.New(rand.NewSource(time.Now().UnixNano()))
)

func main(){
	r   := make(map[int]string)
	r[0] = "桂林米粉"
	r[1] = "宽小面"
	r[2] = "多味轩"
	r[3] = "米FEEL"
	r[4] = "排骨米饭"
	r[5] = "V道食街"
	r[6] = "I米粟"
	r[7] = "后厂村九号"
	r[8] = "再来一次～"
	r[9] = "醉唐轩"
	r[10] = "重庆小面"
	r[11] = "排骨米饭"
	r[12] = "麦当劳"
	r[13] = "他二哥"

	
	randRet := rr.Perm(len(r) - 1)[0]
	fmt.Println(r[randRet])

}

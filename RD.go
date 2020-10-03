package main

import (
	"fmt"
	"bufio"
	"os"
)

func main(){
	testStr := make([]string, 0)

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		testStr = append(testStr, sc.Text())
	}
	if err := sc.Err(); err != nil {
		fmt.Fprintf(os.Stderr, " 从输入流读取参数失败: %s\n ", err)
	}
	for _,afterStr:=range removeDuplicate(testStr){
		fmt.Println(afterStr)
	}
}

func removeDuplicate(arr []string) []string {
    resArr := make([]string, 0)
    tmpMap := make(map[string]interface{})
    for _, val := range arr {
        if _, ok := tmpMap[val]; !ok {
            resArr = append(resArr, val)
            tmpMap[val] = nil
        }
	}
	return resArr
}

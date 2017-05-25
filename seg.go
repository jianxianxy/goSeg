package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var Dict map[string][]string //字典

func main() {
	plan := SegString("苹果胡萝卜泥")

	matl := MatchLevel(plan, "苹果胡萝卜泥苹果")
	fmt.Println(matl)
}

func SegString(str string) []string {
	var plan []string
	//加载字典
	LoadDict()
	//正序分词
	planLr := GetWordLr(str, make([]string, 0, 5))
	//倒序分词
	planRl := GetWordRl(str, make([]string, 0, 5))
	//反转
	SliceReverse(&planRl)
	//比较
	if SliceIsEqual(planLr, planRl) {
		plan = planLr
	} else {
		plan = PlanFilter(planLr, planRl)
	}
	return plan
}

//加载字典
func LoadDict() {
	Dict = make(map[string][]string)
	file, err := os.Open("./data/dict.txt")
	defer file.Close()
	if err != nil {
		panic("加载字典失败")
	}
	reader := bufio.NewReader(file)
	for {
		buf, _, err := reader.ReadLine()
		if err != io.EOF {
			split := strings.Split(string(buf), " ")
			if len(split) > 1 {
				Dict[string(split[0])] = split[1:]
			}
		} else {
			break
		}
	}
}

//正序分词
func GetWordLr(str string, pla []string) []string {
	sta := 0
	add := 2
	strArr := []rune(str)
	strlen := len(strArr)
	var currword string
	if strlen < 2 {
		pla = append(pla, str)
		return pla
	}
	for {
		if sta+add <= strlen {
			word := strArr[sta:add]
			if _, ok := Dict[string(word)]; ok {
				currword = string(word)
				add += 1
			} else {
				add -= 1
				if add == 1 {
					word := strArr[sta:add]
					currword = string(word)
				}
				pla = append(pla, currword)
				pla = GetWordLr(string(strArr[sta+add:]), pla)
				break
			}
		} else {
			pla = append(pla, currword)
			break
		}
	}
	return pla
}

//倒序分词
func GetWordRl(str string, pla []string) []string {
	add := 2
	strArr := []rune(str)
	strlen := len(strArr)
	var currword string
	if strlen < 2 {
		pla = append(pla, str)
		return pla
	}
	for {
		if strlen-add >= 0 {
			word := strArr[strlen-add : strlen]
			if _, ok := Dict[string(word)]; ok {
				currword = string(word)
				add += 1
			} else {
				add -= 1
				if add == 1 {
					word := strArr[strlen-add : strlen]
					currword = string(word)
				}
				pla = append(pla, currword)
				pla = GetWordRl(string(strArr[:strlen-add]), pla)
				break
			}
		} else {
			pla = append(pla, currword)
			break
		}
	}
	return pla
}

//数组切片反转顺序
func SliceReverse(sli *[]string) {
	len := len(*sli)
	for i := 1; i <= len/2; i++ {
		(*sli)[i-1], (*sli)[len-i] = (*sli)[len-i], (*sli)[i-1]
	}
}

//切片内容是否相同
func SliceIsEqual(sliL, sliR []string) bool {
	if len(sliL) != len(sliR) {
		return false
	}
	for key, val := range sliL {
		if val != sliR[key] {
			return false
		}
	}
	return true
}

//方案按照权重筛选
func PlanFilter(sliL, sliR []string) []string {
	if len(sliL) < len(sliR) {
		return sliL
	}
	if len(sliR) < len(sliL) {
		return sliR
	}
	var levl, levr int
	for _, val := range sliL {
		rate, err := strconv.Atoi(Dict[val][0])
		if err == nil {
			levl += rate
		}
	}
	for _, val := range sliR {
		rate, err := strconv.Atoi(Dict[val][0])
		if err == nil {
			levr += rate
		}
	}
	if levl > levr {
		return sliL
	}
	return sliR
}

//匹配度分析
func MatchLevel(plan []string, con string) map[string]int {
	matl := make(map[string]int, 0)
	for _, val := range plan {
		num := strings.Count(con, val)
		matl[val] = num
	}
	return matl
}

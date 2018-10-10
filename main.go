package main

import (
	"flag"
	"fmt"
	"strconv"
	"math"
	"strings"
)

func main(){
	rows := flag.Int("rows", 0, "Rows count.")
	columns := flag.Int("columns", 0, "Columns count.")
	peoples := flag.Int("peoples", 0, "Peoples count.")

	flag.Parse()

	num, err := findUnhappyPoints(*rows,*columns,*peoples)
	fmt.Println("Count unhappy points: ",num)
	if err != ""{fmt.Println(err)}

}

func findUnhappyPoints(rows, columns, peoples int) (num int, err string){
	permutation := ""
	switch {
	case rows <= 0:
		return 0, "Rows not valid"
	case columns <= 0:
		return 0, "Columns not valid"
	case peoples > rows*columns:
		return 0, "Peoples count to match"
	default:
		permutation,num = findBestPermutation(rows,columns,peoples)
		fmt.Println(permutation)
		err = ""
	}
	return
}

func findBestPermutation(rows,columns,peoples int)(permutation string, count int){
	if peoples ==0 {
		return "", 0
	}else{
		permutations:=findAllPermitations(rows,columns,peoples)
		count =rows*columns*2
		for _, p:=range permutations{
			c := countUnheppiPoints(rows,columns,p)
			if c < count{
				count = c
				permutation = p
			}
		}
		return
	}
}

func countUnheppiPoints(rows,columns int, permutation string) (count int){

	p:=make([]string, rows,rows)
	for i:=0; i<=rows-1;i++{
		p[i] = permutation[(i*columns):(i*columns+(columns))]
	}
	for i, item:=range p{
		//fmt.Println(i)
		if i == rows-1 {
			for j := 0; j <= len(item)-2; j++ {
				if string(item[j]) == "1" && string(item[j+1]) == "1" {
					count++
				}
			}
		}else{
			for j:=0; j<=len(item)-2; j++{
				if string(item[j]) == "1" && string(item[j+1]) == "1" {
					count++
				}
				if string(item[j]) == "1" && string(p[i+1][j])=="1"{
					count++
				}
			}
			if string(p[i][columns-1]) == "1" && string(p[i+1][columns-1]) == "1" {
				count++
			}
		}
	}
	return
}

func findAllPermitations(rows, columns, peoples int) (permutations []string){
	roomsCount := rows*columns
	for i := 0; i <= int(math.Pow(2, float64(roomsCount))); i++{
		variant:= strconv.FormatInt(int64(i),2)
		if len(variant)< rows*columns{
			variant = strings.Repeat("0", rows*columns-len(variant))
		}
		if res:=checkCount(variant,peoples); res{
			permutations = append(permutations, variant)
			//fmt.Println(permutations)
		}
	}
	return
}
func checkCount(item string, people int) (result bool){
	count :=0
	for _, elem:= range item{
		if fmt.Sprintf("%c",elem) == "1"{
			count +=1
		}
	}
	if count == people{
		result = true
	}
	return
}
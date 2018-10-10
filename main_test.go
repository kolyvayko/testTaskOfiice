package main

import (
	"testing"
	"fmt"
)

type testdata struct {
	values []int
	result int
}

var tests = []testdata{
	{  []int{5,2,8}, 7 }, {  []int{3,5,14}, 18 }, {  []int{1,16,1}, 0 }, {  []int{3,5,1}, 0 },
	{  []int{8,2,12}, 10 }, {  []int{16,1,1}, 0 }, {  []int{3,3,6}, 3 }, {  []int{2,6,12}, 16 },
	{  []int{15,1,0}, 0 }, {  []int{5,3,7}, 0 }, {  []int{4,3,5}, 0 }, {  []int{3,5,11}, 8 },
	{  []int{7,2,13}, 16 }, {  []int{15,1,6}, 0 }, {  []int{15,1,15}, 14 }, {  []int{4,4,9}, 2 },
	{  []int{5,3,8}, 0}, {  []int{3,5,6}, 0}, {  []int{16,1,7}, 0}, {  []int{1,15,7}, 0},
	{  []int{4,3,12}, 17}, {  []int{5,3,13}, 14}, {  []int{2,4,5}, 2}, {  []int{5,3,5}, 0},
	{  []int{16,1,16}, 15}, {  []int{2,5,8}, 7}, {  []int{5,3,4}, 0}, {  []int{5,3,10}, 6},
	{  []int{4,4,7}, 0}, {  []int{3,5,9}, 3}, {  []int{4,2,2}, 0}, {  []int{4,4,15}, 20},
	{  []int{2,2,4}, 4}, {  []int{5,3,11}, 8}, {  []int{4,4,8}, 0}, {  []int{1,16,9}, 1},
	{  []int{4,4,16}, 24}, {  []int{1,15,6}, 0}, {  []int{15,1,8}, 0}, {  []int{5,3,6}, 0},
	{  []int{16,1,9}, 1}, {  []int{3,5,15}, 22}, {  []int{1,15,1}, 0}, {  []int{1,15,0}, 0 },
	{  []int{2,5,9}, 10}, {  []int{3,5,10}, 6}, {  []int{1,15,15}, 14}, {  []int{3,2,0}, 0},
	{  []int{5,3,2}, 0}, {  []int{5,3,1}, 0}, {  []int{5,2,4}, 0}, {  []int{3,5,4}, 0},
	{  []int{2,7,13}, 16}, {  []int{3,3,0}, 0}, {  []int{7,2,11}, 10}, {  []int{4,4,0}, 0},
	{  []int{1,1,0}, 0}, {  []int{2,6,9}, 7}, {  []int{3,5,3}, 0}, {  []int{5,3,15}, 22},
	{  []int{5,2,6}, 2}, {  []int{3,4,12}, 17}, {  []int{2,3,6}, 7}, {  []int{1,1,1}, 0},
	{  []int{15,1,1}, 0}, {  []int{1,16,16}, 15}, {  []int{2,2,2}, 0}, {  []int{3,3,9}, 12},
	{  []int{16,1,8}, 0}, {  []int{9,1,6}, 2}, {  []int{5,3,12}, 11}, {  []int{2,2,3}, 2},
	{  []int{3,5,7}, 0}, {  []int{7,2,0}, 0}, {  []int{4,3,6}, 0}, {  []int{2,3,4}, 2},
	{  []int{1,15,8}, 0}, {  []int{16,1,0}, 0}, {  []int{5,3,9}, 3}, {  []int{15,1,7}, 0},
	{  []int{2,4,6}, 4}, {  []int{1,16,7}, 0}, {  []int{3,5,12}, 11}, {  []int{1,16,8}, 0},
	{  []int{4,4,1}, 0}, {  []int{3,5,0}, 0}, {  []int{3,5,8}, 0}, {  []int{1,16,0}, 0},
	{  []int{5,3,3}, 0}, {  []int{5,3,0}, 0}, {  []int{1,13,9}, 4}, {  []int{3,5,2}, 0},
	{  []int{1,9,6}, 2}, {  []int{6,2,12}, 16}, {  []int{4,3,8}, 4}, {  []int{3,5,5}, 0},
	{  []int{5,3,14}, 18}, {  []int{4,3,7}, 2}, {  []int{6,2,4}, 0}, {  []int{3,5,1}, 0},
	{  []int{0,3,14},0}, {  []int{1,0,14},0},{  []int{3,3,14},0},{  []int{30,30,700},1800},

}
func TestFindUnhappyPoints(t *testing.T) {
	for _, item:= range tests {
		fmt.Println(item)
		count,err := findUnhappyPoints(item.values[0],item.values[1],item.values[2])
		if count != item.result {
			t.Error("Expected", item.result, " Got: ", count)
		}else{
			fmt.Println(err)
		}
	}
}

func TestCounUnheppiPoints(t *testing.T) {
	permutation := countUnheppiPoints(3,3,"111100011")
	//fmt.Println(permutation)
	if permutation != 4 {
		t.Error(permutation)
	}
	permutation = countUnheppiPoints(3,3,"111111111")
	//fmt.Println(permutation)
	if permutation != 12 {
		t.Error(permutation)
	}
	permutation = countUnheppiPoints(5,5,"1111100000111111111100001")
	//fmt.Println(permutation)
	if permutation != 18 {
		t.Error(permutation)
	}
}

func TestFindAllPermutations(t *testing.T) {
	findAllPermitations(3,3,5)

}
func TestCheckCount(t *testing.T){
	type values struct{
		Value string
		Count int
		Res bool
	}
	var testdata = []values{
		{"011001110", 5, true},
		{"011001110", 3, false},
		{"011101110", 5, false},
		{"011101110", 6, true},
		{"0", 6, false},
		{"0", 1, false},
		{"1", 1, true},

	}
	for _, value:=range testdata{
		res:= checkCount(value.Value, value.Count)
		if res != value.Res {
			t.Error(res)
		}
	}

}

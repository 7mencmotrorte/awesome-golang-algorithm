package Solution

import (
	"reflect"
	"strconv"
	"testing"
)

func TestSolution(t *testing.T) {
	//	测试用例
	cases := []struct {
		name   string
		inputs [][]int
		expect *QuadNode
	}{
		{"TestCase1", [][]int{
			{0, 1}, {1, 0},
		}, &QuadNode{
			IsLeaf:      false,
			Val:         true,
			TopLeft:     &QuadNode{IsLeaf: true, Val: false},
			TopRight:    &QuadNode{IsLeaf: true, Val: true},
			BottomLeft:  &QuadNode{IsLeaf: true, Val: true},
			BottomRight: &QuadNode{IsLeaf: true, Val: false},
		}},
		{"TestCase2", [][]int{
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1},
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0, 0},
			{1, 1, 1, 1, 0, 0, 0, 0},
		}, &QuadNode{
			IsLeaf:      false,
			Val:         true,
			TopLeft:     &QuadNode{IsLeaf: true, Val: true},
			BottomLeft:  &QuadNode{IsLeaf: true, Val: true},
			BottomRight: &QuadNode{IsLeaf: true, Val: false},
			TopRight: &QuadNode{
				IsLeaf:      false,
				Val:         true,
				TopLeft:     &QuadNode{IsLeaf: true, Val: false},
				TopRight:    &QuadNode{IsLeaf: true, Val: false},
				BottomLeft:  &QuadNode{IsLeaf: true, Val: true},
				BottomRight: &QuadNode{IsLeaf: true, Val: true},
			},
		}},
	}

	//	开始测试
	for i, c := range cases {
		t.Run(c.name+" "+strconv.Itoa(i), func(t *testing.T) {
			got := Solution(c.inputs)
			if !reflect.DeepEqual(got, c.expect) {
				t.Fatalf("expected: %v, but got: %v, with inputs: %v",
					c.expect, got, c.inputs)
			}
		})
	}
}

// 压力测试
func BenchmarkSolution(b *testing.B) {
}

// 使用案列
func ExampleSolution() {
}

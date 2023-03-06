package hashing

import (
	"crypto/md5"
	"fmt"
	"math"
	"strconv"
	"testing"
)

func Test_string_to_float(t *testing.T) {
	for i := 0; i < 10; i++ {
		f := hashToUnitInterval(md5.New, []byte(strconv.Itoa(i)))
		t.Log(i, f)
	}
}

func TestNode_compute_weighted_score(t *testing.T) {
	node1 := NewNode("node 1", 1)

	for i := 0; i < 10; i++ {
		f1 := node1.computeWeightedScore(md5.New, strconv.Itoa(i))
		t.Log(i, f1)
	}

}

func TestNode(t *testing.T) {

	node1 := NewNode("A", 10)
	node2 := NewNode("B", 20)
	node3 := NewNode("C", 30)
	node4 := NewNode("D", 40)
	node5 := NewNode("E", 50)

	board := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
	}

	for i := 0; i < 10000; i++ {
		f1 := node1.computeWeightedScore(md5.New, strconv.Itoa(i))
		f2 := node2.computeWeightedScore(md5.New, strconv.Itoa(i))
		f3 := node3.computeWeightedScore(md5.New, strconv.Itoa(i))
		f4 := node4.computeWeightedScore(md5.New, strconv.Itoa(i))
		f5 := node5.computeWeightedScore(md5.New, strconv.Itoa(i))

		// println("f1", f1)
		// println("f2", f2)
		// println("f3", f3)
		// println("f4", f4)
		// println("f5", f5)

		board[IndexMax([]float64{f1, f2, f3, f4, f5})]++
	}

	fmt.Println(board)

	div := func(a, b int) int { return a / b }

	fmt.Println(div(board[0], board[0]))
	fmt.Println(div(board[1], board[0]))
	fmt.Println(div(board[2], board[0]))
	fmt.Println(div(board[3], board[0]))
	fmt.Println(div(board[4], board[0]))
}

func TestHRW(t *testing.T) {

	node1 := NewNode("A", 0.1)
	node2 := NewNode("B", 0.2)
	node3 := NewNode("C", 0.3)
	node4 := NewNode("D", 0.4)
	node5 := NewNode("E", 0.5)

	board := map[string]int{
		node1.Name: 0,
		node2.Name: 0,
		node3.Name: 0,
		node4.Name: 0,
		node5.Name: 0,
	}

	nodes := []*Node{node1, node2, node3, node4, node5}

	for i := 0; i < 10000; i++ {
		node := DetermineResponsibleNode(md5.New, strconv.Itoa(i), nodes)

		board[node.Name]++
	}

	fmt.Println(board)

	div := func(a, b int) int { return a / b }

	fmt.Println(div(board[node1.Name], board[node1.Name]))
	fmt.Println(div(board[node2.Name], board[node1.Name]))
	fmt.Println(div(board[node3.Name], board[node1.Name]))
	fmt.Println(div(board[node4.Name], board[node1.Name]))
	fmt.Println(div(board[node5.Name], board[node1.Name]))
}

func TestMathLog(t *testing.T) {

	weights := []float64{
		0.01, 0.1, 1, 10, 100,
	}

	for i := 10; i < 20; i++ {
		for _, w := range weights {
			log_score := 1.0 / math.Log(float64(i))
			score := w * log_score
			t.Log(i, score)
		}
	}
}

func TestReplicaSets(t *testing.T) {

	max_nodes := 10
	rate := 3

	nodes := make([]int, max_nodes)
	for i := 0; i < len(nodes); i++ {
		nodes[i] = i
	}

	var set = make([][]int, len(nodes))
	for i := 0; i < len(nodes); i++ {
		set[i] = make([]int, rate)
		for j := 0; j < rate; j++ {
			set[i][j] = nodes[(i+j)%len(nodes)]
		}
	}

	for _, set_item := range set {
		// for _, sub_item := range set {
		fmt.Printf("%+v\n", set_item)
		// }
	}

}

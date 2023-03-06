package hashing

import (
	"encoding/binary"
	"hash"
	"math"

	"golang.org/x/exp/constraints"
)

type Hasher func() hash.Hash

// hashToUnitInterval
//
//	문자열 값을 입력 받아 부동 소수 값으로 변환
func hashToUnitInterval(hasher Hasher, b []byte) float64 {
	hash := hasher()
	hash.Write(b)
	sum := hash.Sum([]byte{})
	size := hash.BlockSize()
	uint64_ := binary.BigEndian.Uint64(sum)
	float64_ := float64(uint64_)

	return (float64_ + 1) / math.Exp2(float64(size))
}

type Node struct {
	Name      string
	Weight    float64
	Convertor func()
}

func NewNode(name string, weight float64) *Node {
	return &Node{Name: name, Weight: weight}
}

func (node Node) computeWeightedScore(hasher Hasher, key string) float64 {
	score := hashToUnitInterval(hasher, []byte(node.Name+key))
	log_score := 1.0 / -math.Log(score)

	return node.Weight * log_score
}

func DetermineResponsibleNode(hasher Hasher, key string, nodes []*Node) *Node {
	var values = make([]float64, len(nodes))
	for i := range nodes {
		values[i] = nodes[i].computeWeightedScore(hasher, key)
	}

	idx := IndexMax(values)
	if idx == -1 {
		return nil
	}

	return nodes[idx]
}

type comparable = constraints.Ordered

func IndexMax[T comparable](values []T) int {
	// 비교할 값이 없음
	if len(values) == 0 {
		return -1
	}

	var cmp = values[0]
	var idx int = 0

	// 비교할 값이 한개
	if len(values) == 1 {
		return idx
	}

	for i := 1; i < len(values); i++ {
		// compare
		if Less(cmp, values[i]) {
			cmp = values[i] // reset acc
			idx = i         // reset idx
		}
	}

	return idx
}

func Less[T comparable](a, b T) bool {
	return a < b
}

func Max[T comparable](values ...T) (T, bool) {
	var cmp T
	if len(values) == 0 {
		return cmp, false
	}

	cmp = values[0]
	if len(values) == 1 {
		return cmp, false
	}

	for i := 1; i < len(values); i++ {
		if Less(cmp, values[i]) {
			cmp = values[i]
		}
	}

	return cmp, true
}

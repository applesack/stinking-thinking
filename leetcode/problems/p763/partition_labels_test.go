package p763

import (
	"fmt"
	"testing"
)

func Test763_1(t *testing.T) {
	fmt.Println(partitionLabels("ababcbacadefegdehijhklij"))
	fmt.Println(partitionLabels("eccbbbbdec"))
	fmt.Println(partitionLabelsV2("eccbbbbdec"))
}

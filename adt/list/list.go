package list

import "fmt"

func checkIndex(index, size int) error {
	if index < 0 || index >= size {
		return fmt.Errorf("index out of bound [%d:%d]", size, index)
	}
	return nil
}

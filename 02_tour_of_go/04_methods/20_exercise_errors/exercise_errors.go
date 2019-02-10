package main

import "fmt"

// ErrNegativeSqrt 独自エラークラス
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

// Sqrt 平方根を返す
func Sqrt(root float64) (float64, error) {
	if root < 0 {
		return root, ErrNegativeSqrt(root)
	}
	initial := float64(1)
	x := babylonianMethod(initial, root)
	lastX := float64(0)

	for x != lastX {
		lastX = x
		x = babylonianMethod(x, root)
	}
	return x, nil
}

// babylonianMethod 平方根近似値計算(バビロニア人の方法)
func babylonianMethod(x float64, y float64) float64 {
	return (x + y/x) / 2
}

func main() {
	show(2)
	show(-2)
}

func show(root float64) {
	if sqrt, err := Sqrt(root); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sqrt)
	}
}

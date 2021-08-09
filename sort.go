package main

import (
	"fmt"
	"net/http"
	"sort"
	"sync"
)

func main() {
	http.ListenAndServe("https://shalove.github.io/emma/:9999", http.FileServer(http.Dir("/Users/shazinian/go/src/test")))

	//基本排序 int、float64、string,判断是否排序
	// basicSort()

	//希尔排序
	// arr := []int{1, 44, 22, 33, 66}
	// shellSort(arr)

	//堆排序
	// heapSort(arr)
}

func printi(i int, wg *sync.WaitGroup) error {
	defer wg.Done()
	fmt.Println(i)
	return nil
}

func basicSort() {
	a := []int{1, 2, 22, 3453, 11}
	a_pre := sort.IntsAreSorted(a)
	sort.Ints(a)
	a_after := sort.IntsAreSorted(a)

	b := []float64{1.0, 2.2, 22.3, 3453.1, 1.1}
	b_pre := sort.Float64sAreSorted(b)
	sort.Float64s(b)
	b_after := sort.Float64sAreSorted(b)

	c := []string{"be", "aa", "dd", "cca"}
	c_pre := sort.StringsAreSorted(c)
	sort.Strings(c)
	c_after := sort.StringsAreSorted(c)

	fmt.Println(a, a_pre, a_after)
	fmt.Println(b, b_pre, b_after)
	fmt.Println(c, c_pre, c_after)
}

func shellSort(arr []int) {
	le := len(arr)

	for gap := le / 2; gap > 0; gap /= 2 {
		for i := gap; i < le; i++ {
			//轮流对各组元素排序
			insertSort(arr, gap, i)
		}
	}
	fmt.Println("希尔排序结果")
	fmt.Println(arr)
}

func insertSort(arr []int, gap int, i int) {
	//插入排序
	for j := i - gap; j >= 0 && arr[i] < arr[j]; j -= gap {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

func heapSort(arr []int) {
	le := len(arr)
	for i := (le - 2) / 2; i >= 0; i-- {
		arr = downAdjust(arr, i, le)
	}
	fmt.Println("堆排序结果")
	fmt.Println(arr)
	for i := le - 1; i >= 1; i-- {
		//把堆顶的元素与最后一个元素交换
		arr[i], arr[0] = arr[0], arr[i]
		arr = downAdjust(arr, 0, i)
	}
	fmt.Println("堆排序结果")
	fmt.Println(arr)
}

func downAdjust(arr []int, parent int, n int) []int {
	tmp := arr[parent]

	child := 2*parent + 1

	for child < n {
		//右孩子更小的话，指向右孩子
		if child+1 < n && arr[child] > arr[child+1] {
			child++
		}
		if tmp <= arr[child] {
			break
		}
		//父节点下沉
		arr[parent] = arr[child]
		parent = child
		child = 2*parent + 1
	}
	arr[parent] = tmp
	return arr
}

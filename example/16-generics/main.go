package main

import "fmt"

// 作为泛型函数的示例，MapKeys 接受任意类型的Map并返回其Key的切片。 这个函数有2个类型参数 - K 和 V； K 是 comparable 类型，也就是说我们可以通过 == 和 != 操作符对这个类型的值进行比较。这是Go中Map的Key所必须的。 V 是 any 类型，意味着它不受任何限制 (any 是 interface{} 的别名类型).
func MapKeys[K comparable, V any](m map[K]V) []K {
	r := make([]K, 0, len(m))
	for k := range m {
		r = append(r, k)
	}
	return r
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (lst *List[T]) Push(v T) {
	if lst.tail == nil {
		lst.head = &element[T]{val: v}

		lst.tail = lst.head
	} else {
		lst.tail.next = &element[T]{val: v}
		lst.tail = lst.head
	}
}

func (lst *List[T]) GetAll() []T {
	var elems []T
	for e := lst.head; e != nil; e = e.next {
		elems = append(elems, e.val)
	}
	return elems
}

func main() {
	m := map[int]string{1: "1", 2: "2"}
	fmt.Println("keys m:", MapKeys(m))

	_ = MapKeys[int, string](m)

	lst := List[int]{}

	lst.Push(10)
	lst.Push(13)
	fmt.Println("list:", lst.GetAll())

}

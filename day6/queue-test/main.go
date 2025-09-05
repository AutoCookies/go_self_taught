package main

import "fmt"

type Queue struct {
	data []interface{}
}

type Stack struct {
	data []interface{}
}

func (s *Stack) Push(v interface{}) {
	s.data = append(s.data, v) // thêm vào cuối
}

func (s *Stack) Pop() interface{} {
	length := len(s.data)
	if length == 0 {
		return nil
	}
	top := s.data[length-1]    // lấy phần tử cuối
	s.data = s.data[:length-1] // bỏ phần tử cuối
	return top
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (q *Queue) Enqueue(v interface{}) {
	q.data = append(q.data, v) // thêm cuối
}

func (q *Queue) Dequeue() interface{} {
	if len(q.data) == 0 {
		return nil
	}
	front := q.data[0]
	q.data = q.data[1:]
	return front
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func (q *Queue) Size() int {
	return len(q.data)
}

// Queue: [1, 2, 3], I want to delete 2. Dequeu push them to stack except the 2 value
// Pop out the stack, enqueue them to queue

func (q *Queue) Delete(v interface{}) {
	temp := &Queue{}
	for !q.IsEmpty() {
		item := q.Dequeue()
		if item != v {
			temp.Enqueue(item)
		}
	}
	*q = *temp
}

// Queue: [1, 2, 3, 2, 3, 4]
// Temp Queue: []
// Iterated throught 1, and 1 still not in seen, add 1 to seen, add 1 to temp queue set seen to true.
// Continue with 2, 2 is not in seen, add 2 to seen, add 2 to temp queue set seen to true.
// Continue with 3, 3 is not in seen, add 3 to seen, add 3 to temp queue set seen to true.
// But in next iteration, 2 already in seen, so do not add 2 to temp queue

func (q *Queue) DeleteDuplicated() {
	seen := make(map[interface{}]bool)
	temp := &Queue{}

	for !q.IsEmpty() {
		item := q.Dequeue()
		if !seen[item] { // nếu chưa thấy item
			seen[item] = true  // đánh dấu đã thấy
			temp.Enqueue(item) // giữ lại trong queue tạm
		}
	}

	*q = *temp // copy lại dữ liệu
}

func (q *Queue) Print() {
	for i := 0; i < q.Size(); i++ {
		println(q.data[i].(int))
	}
}

// Queue: [1, 2, 3]
// InsertAt(1, 1)
// Go to the pos 1 and put the value to 1 pos.
// But in pos 2 now is 2, so It will append one more
// from q.data[pos] It will append one more time, with value append([]interface{}{v}, q.data[pos:]...) but
// append([]interface{}{v}, q.data[pos:]...) means add the rest of q.data begin from pos to the end. Then the this value will be
// append after q.data[:pos]

func (q *Queue) InsertAt(pos int, v interface{}) {
	if pos < 0 || pos > len(q.data) {
		return // vị trí không hợp lệ
	}
	q.data = append(q.data[:pos], append([]interface{}{v}, q.data[pos:]...)...)
}

func (q *Queue) DeleteAt(pos int) {
	if pos < 0 || pos > len(q.data) {
		return // vị trí không hợp lệ
	}
	q.data = append(q.data[:pos], q.data[pos+1:]...)
}

func main() {
	q := &Queue{}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	q.Delete(2)
	q.Print()
	fmt.Println()
	q.InsertAt(1, 2)

	q.Print()
	fmt.Println()
	q.DeleteAt(1)
	q.Print()
}

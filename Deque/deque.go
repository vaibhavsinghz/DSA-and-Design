package deque

type deque []int

func (d *deque) len() int {
	return len(*d)
}

func (d *deque) pushFront(x int) {
	*d = append([]int{x}, (*d)...)
}

func (d *deque) pushBack(x int) {
	*d = append(*d, x)
}

func (d *deque) popFront() {
	if len(*d) <= 0 {
		return
	}
	*d = (*d)[1:]
}

func (d *deque) popBack() {
	l := len(*d)
	if l <= 0 {
		return
	}
	*d = (*d)[:l-1]
}

func (d *deque) front() int {
	if len(*d) <= 0 {
		return 0
	}
	return (*d)[0]
}

func (d *deque) back() int {
	l := len(*d)
	if l <= 0 {
		return 0
	}
	return (*d)[l-1]
}

func main() {

}

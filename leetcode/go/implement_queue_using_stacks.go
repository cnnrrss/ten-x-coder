type MyQueue struct {
    Items []int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.Items = append(this.Items, x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    if len(this.Items) == 0 {
        return -1
    }
    
    item, items := this.Items[0], this.Items[1:]
    this.Items = items
    return item
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    if len(this.Items) == 0 {
        return -1
    }
    
    return this.Items[0]
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return len(this.Items) == 0   
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
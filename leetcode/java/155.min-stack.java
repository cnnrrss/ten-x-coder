import java.util.Stack;

/*
 * @lc app=leetcode id=155 lang=java
 *
 * [155] Min Stack
 */
class MinStack {
    Stack<Integer> s;
    Integer minEle;

    /** initialize your data structure here. */
    public MinStack() { s = new Stack<Integer>(); }
    
    public void push(int x) {
        if (s.isEmpty()) {
            minEle = x;
            s.push(x);
            return;
        }

        // If new number is less than original minEle
        if (x < minEle) {
            s.push(2 * x - minEle);
            minEle = x;
        } else {
            s.push(x);
        }
    }
    
    public void pop() {
        if (s.isEmpty()) {
            return;
        }
        Integer t = s.pop();
        if (t < minEle) {
            minEle = 2*minEle - t;
        }
    }
    
    public int top() {
        if (s.isEmpty()) { return -1; }
        else {
            Integer t = s.peek();
            if (t < minEle) {
                return minEle;
            } else {
                return t;
            }
        }
    }
    
    public int getMin() {
        if (s.isEmpty()) { return -1; }
        else { return minEle; }
    }
}

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack obj = new MinStack();
 * obj.push(x);
 * obj.pop();
 * int param_3 = obj.top();
 * int param_4 = obj.getMin();
 */


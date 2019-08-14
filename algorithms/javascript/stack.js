/**
 * initialize your data structure here.
 */
var MaxStack = function () {
    this.stack = [];
    this.maxStack = [];
};

/** 
 * @param {number} x
 * @return {void}
 */
MaxStack.prototype.push = function (x) {
    let max = this.maxStack.length === 0 ? x : this.maxStack[this.maxStack.length - 1];
    this.maxStack.push(max > x ? max : x);
    this.stack.push(x);
};

/**
 * @return {number}
 */
MaxStack.prototype.pop = function () {
    this.maxStack.pop();
    return this.stack.pop();
};

/**
 * @return {number}
 */
MaxStack.prototype.top = function () {
    return this.stack[this.stack.length - 1];
};

/**
 * @return {number}
 */
MaxStack.prototype.peekMax = function () {
    return this.maxStack[this.maxStack.length - 1];
};

/**
 * @return {number}
 */
MaxStack.prototype.popMax = function () {
    let max = this.peekMax();
    let temp = [];

    while (this.top() !== max) {
        temp.push(this.pop());
    }

    this.pop();

    while (temp.length !== 0) {
        this.push(temp.pop());
    }

    return max;
};

/**
 * Your MaxStack object will be instantiated and called as such:
 * var obj = new MaxStack()
 * obj.push(x)
 * var param_2 = obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.peekMax()
 * var param_5 = obj.popMax()
 */

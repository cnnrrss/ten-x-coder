var IntQueue = function() {
    this.items = [];
    this.size = 0;
};

/** 
 * @param {number} x
 * @return {void}
*/
IntQueue.prototype.enqueue = function(x) {
    this.items.push(x);
    this.size++
};

/**
 * @return {number}
 */
IntQueue.prototype.dequeue = function() {
    let item = this.items.shift();
    this.size--;
    return item;
};
/**
 * IntQueue object will be instantiated and called as such.
* var q = new IntQueue();
* q.enqueue(1);
* q.enqueue(2);
* q.enqueue(3);
* q.dequeue();
* q.dequeue();
 */
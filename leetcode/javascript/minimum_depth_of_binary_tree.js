/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number}
 */
var minDepth = function(root) {
    if (!root) return 0;
    let depth = 1;
    let queue = [root];
    console.log(root);
    
    if (!root.left && !root.right) return depth;

    while (queue.length > 0) {
        let nodeDepth = queue.length;
        
        for (let i = 0; i < nodeDepth; i++) {
            let node = queue.shift();
            if (!node.left && !node.right) return depth;
            else {
                console.log(node.right, node.left);
                if (node.left !== null) queue.push(node.left);
                if (node.right !== null) queue.push(node.right);
            }
        }
        depth++;
    }
    return depth;
};
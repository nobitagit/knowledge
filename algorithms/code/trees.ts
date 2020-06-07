/**
 * 1. How to invert a bunary tree
 */

/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */

class TreeNode {
  constructor(
    public val: number,
    public left: TreeNode,
    public right: TreeNode
  ) {}
}

/**
 * @param {TreeNode} root
 * @return {TreeNode}
 */
var invertTree = function (root: TreeNode) {
  // base case: if root is null
  // it means we already reached a leaf node
  // so return null and stop the iteration
  if (root == null) {
    return null;
  }

  // Start inverting the tree.
  // Store a temporary child node, so
  // we can mutate root
  let temp = root.left;
  // swap left
  root.left = invertTree(root.right);
  // swap right
  root.right = invertTree(temp);
  // return the newly swapped root
  return root;
};

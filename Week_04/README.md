# 学习笔记

## 二叉树的深度优先遍历

1. 前序深度优先遍历
```
/* pre-oder DFS traversal */
func preorder(n *node) {
    if n != nil {
        fmt.Printf(n.value + " ")
        preorder(n.left)
        preorder(n.right)
    }
}
```
2. 中序深度优先遍历
```
/* in-oder DFS traversal */
func inorder(n *node) {
    if n != nil {
        inorder(n.left)
        fmt.Printf(n.value + " ")
        inorder(n.right)
    }
}
```
3. 后序深度优先遍历
```
/* post-oder DFS traversal */
func postorder(n *node) {
    if n != nil {
        postorder(n.left)
        postorder(n.right)
        fmt.Printf(n.value + " ")
    }
}
```

* 广度优先遍历
```
/* breadth first traversal */
func breadth(n *node) {
    if n != nil {
        s := []*node{n}
        for len(s) > 0 {
            current_node := s[0]
            fmt.Printf(current_node.value + " ")
            s = s[1:]
            if current_node.left != nil {
                s = append(s, current_node.left)
            }
            if current_node.right != nil {
                s = append(s, current_node.right)
            }
        }
    }
}

```

## 二分法

* 二分法一般把能减少一半的搜索空间.是 O(log n) 的时间复杂度

* 二分法是工程中常用的方法.git 中的 `git bisect` 就是通过二分配合测试找 bad commit.

* 二分作用在有序的数据集上

* 二分需要注意求中点时防止越界

### 二分法的模板

*标准姿势*
```
int binarySearch(int[] nums, int target){
  if(nums == null || nums.length == 0)
    return -1;

  int left = 0, right = nums.length - 1;
  while(left <= right){
    // Prevent (left + right) overflow
    int mid = left + (right - left) / 2;
    if(nums[mid] == target){ return mid; }
    else if(nums[mid] < target) { left = mid + 1; }
    else { right = mid - 1; }
  }

  // End Condition: left > right
  return -1;
}
```

* 二分查找的最基础和最基本的形式。
* 查找条件可以在不与元素的两侧进行比较的情况下确定（或使用它周围的特定元素）。
* 不需要后处理，因为每一步中，你都在检查是否找到了元素。如果到达末尾，则知道未找到该元素。

*姿势二*

```
int binarySearch(int[] nums, int target){
  if(nums == null || nums.length == 0)
    return -1;

  int left = 0, right = nums.length;
  while(left < right){
    // Prevent (left + right) overflow
    int mid = left + (right - left) / 2;
    if(nums[mid] == target){ return mid; }
    else if(nums[mid] < target) { left = mid + 1; }
    else { right = mid; }
  }

  // Post-processing:
  // End Condition: left == right
  if(left != nums.length && nums[left] == target) return left;
  return -1;
}
```

* 一种实现二分查找的高级方法。
* 查找条件需要访问元素的直接右邻居。
* 使用元素的右邻居来确定是否满足条件，并决定是向左还是向右。
* 保证查找空间在每一步中至少有 2 个元素。
* 需要进行后处理。 当剩下 1 个元素时，循环 / 递归结束。 需要评估剩余元素是否符合条件。

*姿势三*

```
int binarySearch(int[] nums, int target) {
    if (nums == null || nums.length == 0)
        return -1;

    int left = 0, right = nums.length - 1;
    while (left + 1 < right){
        // Prevent (left + right) overflow
        int mid = left + (right - left) / 2;
        if (nums[mid] == target) {
            return mid;
        } else if (nums[mid] < target) {
            left = mid;
        } else {
            right = mid;
        }
    }

    // Post-processing:
    // End Condition: left + 1 == right
    if(nums[left] == target) return left;
    if(nums[right] == target) return right;
    return -1;
}

```
* 一种实现二分查找的高级方法。
* 查找条件需要访问元素的直接右邻居。
* 使用元素的右邻居来确定是否满足条件，并决定是向左还是向右。
* 保证查找空间在每一步中至少有 2 个元素。
* 需要进行后处理。 当你剩下 1 个元素时，循环 / 递归结束。 需要评估剩余元素是否符合条件。

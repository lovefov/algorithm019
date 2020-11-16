# 第二周学习笔记

## go 的哈希表

-golang 的 map 底层就是 hashmap

- go 的 map 通过拉链法缓解 hash 冲突的问题

- 通过 [hmap] (https://github.com/golang/go/blob/ed15e82413c7b16e21a493f5a647f68b46e965ee/src/runtime/map.go#L115-L129) 定义

```
type hmap struct {
	count     int
	flags     uint8
	B         uint8
	noverflow uint16
	hash0     uint32

	buckets    unsafe.Pointer
	oldbuckets unsafe.Pointer
	nevacuate  uintptr

	extra *mapextra
}

type mapextra struct {
	overflow    *[]*bmap
	oldoverflow *[]*bmap
	nextOverflow *bmap
}
```

- 定义:
  - count 表示当前哈希表中的元素数量；
  - B 表示当前哈希表持有的 buckets 数量，但是因为哈希表中桶的数量都 2 的倍数，所以该字段会存储对数，也就是 len(buckets) == 2^B；
  - hash0 是哈希的种子，它能为哈希函数的结果引入随机性，这个值在创建哈希表时确定，并在调用哈希函数时作为参数传入；
  - oldbuckets 是哈希在扩容时用于保存之前 buckets 的字段，它的大小是当前 buckets 的一半；

- map 的增删改查涉及到到 golang 运行时的一些东西,等搞明白了再添加

## 关于使用的 hashmap 的题,好像都可以通过一些简单的自定义的 hash 函数来做,比如记录每个元素出的次数然后进行相应的处理.

## 树,二叉树,二叉搜索树

- 树是非线性数据结构,通过递归定义.

- 因为树是递归定义的,所以通过递归的方式处理树的问题也是可以的.

## 二叉树的前中后序遍历.
```
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var inorder func(*TreeNode)
	inorder = func(n *TreeNode) {
		if n == nil {
			return
		}
		inorder(n.Left)
		res = append(res, n.Val)
		inorder(n.Right)
	}
	inorder(root)
	return res
}
```
- 前序: 根左右 中序: 左根右 后续: 左右根

- 只是取值的时机不同.

## 堆

- 堆是一个接口,底层有多种实现

- 一般分为大顶堆小顶堆.

- golang堆的接口:

```
func Fix(h Interface, i int)
func Init(h Interface)
func Pop(h Interface) interface{}
func Push(h Interface, x interface{})
func Remove(h Interface, i int) interface{}
type Interface
```

- golang 通过 heap 实现了优先队列


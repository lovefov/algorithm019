# 第三周学习笔记

## 递归

- 定义:（英语：Recursion），又译为递回，在数学与计算机科学中，是指在函数的定义中使用函数自身的方法。递归一词还较常用于描述以自相似方法重复事物的过程。例如，当两面镜子相互之间近似平行时，镜中嵌套的图像是以无限递归的形式出现的。也可以理解为自我复制的过程。

- 为啥要用递归: 递归描述的也是一种重复,自相似性的重复.计算机能解决重复性问题,那使用递归解决
那些自相似性的重复问题最为合适.

- 正视递归: 在合适的场景下使用递归,加上必要的优化,性能不一定差.

- 一般的递归可以转化成循环.可以通过维护一个栈,自已控制栈,实现递归.

- 模板:


```
def recursion(level, param1, param2, ...):
     # recursion terminator // 递归的终止条件,
     if level > MAX_LEVEL:
       process_result
       return
     # process logic in current level  //处理逻辑
     process(level, data...)
     # drill down
     self.recursion(level + 1, p1, ...) //进入下一层
     # reverse the current level status if needed //后处理当前层状态
```

## 分治

- 分治法是一种特殊的递归

- 分治模板比递归模板多了一个 将子问题解合并的动作.

- 循环递归
  - 在每一层递归上都有三个步骤：
  - 分解：将原问题分解为若干个规模较小，相对独立，与原问题形式相同的子问题。
  - 解决：若子问题规模较小且易于解决时，则直接解。否则，递归地解决各子问题。
  - 合并：将各子问题的解合并为原问题的解。

- 归并排序

```c++
 void merge_sort(int array[], unsigned int first, unsigned int last)
 {
    int mid = 0;
    if(first<last)
    {
        mid = (first+last)/2;
        merge_sort(array, first, mid);
        merge_sort(array, mid+1,last);
        merge(array,first,mid,last);
    }
 }
```

## 回溯

- 模板

```Python
result = []
def backtrack(路径, 选择列表):
    if 满足结束条件:
        result.add(路径)
        return

    for 选择 in 选择列表:
        做选择
        backtrack(路径, 选择列表)
        撤销选择
```

- 最经典的例题,八皇后问题

- 回溯法的特点:
  - 不像动态规划存在重叠子问题可以优化，回溯算法就是纯暴力穷举，复杂度一般都很高

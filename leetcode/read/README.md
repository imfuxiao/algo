## 数组与字符串

3. [无重复字符的最长子串](https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/)

难度: 中等
   
滑动窗口: 
   
1. 使用两个指针表示字符串中的某个子串
2. 判断是否有重复的字符


4. [寻找两个正序数组的中位数](https://leetcode-cn.com/problems/median-of-two-sorted-arrays/)

难度: 困难
需要使用二分查找, #TODO

15. [三数之和](https://leetcode-cn.com/problems/3sum/)

难度: 中等
排序+双指针

238. [除自身以外数组的乘积](https://leetcode-cn.com/problems/product-of-array-except-self/)

难度: 中等

因为不能使用除发, 所以不能直接使用总乘积除以当前数的方法.

需要分别使用两个for, 分别作为左侧数字的乘积和右侧数字的乘积.

415. [字符串相加](https://leetcode-cn.com/problems/add-strings/)

难度: 简单

两个变量, ans记录返回值, add记录进位, 其余在for中申请, 注意for的条件||, 还有for中边界值>=

14. [最长公共前缀](https://leetcode-cn.com/problems/longest-common-prefix/)



## 哈希表

974. [和可被 K 整除的子数组](https://leetcode-cn.com/problems/subarray-sums-divisible-by-k/)

难度: 中等

通常涉及连续子数组问题, 都使用**前缀和**来解决.

146. [LRU 缓存机制](https://leetcode-cn.com/leetbook/read/2021-spring-recruitment/5f5qts/)

难度: 中等
双向链表+哈希表

560. [和为K的子数组](https://leetcode-cn.com/problems/subarray-sum-equals-k/)

难度: 中等

## 并查集

200. [岛屿数量](https://leetcode-cn.com/problems/number-of-islands/)

难度: 中等

DFS, BFS

56. [合并区间](https://leetcode-cn.com/problems/merge-intervals/)

难度: 中等

需要先排序

990. [等式方程的可满足性](https://leetcode-cn.com/problems/satisfiability-of-equality-equations/)

难度: 中等

## 动态规划

5. [最长回文子串](https://leetcode-cn.com/problems/longest-palindromic-substring/)

难度: 中等

注意状态转移公式 

53. [最大子序和](https://leetcode-cn.com/problems/maximum-subarray/)

难度: 简单

42. [接雨水](https://leetcode-cn.com/problems/trapping-rain-water/)

难度: 困难 TODO

1014. [最佳观光组合](https://leetcode-cn.com/problems/best-sightseeing-pair/)

难度: 中等

注意公式符号

121. [买卖股票的最佳时机](https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock/)

难度: 简单

139. [单词拆分](https://leetcode-cn.com/problems/word-break/)

难度: 中等

## 查找类问题

199. [二叉树的右视图](https://leetcode-cn.com/problems/binary-tree-right-side-view/)

难度: 中等

1. 我的方案: 按层遍历, 即使用queue, 如果是当前层的最后一个节点则加入返回值中.
2. 广度优先
3. 深度优先

46. [全排列](https://leetcode-cn.com/problems/permutations/)

难度: 中等

回溯: 注意数组递归前后交换位置

33. [搜索旋转排序数组](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)

难度: 中等

二分查找, 注意不是普通的二分查找, 因为部分可能是无序的.

79. [单词搜索](https://leetcode-cn.com/problems/word-search/)

难度: 中等

只能使用深度优先搜索 TODO

297. [二叉树的序列化与反序列化](https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/)

难度: 困难

1300. [转变数组后最接近目标值的数组和](https://leetcode-cn.com/problems/sum-of-mutated-array-closest-to-target/)

难度: 中等 TODO


## 排序

215. [数组中的第K个最大元素](https://leetcode-cn.com/problems/kth-largest-element-in-an-array/)

难度: 中等

347. [前 K 个高频元素](https://leetcode-cn.com/problems/top-k-frequent-elements/)

难度: 中等 

## 链表

25. [K 个一组翻转链表](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/)

难度: 困难 TODO

## 递归

124. [二叉树中的最大路径和](https://leetcode-cn.com/problems/binary-tree-maximum-path-sum/)

难度: 困难 TODO

22. [括号生成](https://leetcode-cn.com/problems/generate-parentheses/)

难度: 中等

394. [字符串解码](https://leetcode-cn.com/problems/decode-string/submissions/)

难度: 中等 TODO

236. [二叉树的最近公共祖先](https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/)

难度: 中等 TODO

## 栈和队列

23. [合并K个升序链表](https://leetcode-cn.com/problems/merge-k-sorted-lists/)

难度: 困难

20. [有效的括号](https://leetcode-cn.com/problems/valid-parentheses/)

难度: 简单



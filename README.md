# go-algorithm
基于Go语言的算法实现 algorithm go implement

![[env]](https://img.shields.io/badge/env-darwin-inactive?style=for-the-badge&logo=appveyor) ![go1.14](https://img.shields.io/badge/language-Go1.14-blue?style=for-the-badge&logo=appveyor) ![license MIT](https://img.shields.io/badge/license-MIT-success?style=for-the-badge&logo=appveyor)

欢迎👏访问[libvulcan的博客](https://blog.libvulcan.com)

> 内容持续补充中...

- [x] 回溯算法
  - [x] [生成括号](https://github.com/libvulcan/go-algorithm/blob/main/back_track/generate_brackets.go)
  - [x] [N皇后](https://github.com/libvulcan/go-algorithm/blob/main/back_track/n_queen.go)
  - [x] [全排列](https://github.com/libvulcan/go-algorithm/blob/main/back_track/permutation.go)
  - [x] [K个数字的所有组合](https://github.com/libvulcan/go-algorithm/blob/main/back_track/combine.go)
  - [x] [数组是否能划分为K个元素和相同子集(从数字的视角)](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/back_track/sub_set.go#L9)
  - [x] [数组是否能划分为K个元素和相同子集(从桶/子集的视角)](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/back_track/sub_set.go#L55)
  - [x] [求一组没有重复元素的数组的所有子排列](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/back_track/sub_set.go#L111)

- [x] 二分查找
  - [x] [查找索引](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/binary_search/binary_search.go#L10)
  - [x] [优先查找左边界索引](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/binary_search/binary_search.go#L36)
  - [x] [优先查找右边界索引](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/binary_search/binary_search.go#L64)

- [x] 双指针(原地修改)
  - [x] [数组去重](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/double_pointer/remove_duplicates.go#L12)
  - [x] [链表去重](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/double_pointer/remove_duplicates.go#L29)
  - [x] [批量删除元素](https://github.com/libvulcan/go-algorithm/blob/main/double_pointer/remove_element.go)
  
- [x] 链表操作
  - [x] 快慢指针暴打链表
    - [x] [判断链表中是否有环](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/linked_list/cycle_linked_list.go#L8)
    - [x] [获取环的起始位置](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/linked_list/cycle_linked_list.go#L22)
    - [x] [查找链表中点](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/linked_list/find_linked_list_node.go#L7)
    - [x] [删除链表倒数第K个节点](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/linked_list/find_linked_list_node.go#L20)
  - [x] [反转整个链表](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/linked_list/reverse_linked_list.go#L14)
  - [x] [反转链表前N个节点](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/linked_list/reverse_linked_list.go#L26)
  - [x] [反转链表N~M节点](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/linked_list/reverse_linked_list.go#L40)
  - [x] [反转链表全功能实现(非递归，快慢指针)](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/linked_list/reverse_linked_list.go#L51)

- [x] 前缀和
  - [x] [求数组中一共有多少个和为K的子数组](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/prefix_sum/sub_array_sum.go#L8)
  - [x] [求数组中一共有多少个和为K的子数组(优化解法)](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/prefix_sum/sub_array_sum.go#L30)
  - [x] 差分数组
    - [x] [生成差分数组](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/prefix_sum/diff_array.go#L13)
    - [x] [还原拆分数组](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/prefix_sum/diff_array.go#L26)
    - [x] [给数组区间内的数批量加减值](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/prefix_sum/diff_array.go#L38)
    - [x] [航班预定统计(差分数组应用)](https://github.com/libvulcan/go-algorithm/blob/de0eb2137ad7dac68e8eada35b5ce8b6e1012ff8/prefix_sum/diff_array.go#L50)

- [x] 滑动窗口
  - [x] [找出所有字母的异位词](https://github.com/libvulcan/go-algorithm/blob/main/sliding_window/find_all_anagrams_in_a_string.go)
  - [x] [最长无重复字符的子串长度](https://github.com/libvulcan/go-algorithm/blob/main/sliding_window/longest_substring_without_repeating_characters.go)
  - [x] [最小覆盖子串](https://github.com/libvulcan/go-algorithm/blob/main/sliding_window/min_window_substring.go)
  - [x] [字符串排列(子串)](https://github.com/libvulcan/go-algorithm/blob/main/sliding_window/permutation_string.go)
  
- [ ] [负载均衡算法](https://github.com/libvulcan/go-algorithm/tree/main/load_balance)
  - [x] 轮询
  - [x] 随机
  - [x] IP Hash
  - [x] 加权轮询
  - [x] 平滑加权轮询
  - [ ] ...

- [ ] 排序
  - [x] 冒泡排序
  - [x] 选择排序
  - [x] 归并排序
  - [x] 插入排序
  - [x] 希尔排序
  - [ ] 堆排序
  - [ ] 计数排序
  - [ ] 基数排序
  - [ ] 桶排序

- [ ] 限流算法

- [ ] 二叉树操作

- [ ] 动态规划

- [ ] ...


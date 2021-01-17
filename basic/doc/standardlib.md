## sort
#### sort.Interface
对数据集合（包括自定义数据类型的集合）排序需要实现 sort.Interface 接口的三个方法:
```go
type Interface interface {
    // 获取数据集合元素个数
    Len() int
    // 如果 i 索引的数据小于 j 索引的数据，返回 true，且不会调用下面的 Swap()，即数据升序排序。
    Less(i, j int) bool
    // 交换 i 和 j 索引的两个元素的位置
    Swap(i, j int)
}
```


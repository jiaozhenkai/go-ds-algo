# queue

#### 特点

先进先出

#### 操作

* NewArrayQueue
  * 需要传递容量 cap, cap 范围应该是： 0 <= cap <= MaxUint64-1,  这是因为有
    一个空间的浪费

* Front 返回队头元素

* Back 返回队尾元素

* EnQueue 入队

* DeQueue 出队

* String 格式化输出

* IsEmpty

* IsFull

* Len

#### 适合解决的问题

> TODO

#### 应用场景

> TODO

#### 具体实现

数组，链表

## array queue 

使用切片实现队列

为了避免浪费空间其实就是循环队列

队尾指针 back 处不存放任何值, 浪费这一个存储空间来避免更多的空间的浪费

## linked queue

> TODO

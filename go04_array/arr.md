# 数组遍历

```
sum := 0
for _, v := range numbers{
    sum += v
}
```

* 可通过“_”省略变量
* 不仅仅range, 任何地方都可以通过“_”省略变量
* 如果只要 index, 可以写成 for i:= range numbers{}



# 数组是值类型

* [10]int 和 [20]int 是不同类型
* 调用func f(arr [10]int) 会拷贝数组（值传递，不是引用传递）
* 在go语言中一般不直接使用数组



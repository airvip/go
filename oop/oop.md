# 面向对象

* go 语言仅支持封装，不支持继承和多态
* go 语言没有class, 只有struct


# 值接收者 vs 指针接收者

* 要改变内容必须使用指针接收者
* 结构过大也考虑使用指针接收者
* 一致性：如有指针接收者，最好使用指针接收者

# 封装

* 名字一般使用驼峰法
* 首字母大写：public
* 首字母小写：private

# 包

* 每个目录一个包
* main包包含可执行入口
* 为定义的方法必须放在同一个包内
* 可以放在不同的文件

如何扩充系统类型或者别人的类型

* 定义别名
* 使用组合


#     3
#   /   \
#  0     5
#   \   /
#    2 0
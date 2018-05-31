# map操作

* 创建：make(map[string]int)
* 获取元素： m[key]
* ley不存在时，获取value类型的初始值
* 用value,ok := m[key] 来判断是否存在key
* 用delete删除一个key

# map遍历

* 使用range遍历key,或者遍历key,value对
* 不保证遍历顺序，如需顺序，需要手动对key排序
* 使用len获取元素个数

# map的key

* map使用哈希表，必须可以比较相等
* 除了slice,map,function的内建类型都可以作为key
* Struct类型不包含上述字段也可以作为key
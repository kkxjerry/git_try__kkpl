



# GO开发

```go
		var stockcode="000987"
    var enddate="2020-12-31"
    var url="Code=%s&endDate=%s"
    var target_url=fmt.Sprintf(url,stockcode,enddate)
    fmt.Println(target_url)

		const name, age = "Kim", 22
    s := fmt.Sprintf("%s is %d years old.\n", name, age)
		
```

| %v   | 按值的本来值输出                         |
| ---- | ---------------------------------------- |
| %+v  | 在 %v 基础上，对结构体字段名和值进行展开 |
| %#v  | 输出 Go 语言语法格式的值                 |
| %T   | 输出 Go 语言语法格式的类型和值           |
| %%   | 输出 % 本体                              |
| %b   | 整型以二进制方式显示                     |
| %o   | 整型以八进制方式显示                     |
| %d   | 整型以十进制方式显示                     |
| %x   | 整型以十六进制方式显示                   |
| %X   | 整型以十六进制、字母大写方式显示         |
| %U   | Unicode 字符                             |
| %f   | 浮点数                                   |
| %p   | 指针，十六进制方式显示                   |

```go
name := "John"
    age := 30
    height := 175.5

    result := fmt.Sprintf("Name: %s, Age: %d, Height: %.2f", name, age, height)
    fmt.Println(result)


p := point{1, 2}
    fmt.Printf("%v\n", p)

    fmt.Printf("%+v\n", p)

    fmt.Printf("%#v\n", p)

    fmt.Printf("%T\n", p)


{1 2}
{x:1 y:2}
main.point{x:1, y:2}
main.point
```

字符串格式，可以使用以下对齐参数：

- `%s`：默认对齐方式，左对齐。
- `%10s`：指定宽度为 10 的右对齐。
- `%-10s`：指定宽度为 10 的左对齐。

```go
package main
import "fmt"
func main() {
    var a string = "Runoob"
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)
}
a, b, c := 5, 7, "abc"
```



```go
package main

import "unsafe"
const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a)
)

func main(){
    println(a, b, c)
}

```



```go
package main

import "fmt"

func main() {
    const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i)
}

package main

import "fmt"
const (
    i=1<<iota
    j=3<<iota
    k
    l
)

func main() {
    fmt.Println("i=",i)
    fmt.Println("j=",j)
    fmt.Println("k=",k)
    fmt.Println("l=",l)
}
```

- **k=3**：左移 2 位，变为二进制 **1100**，即 12。
- **l=3**：左移 3 位，变为二进制 **11000**，即 24。



| 运算符 | 描述 | 实例               |
| :----- | :--- | :----------------- |
| +      | 相加 | A + B 输出结果 30  |
| -      | 相减 | A - B 输出结果 -10 |
| *      | 相乘 | A * B 输出结果 200 |
| /      | 相除 | B / A 输出结果 2   |
| %      | 求余 | B % A 输出结果 0   |
| ++     | 自增 | A++ 输出结果 11    |
| --     | 自减 | A-- 输出结果 9     |



| 运算符 | 描述                                                         | 实例              |
| :----- | :----------------------------------------------------------- | :---------------- |
| ==     | 检查两个值是否相等，如果相等返回 True 否则返回 False。       | (A == B) 为 False |
| !=     | 检查两个值是否不相等，如果不相等返回 True 否则返回 False。   | (A != B) 为 True  |
| >      | 检查左边值是否大于右边值，如果是返回 True 否则返回 False。   | (A > B) 为 False  |
| <      | 检查左边值是否小于右边值，如果是返回 True 否则返回 False。   | (A < B) 为 True   |
| >=     | 检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。 | (A >= B) 为 False |
| <=     | 检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。 | (A <= B) 为 True  |



| 运算符 | 描述                                                         | 实例               |
| :----- | :----------------------------------------------------------- | :----------------- |
| &&     | 逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。 | (A && B) 为 False  |
| \|\|   | 逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。 | (A \|\| B) 为 True |
| !      | 逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。 | !(A && B) 为 True  |



| 运算符 | 描述                                           | 实例                                  |
| :----- | :--------------------------------------------- | :------------------------------------ |
| =      | 简单的赋值运算符，将一个表达式的值赋给一个左值 | C = A + B 将 A + B 表达式结果赋值给 C |
| +=     | 相加后再赋值                                   | C += A 等于 C = C + A                 |
| -=     | 相减后再赋值                                   | C -= A 等于 C = C - A                 |
| *=     | 相乘后再赋值                                   | C *= A 等于 C = C * A                 |
| /=     | 相除后再赋值                                   | C /= A 等于 C = C / A                 |
| %=     | 求余后再赋值                                   | C %= A 等于 C = C % A                 |
| <<=    | 左移后赋值                                     | C <<= 2 等于 C = C << 2               |
| >>=    | 右移后赋值                                     | C >>= 2 等于 C = C >> 2               |
| &=     | 按位与后赋值                                   | C &= 2 等于 C = C & 2                 |
| ^=     | 按位异或后赋值                                 | C ^= 2 等于 C = C ^ 2                 |
| \|=    | 按位或后赋值                                   | C \|= 2 等于 C = C \| 2               |



| 优先级 | 运算符           |
| :----- | :--------------- |
| 5      | * / % << >> & &^ |
| 4      | + - \| ^         |
| 3      | == != < <= > >=  |
| 2      | &&               |
| 1      | \|\|             |



```go
package main

import "fmt"

func main() {
    var age int = 23
    if age == 25 {
        fmt.Println("true")
    } else if age < 25 {
        fmt.Println("too small")
    } else {
        fmt.Println("too big")
    }
}
```



### switch

```go
package main

import "fmt"

func main() {
    day := "Monday"
    
    switch day {
    case "Monday":
        fmt.Println("It's Monday!")
    case "Tuesday":
        fmt.Println("It's Tuesday!")
    case "Wednesday":
        fmt.Println("It's Wednesday!")
    case "Thursday":
        fmt.Println("It's Thursday!")
    case "Friday":
        fmt.Println("It's Friday!")
    case "Saturday":
        fmt.Println("It's Saturday!")
    case "Sunday":
        fmt.Println("It's Sunday!")
    default:
        fmt.Println("It's not a valid day of the week.")
    }
}

```

`case`块中的代码执行后，会自动跳出`switch`语句，不会继续执行其他`case`块。

```go
package main

import "fmt"

func main() {
    num := 3

    switch num {
    case 1:
        fmt.Println("One")
        fallthrough
    case 2:
        fmt.Println("Two")
        fallthrough
    case 3:
        fmt.Println("Three")
    }
}

```

由于在块中使用了`fallthrough`关键字，它会继续执行下一个`case 2`块

| 控制语句                                                     | 描述                                             |
| :----------------------------------------------------------- | :----------------------------------------------- |
| [break 语句](https://www.runoob.com/go/go-break-statement.html) | 经常用于中断当前 for 循环或跳出 switch 语句      |
| [continue 语句](https://www.runoob.com/go/go-continue-statement.html) | 跳过当前循环的剩余语句，然后继续进行下一轮循环。 |
| [goto 语句](https://www.runoob.com/go/go-goto-statement.html) | 将控制转移到被标记的语句。                       |

------





### array

```go
var numbers = [5]int{1, 2, 3, 4, 5}
numbers := [5]int{1, 2, 3, 4, 5}

如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度：
var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

```



### Slice 

```go
var numbers = make([]int,3,5)
numbers2 := []int{0,1,2,3,4,5,6,7,8}   //[]中没有写数字
len()长度
cap()容量 
append()扩容，扩容后cap为2的N次方
copy()

```

### map

```go
// 创建一个空的 Map
m := make(map[string]int)

// 创建一个初始容量为 10 的 Map
m := make(map[string]int, 10)

// 使用字面量创建 Map
m := map[string]int{
    "apple": 1,
    "banana": 2,
    "orange": 3,
}
delete(m,"banana")

// 获取键值对
v1 := m["apple"]
v2, ok := m["pear"]  // 如果键不存在，ok 的值为 false，v2 的值为该类型的零值


```

### range

```go
for key, value := range oldMap {}
for key := range oldMap
for _, value := range oldMap
```

Strconv

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "123"
    num, err := strconv.Atoi(str)
    if err != nil {
        fmt.Println("转换错误:", err)
    } else {
        fmt.Printf("字符串 '%s' 转换为整数为：%d\n", str, num)
    }
}
```

​	

```go
package main

import (
    "fmt"
    "strconv"
)

func main() {
    str := "3.14"
    num, err := strconv.ParseFloat(str, 64)
    if err != nil {
        fmt.Println("转换错误:", err)
    } else {
        fmt.Printf("字符串 '%s' 转为浮点型为：%f\n", str, num)
    }
}

package main

import "fmt"

func main() {
    var i interface{} = "Hello, World"
    str, ok := i.(string)
    if ok {
        fmt.Printf("'%s' is a string\n", str)
    } else {
        fmt.Println("conversion failed")
    }
}
```

interface

```go
type DivideError struct {
    dividee int
    divider int
}

// 实现 `error` 接口
func (de *DivideError) Error() string {
    strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
    return fmt.Sprintf(strFormat, de.dividee)
}
//	吟诗实现
package main

import "fmt"

type Shape interface {
    area() float64
}

type Rectangle struct {
    width  float64
    height float64
}

func (r Rectangle) area() float64 {
    return r.width * r.height
}

type Circle struct {
    radius float64
}

func (c Circle) area() float64 {
    return 3.14 * c.radius * c.radius
}

func main() {
    var s Shape

    s = Rectangle{width: 10, height: 5}
    fmt.Printf("矩形面积: %f\n", s.area())

    s = Circle{radius: 3}
    fmt.Printf("圆形面积: %f\n", s.area())
}


随便写



通过reflect.Typeof 获得gettype，reflect.Valueof 获得getvalue，gettype.filed(i) 获得field，field.name ,field.type ,field.tag , value=reflect.Valueof(input).field(i).Interface()

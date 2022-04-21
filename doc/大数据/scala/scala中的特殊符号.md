<-只会出现在for循环里面

->只会出现在k->v里面

简单函数表达式用()

复杂函数表达式用{}

但是=>就较为复杂了，有四种用法

***\*1.表示函数的返回类型(Function Type)\****

```php
//定义函数 
scala> def double(x: Int): Int = x*2
 double: (x: Int)Int

 //定义一个函数变量: 
 scala> var x : (Int) => Int = double
 x: Int => Int = <function1>

 //调用
 scala> x(2)
 res1: Int = 4
```

函数double的类型就是 (x: Int) => Int 或者 Int => Int。***\*左边是参数类型，右边是方法返回值类型\****。
**备注：** 当函数只有一个参数的时候，函数类型里面括起来函数参数的括号是可以省略的。

***\*2.匿名函数\****

```kotlin
//通过匿名函数定义一个函数变量xx
 scala> var xx = (x: Int) => x + 1
 xx: Int => Int = <function1>

 //给一个高阶函数，传递一个函数：
 scala> val newList = List(1,2,3).map { (x: Int) => x * 2 }
 newList: List[Int] = List(2, 4, 6)
```

***\*匿名函数\****定义，=>***\*左边是参数 右边是函数实现体\**** （x: Int）=>{}



***\*3.case语句\****

```kotlin
scala> val x = 10; val y = 20
x: Int = 10
y: Int = 20

scala> val max = x > y match {
  case true => x
  case false => y
}
max: Int = 20
```

在模式匹配 match 和 try-catch 都用 “=>” **表示输出的结果或返回的值**



**4.By-Name Parameters(传名参数)**

传名参数在函数调用前表达式不会被求值，而是会被包裹成一个匿名函数作为函数参数传递下去，例如参数类型为无参函数的参数就是传名参数。

```kotlin
//函数double 
scala> def doubles(x: => Int) = {
     println("Now doubling " + x)
     x*2  
 }
 doubles: (x: => Int)Int

 //调用函数
 scala> doubles(3)
 Now doubling 3
 res2: Int = 6

 scala> def f(x: Int): Int = {
   println(s"Calling f($x)")
   x+1
 }
 f: (x: Int)Int

 //调用函数
 scala> doubles(f(3))
 Calling f(3)
 Now doubling 4
 Calling f(3)
 res9: Int = 8
```

对于函数doubles而言，它的参数x就是by-name的。如果调用doubles的时候，直接给个普通的值或者非函数变量。那么doubles的执行结果就跟普通的函数没有区别。**但是当把一个返回值为Int类型的函数**，例如f(2)，传递给doubles的时候。那么f(2)会被先计算出返回值3，返回值3传入doubles参与运算。运算完成以后得8，**f(2)会被doubles在执行以后，再调用一遍**。
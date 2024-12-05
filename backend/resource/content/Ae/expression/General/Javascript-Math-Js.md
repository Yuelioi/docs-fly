---
title: Javascript-Math-Js
order: 4
---

# JavaScript Math JS数学方法

## abs()

用法：Math.abs(Number)

说明：用于返回某数的绝对值

参数：数值

类型：数值

示例:

```javascript
Math.abs(-8); //返回8。-8的绝对值为8
```

## acos()

用法：Math.acos(Number)

说明：返回一个数的反余弦

参数：数值

类型：数值

图示:

![](https://mir.yuelili.com/user/source/2020/05/cos.jpg)

示例:

```javascript
Math.acos(1); //返回0。1的反余弦为0
```

## acosh()

用法：Math.acosh(Number)

说明：返回一个数的反双曲余弦值

参数：数值

类型： 数值

图示 :

![](https://mir.yuelili.com/user/source/6779/acosh.jpg)

示例:

```javascript
Math.acosh(1); //返回0。1的反双曲线余弦值为0
```

## asin()

用法：Math.asin(Number)

说明：返回数字的反正弦值。 反正弦值是指正弦值为 Number 的角度。 返回的角度以弧度表示，弧度值在 -pi/2 到 pi/2 之间。

参数：Number 为数值

类型：数值（弧度）

```javascript
Math.asin(0.5); //返回 0.523598776
(Math.asin(0.5) * 180) / Math.PI; //返回30。sin30°=0.5，所以把0.5的弧度制转为角度即可得出30°
radiansToDegrees(Math.asin(0.5)); //返回30。同上，只不过使用了弧度转角度radiansToDegrees()函数。
```

## asinh()

用法：Math.asinh(Number)

说明：返回反双曲正弦值

参数：数值

类型：数值

示例：

```javascript
Math.asinh(1); // expected output: 0.881373587019543

Math.asinh(0); // expected output: 0
```

## atan

用法：Math.atan(Number)

说明： 返回一个数值的反正切（以弧度为单位）

参数：数值

类型：返回一个 - π/2 到 π/2 弧度之间的数值

示例：

```javascript
Math.atan(1); // 0.7853981633974483
Math.atan(0); // 0
```

## atan2()

用法：Math.atan2(y,x)

说明：返回从原点(0,0)到(x,y)点的线段与 x 轴正方向之间的平面角度(弧度值)

参数：二维坐标点

类型：数值（弧度角）

示例：

```javascript
Math.atan2(90, 15); // 1.4056476493802699
Math.atan2(15, 90); // 0.16514867741462683
```

## atanh()

用法：Math.atanh(Number)

说明：返回一个数值反双曲正切值

参数：数值

类型：数值

示例：

```javascript
Math.atanh(-2); // NaN
Math.atanh(-1); // -Infinity
Math.atanh(0); // 0
```

## cbrt()

用法：Math.cbrt(number)

说明：返回数字的立方根

参数：数值

类型：数值

示例：

```javascript
Math.cbrt(8); //返回2，因为2的三次方=8
```

## ceil()

用法：Math.ceil(number)

说明：返回大于或等于给定数字的最小整数

参数：数值

类型：数值

示例：

```javascript
Math.ceil(-7.4); //返回-7
Math.ceil(7.4); //返回8
```

## clz32()

用法：Math.clz32(number) ,

说明：返回一个数字在转换成 32 位无符号整形数字的二进制形式后, 开头的 0 的个数

参数：数值

类型：数值

示例：

```javascript
Math.clz32(1000000); //返回12。1000000转换成32位无符号整形数字的二进制形式后是 00000000000011110100001001000000, 开头0的个数是 12 个
```

## cos()

用法：Math.cos(_x_)

说明：返回一个数值的余弦值。

参数：x 为一个以弧度为单位的数值。

类型：余弦数值。

示例：

```javascript
Math.cos(0); // 1
Math.cos(1); // 0.5403023058681398
```

## cosh()

用法：Math.cosh(_x_)

说明：返回数值的双曲余弦函数

参数：x 为数值。

类型：数值。

示例：

```javascript
Math.cosh(0); // 1
Math.cosh(1); // 1.5430806348152437
Math.cosh(-1); // 1.5430806348152437
```

## Math.E

自然常数 e，2.71828....

## exp()

用法：Math.exp(x)

说明：返回 e 的 x 次方，e 是欧拉常数（Euler's constant），自然对数的底，约为 2.718

参数：x 为数值

类型：数值

示例：

```javascript
Math.exp(2); //返回7.3890...
```

## floor()

用法：Math.floor(x)

说明：向下取整

参数：x 为数值

类型：数值

示例：

```javascript
Math.floor(5.05); //5
Math.floor(4); //4
Math.floor(-5.05); //-6
```

## log()

用法：Math.log(x)

说明：Math.log() 函数返回一个数的自然对数

参数：x 为一个数

类型：数值

示例：

```javascript
Math.log(-1); // NaN
Math.log(0); // -Infinity
Math.log(1); // 0
Math.log(10); // 2.302585092994046
```

## log10()

用法：Math.log10(x)

说明：返回以 10 为底的对数。如果传入的参数小于 0, 则返回 NaN.

参数：x 为一个数

类型：数值

示例：

```javascript
Math.log10(10); // 1
Math.log10(100); // 2
Math.log10("100"); // 2
Math.log10(1); // 0
Math.log10(0); // -Infinity
Math.log10(-2); // NaN
```

## log1p()

用法：Math.log1p(x)

说明：返回一个数字加 1 后的自然对数 (底为 E), 既 log(x+1)。如果参数的值小于 -1, 则返回 NaN.

参数：x 为数值

类型：数值

示例：

![log(x+1)](https://mir.yuelili.com/user/source/2020/07/logx1-1.jpg)

```javascript
Math.log1p(Math.E - 1); // 1
Math.log1p(0); // 0
Math.log1p("0"); // 0
Math.log1p(-1); // -Infinity
Math.log1p(-2); // NaN
Math.log1p("hello"); // NaN
```

## log2()

用法：Math.log2(x)

说明：返回以 2 为底的对数.

参数：数值

类型：数值

示例：

```javascript
Math.log2(2); // 1
Math.log2(1024); // 10
Math.log2(1); // 0
Math.log2(0); // -Infinity
Math.log2(-2); // NaN
Math.log2("1024"); // 10
Math.log2("foo"); // NaN
```

## max()

用法：Math.max(value1,value2, ...)

说明：返回一组数字中的最大值

参数：value 为数值

类型：数值

示例：

```javascript
Math.max(10, 20); //返回20
```

## min()

用法：Math.min(value1,value2, ...)

说明：返回一组数字中的最小值

参数：value 为数值

类型：数值

示例：

```javascript
Math.min(10, 20); //返回10
```

## Math.PI

圆周率 π，3.1415....

## pow()

用法：Math.pow(base, exponent)

说明：返回基数（`base`）的指数（`exponent`）次幂，即`baseexponent`。

参数：base 为基数，exp 为指数

类型：数值

示例：

```javascript
Math.pow(4, 0.5); // 类型：2
Math.pow(7, -2); //0.02040816326530612 (也就是1/49)
```

## round()

用法：Math.round(x)

说明：返回一个数字四舍五入后最接近的整数。向绝对值最大的方向入

参数：x 为数值

类型：整数

示例：

```javascript
Math.round(3.5); //返回4
Math.round(-3.5); //返回-3
```

## sign()

用法：Math.sign(x)

说明：返回一个数字的符号, 指示数字是正数，负数还是零。

参数：x 为数值

类型：数值

示例：

```javascript
Math.sign(3); //  1
Math.sign(-3); // -1
Math.sign("-3"); // -1
Math.sign(0); //  0
Math.sign(NaN); // NaN
Math.sign("foo"); // NaN
Math.sign(); // NaN
```

## sin()

用法：Math.sin(x)

说明：返回一个数值的正弦值。

参数：x 为数值

类型：-1 到 1 之间的数值

示例：

```javascript
Math.sin(0); // 0
Math.sin(1); // 0.8414709848078965
Math.sin(Math.PI / 2); // 1
```

## sinh()

用法：Math.sinh(x)

说明：返回一个数字(单位为角度)的双曲正弦值.

参数：x 为角度数值

类型：数值

示例：

![](https://mir.yuelili.com/user/source/2020/07/Sinh.png)

```javascript
Math.sinh(0); // 0
Math.sinh(1); // 1.1752011936438014
Math.sinh("-1"); // -1.1752011936438014
Math.sinh("foo"); // NaN
```

## sqrt()

用法：Math.sqrt(x)

说明：返回一个数的平方根，即：

参数：x 为数值

类型：数值

示例：

```javascript
Math.sqrt(9); // 3
Math.sqrt(2); // 1.414213562373095

Math.sqrt(1); // 1
Math.sqrt(0); // 0
Math.sqrt(-1); // NaN
```

## Math.SQRT2

返回 2 的平方根（约等于 1.414）。

## Math.SQRT1_2

返回返回 2 的平方根的倒数（约等于 0.707）。

## tan()

用法：Math.tan(x)

说明：返回一个数值的正切值。

参数：x 为数值

类型：数值

示例：

```javascript
Math.tan(Math.PI / 4); //返回1,45°的正切值为1
```

## tanh()

用法：Math.tanh(x)

说明：会返回一个数的双曲正切函数值

参数：x 为数值

类型：数值

示例：

```javascript
Math.tanh(0); // 0
Math.tanh(Infinity); // 1
Math.tanh(1); // 0.7615941559557649
```

## trunc()

用法：Math.trunc(value)

说明：去除数字的小数部分，不看正负，直接去

参数：value 为数字

类型：整数

示例：

```javascript
Math.trunc(13.37); //返回13
```

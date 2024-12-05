---
title: Interpolation
order: 3
---

# Interpolation 线性插值

## linear()

用法：linear(t, value1, value2)或 linear(t,tMin,tMax,value1, value2)

说明 1：linear(t, value1, value2)

```javascript
当 t <= 0 时，返回 value1。
当 t 介于 0 到 1 时，返回从 value1 到 value2 的线性插值。
当 t >= 1 时，返回 value2。
```

说明 2：linear(t,tMin,tMax,value1, value2)

```javascript
当 t <= tMin 时，返回 value1。
当 t 介于 tMin 到 tMax 时，返回从 value1 到 value2 的线性插值。
当 t >= tMax 时，返回 value2。
```

参数：t 是数值，value1 和 value2 是数值或数组

类型：数值或数组

图示：当 t 取 time 时(以秒记)

![](https://mir.yuelili.com/user/docs/exp-a-z/exp-linear.png)

```javascript
linear(0.33, 10, 30); //返回16.6 数学解释(线性插值): 10*0.33 + (30-10)*(1-0.33) = 16.6。图形解释:把value1=10,value2=30,t=10帧带入上图，返回10f的value值
linear(1, 10, 30); //返回30
linear(0.5, [10, 30], [100, 200]); //返回[55,115] 各个元素单独计算插值,55=100。
linear(20, 10, 30, 100, 200); //返回150。20处于10和30中间，所以结果为100和200中间
```

## ease()

用法：ease(t, value1, value2)或 ease(t, tMin, tMax, value1, value2)

说明：与[linear()函数](<https://expression.yuelili.fun/interpolation.html#linear()>)一样，只不过在 tMin 和 tMax 点处，进行缓入缓出，使数据更加平滑

参数：t、tMin 和 tMax 是数值，value1 和 value2 是数值或数组

类型：新的数字或数组

示例：

```javascript
详见linear()函数
```

## easeIn()

用法：easeIn(t, value1, value2)或 easeIn(t, tMin, tMax, value1, value2)

说明：与[linear()函数](<https://expression.yuelili.fun/interpolation.html#linear()>)一样，只不过在 tMin 处，进行缓入，使数据更加平滑

参数：t、tMin 和 tMax 是数值，value1 和 value2 是数值或数组

类型：新的数字或数组

示例：

```javascript
详见linear()函数
```

## easeOut()

用法：easeOut(t, value1, value2)或 easeOut(t, tMin, tMax, value1, value2)

说明：与[linear()函数](<https://expression.yuelili.fun/interpolation.html#linear()>)一样，只不过在 tMax 点处，进行缓出，使数据更加平滑

参数：t、tMin 和 tMax 是数值，value1 和 value2 是数值或数组

类型：新的数字或数组

示例：

```javascript
详见linear()函数
```

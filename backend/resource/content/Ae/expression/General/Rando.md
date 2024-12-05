---
title: Random
order: 6
---

# Random 随机

## seedRandom()

用法：seedRandom(seed, timeless=false)

说明：控制随机种子序列。随机方式与唯一图层标识符的函数、图层中的属性、当前时间以及种子值 （seed）有关。一般配合其他随机方法，不然定义了也没用。

参数 seed ：数值。用于控制 随机 函数的初始值。改变该值，可以创建其他随机序列。

参数 timeless ：布尔值。为 true 时，不使用当前时间作为随机种子输入。能够生成一个随机数值，但该值不会随时间而改变。

类型：无

示例：不透明度添加表达式可将不透明度值设为不随时间而改变的随机值：

```javascript
seedRandom(0, true); //timeless为ture ,拖动时间轴，随机数也不会改变，一直是4.72。
random() * 10; //random为0–1范围的随机值，所以表达式返回0~10的随机值
```

![](https://mir.yuelili.com/user/AE/plugins/particular/random-timeless.gif?imageView2/1/w/1674/h/616#)

```javascript
seedRandom(1, true); //timeless为true时,改变种子数可以改变随机值。下图种子从0变为1，数值6.93→9.96
```

![](https://mir.yuelili.com/user/AE/plugins/particular/random-timeless2.gif?imageView2/1/w/1674/h/616#)

## random()

用法：random() 或 random(minNumber[,maxNumber]) 或 random(minValOrArray[, maxValOrArray])

说明：返回 0–1 内的随机数，或 min-max 内的数值或数组

参数：无参数或数值或数组。只有一个参数时,剩下参数默认为 0

类型：数值或数组

示例：

![](https://mir.yuelili.com/user/AE/plugins/particular/random.bmp)
![](https://mir.yuelili.com/user/AE/plugins/particular/random.bmp)

```javascript
random(); //无参数：返回0.61651651。一个0–1内的随机数随机
random(10, 20); //数值：返回16.15648。一个10–20内的随机数随机
random([1, 2], [10, 20]); //数组：返回一个二维数组。数组第一个元素为1-10间随机，第二个元素为2-20间随机
```

## gaussRandom()

用法：gaussRandom()或 gaussRandom(maxValOrArray)或 gaussRandom(minValOrArray, maxValOrArray)

说明：返回随机数。结果为一个高斯（钟形）分布。大约 90% 的结果都处于范围 min(0 或数值或数组)–max (1 或数值或数组)内，其余 10% 在此范围之外（负无穷到 0 以及 0 到正无穷）

参数：无参数 或数值 或数组

类型：数值或数组

示例：  
![](https://mir.yuelili.com/user/source/2020/06/gaussRandom-sample.png)

```javascript
gaussRandom(); //无参数：高斯分布。一个0~1内的随机数随机
gaussRandom(10, 20); //数值：返回16.15648。一个10~20 内的随机数随机
gaussRandom([1, 2], [10, 20]); //数组：返回一个二维数组。数组第一个元素为1-10间随机，第二个元素为2-20间随机
```

## noise()

用法：noise(valOrArray)

说明：返回范围 -1 到 1 中的数值。noise 不是随机的；基于柏林噪声，这意味着输入两个相同值的返回值也相同。

- 用 random(-10,10)，理论取值是-10 到 10，这一帧是 8，下一帧可能是-10，这是完全随机。

- 但用 noise(time)\*10 ，理论取值也是-10 到 10，但是这帧是 8，下一帧可能是 6.5、9.2。不可能这帧 8，下帧直接到-10 跨度这么大。所以这是相差不大的随机

参数：数值或数组

类型：数值

示例：

| id  | name     | 说明                      |
| --- | -------- | ------------------------- |
| 1   | noise(1) | 始终返回 -0.015323279984  |
| 2   | noise(2) | 始终返回 -0.0068915582945 |
| 3   | noise(3) | 始终返回 0.01913710166916 |
| 4   | noise(4) | 始终返回-0.01711790583265 |
| 5   | noise(5) | 始终返回 -0.0238015213381 |

![](https://mir.yuelili.com/user/AE/expression/exp-2-1.bmp)  
![](https://mir.yuelili.com/user/AE/expression/exp-2-1.bmp?imageView2/1/w/2682/h/484#)

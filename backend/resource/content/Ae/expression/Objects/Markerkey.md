---
title: Markerkey
order: 8
---

# MarkerKey 标记

## 标记设置

AE2019/2020 版本标记设置

![](https://mir.yuelili.com/user/AE/expression/exp-12-1.bmp)

AE2018 版本标记设置

![](https://mir.yuelili.com/user/AE/expression/exp-10-1.bmp)

图层标记点：小键盘\*

合成标记点：shift+ 大键盘的数字（记得英文输入法）或者直接拖

双击标记点弹出上图菜单，右键标志点可以进行删减等操作

注：19 版本缺少了很多功能

## duration

说明：标记的持续时间（以秒为单位）。

类型：数值。

示例

```javascript
thisLayer.marker.key(1).duration;
```

## comment

用法：XX.key.comment

说明：返回标记的评论

参数：无

类型：字符串

示例：

```javascript
thisComp.marker.key(1).comment; //返回本合成第1个标记的评论
thisLayer.marker.key(1).comment; //返回本图层第1个标记的评论
```

## chapter

用法：XX.chapter

说明：图层/合成标记对话框中的章节内容字段。不过 2019 以后去掉了，实在没啥用

参数：无

类型：字符串

示例：

```javascript
thisLayer.marker.key(1).chapter; //你写啥内容，就返回啥内容
```

## url

用法：marker.key().url

说明：标记对话框里的 url 字段。2019 以后没了，因为没啥用

参数：无

类型：字符串

示例：

```javascript
thisLayer.marker.key(1).url;
```

## frameTarget

用法：marker.key().frameTarget

说明：标记对话框中的帧目标内容字段。2019 后没了，因为没啥用

参数：无

类型：字符串

示例：

```javascript
thisLayer.marker.key(1).frameTarget;
```

## eventCuePoint

用法：marker.key().eventCuePoint

说明：标记对话框中的提示点类型设置。对于事件为 True；对于导航为 False。2019 后没了，因为没啥用

参数：无

类型：布尔值

示例：

```javascript
thisLayer.marker.key(1).eventCuePoint;
```

## cuePointName

用法：

说明：标记对话框中的提示点名称内容字段。2019 以后没了，实在没啥用

参数：无

类型：字符串

示例：

```javascript
thisLayer.marker.key(1).cuePointName; //返回图层第1个标记的提示点名称
```

## parameters

说明：标记对话框中的参数名称和参数值内容字段。2019 以后没了，因为没啥用

类型：字符串值的关联数组。

示例：设置好"背景颜色"的参数，则您可以使用以下表达式在最近的标记中访问其值：

```javascript
thisComp.marker.nearestKey(time).parameters["background color"];
```

## protectedRegion

用法：XXComp.marker.key(index).protectedRegion

说明：此值为 True 时，合成标记可看作受保护区域。可以双击合成的标记进行设置。图层标记没有哦。具体有啥用，我也不清楚

类型：布尔值；只读。

```javascript
thisLayer.marker.key(1).protectedRegion; //默认为false
```

Process finished with exit code 0

Updated on 2021 年 10 月 24 日

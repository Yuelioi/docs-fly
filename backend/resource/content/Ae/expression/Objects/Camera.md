---
title: Camera
order: 1
---

# Camera 摄像机

除了 source、effect、mask、width、height、anchorPoint、scale、opacity、audioLevels、timeRemap 以及所有材质属性外，"摄像机"对象具有与"图层"对象相同的属性和方法。

## pointOfInterest

用法：摄像机.transform.pointOfInterest

说明：摄像机的兴趣点属性

参数：无

类型：三维数组

示例：

```javascript
thisComp.layer("Camera 1").transform.pointOfInterest; //返回[1000,1000,0]
```

## zoom

用法：cameraOption.zoom

说明：摄像机中摄像机选项的缩放属性

参数：无

类型：数值

示例 1：

```javascript
thisComp.layer("Camera 1").cameraOption.zoom;
thisComp.layer("摄像机 1").cameraOption.zoom;
```

示例 2：在图层的缩放 scale 属性使用，可在更改图层的 z 位置（深度）或摄像机的缩放 zoom 值时保持图层的相对大小（视觉上没变化）

```javascript
cam = thisComp.activeCamera;
distance = length(sub(transform.position, cam.transform.position));
(transform.scale * distance) / cam.zoom;
```

## depthOfField

用法：XX.depthOfField

说明：摄像机图层的"景深"选项。如果摄像机的景深属性打开，则返回 1；如果景深属性关闭，则返回 0。

参数：无

类型：布尔值（0/1）

示例：

```javascript
thisComp.layer("Camera 1").cameraOption.depthOfField;
```

## focusDistance

用法：无

类型：摄像机属性

说明：摄像机的焦距（以像素为单位）。

参数：无

类型：数值

示例：

```javascript
thisComp.layer("Camera 1").cameraOption.focusDistance;
```

## aperture

说明：返回摄像机的光圈值

类型：数值 （以像素为单位）

示例：

```javascript
thisComp.layer("Camera 1").cameraOption.aperture; // 随便调的光圈值 返回25
```

## blurLevel

用法：XX 摄像机.blurLevel

说明：摄像机的模概层次，百分比为单位

参数：无

类型：数值

示例：

```javascript
thisComp.layer("Camera 1").cameraOption.blurLevel; //默认为100%
```

## irisShape

用法：摄像机.cameraOption.irisShape

说明：摄像机的光圈形状，有快矩形，三角形、正方形直到十边形。

参数：无

类型：形状在参数列表的序号，从 1 开始

示例：

```javascript
thisComp.layer("Camera 1").cameraOption.irisShape; //默认的快矩形在第1位，返回1
```

## irisRotation

用法：摄像机.cameraOption.irisRotation

说明：摄像机的光圈旋转

参数：无

类型：数值

示例：

```javascript
thisComp.layer("Camera 1").cameraOption.irisRotation;
```

## irisRoundness

用法：摄像机.cameraOption.irisRoundness

说明：摄像机光圈圆度

参数：无

类型：数值，百分比为单位

示例：

```javascript
thisComp.layer("Camera 1").cameraOption.irisRoundness;
```

## irisAspectRatio

用法：摄像机..cameraOption.irisAspectRatio

说明：摄像机的光圈长宽比

参数：无

类型：数值

示例：

```javascript
thisComp.layer("Camera 1").cameraOption.irisAspectRatio;
```

## irisDiffractionFringe

用法：摄像机..cameraOption.irisDiffractionFringe

说明：摄像机的光圈衍射条纹

参数：无

类型：数值

示例：

```javascript
thisComp.layer("Camera 1").cameraOption.irisDiffractionFringe;
```

## highlightGain

用法：摄像机.cameraOption.highlightGain

说明：摄像机的高光增益

参数：无

类型：数值

示例：

```javascript
thisComp.layer("Camera 1").cameraOption.highlightGain;
```

## highlightThreshold

用法：摄像机.cameraOption.highlightThreshold

说明：摄像机的高光阈值

参数：无

类型：数值

示例：

```javascript
thisComp.layer("Camera 1").cameraOption.highlightThreshold;
```

## highlightSaturation

用法：摄像机.cameraOption.highlightSaturation

说明：摄像机的高光饱和度

参数：无

类型：数值

示例：

```javascript
thisComp.layer("Camera 1").cameraOption.highlightSaturation;
```

## active

用法：XX.active

类型：属性(状态)

说明：返回图层/属性是否激活的布尔值。何谓激活：摄像机图层的"视频"开关 打开；且当前时间处于范围摄像机图层的入点到摄像机图层的出点范围内，且它是"时间轴"面板中列出的第一个（最高）此类摄像机图层。否则，返回 false。

参数：无参数

类型：布尔值

示例:

```javascript
thisComp.layer("Camera 1").active; //返回True
thisComp.layer("测试").active; //对图层使用。如果小眼睛打开,且当前时间点该图层存在于时间轴上，返回true
effect("Threshold")("Level").active; //对效果使用。如果有该效果，返回true
```

Updated on 2021 年 9 月 9 日

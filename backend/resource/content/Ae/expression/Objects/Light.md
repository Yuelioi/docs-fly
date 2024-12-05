---
title: Light
order: 7
---

# Light 灯光

除了 source、effect、mask、width、height、anchorPoint、scale、opacity、audioLevels、timeRemap 以及所有材质属性外，  
"光灯光"对象具有与"图层"对象相同的属性和方法。

## pointOfInterest

用法：灯光.transform.pointOfInterest

说明：摄像机的兴趣点属性

参数：无

类型：三维数组

示例：

```javascript
thisComp.layer("Camera 1").transform.pointOfInterest; //返回[1000,1000,0]
```

## intensity

用法：XX.intensity

说明：强度值

参数：无

类型：数值

示例：

```javascript
thisComp.layer("Point Light 1").lightOption.intensity;
```

## color

用法：XX.color

说明：灯光的颜色值

参数：无

类型：RGBA 四维数组，比如[1,0,0,1]代表红色

示例：

```javascript
thisComp.layer("Point Light 1").lightOption.color; //灯光的颜色值
```

## coneAngle

用法：lightOption.coneAngle

说明：聚光灯图层的灯光选项-锥形角度

参数：无

类型：数值

示例：

```javascript
thisComp.layer("Spot Light 1").lightOption.coneAngle;
```

## coneFeather

用法：lightOption.coneFeather

说明：聚光灯图层的灯光选项-锥形羽化

参数：无

类型：数值

示例：

```javascript
thisComp.layer("Spot Light 1").lightOption.coneFeather;
```

## shadowDarkness

用法：XX.shadowDarkness

说明：灯光的阴影深度

参数：无

类型：数值

示例：

```javascript
thisComp.layer("Point Light 1").lightOption.shadowDarkness;
```

## shadowDiffusion

用法：XX.shadowDiffusion

说明：灯光的阴影扩散

参数：无

类型：数值

示例：

```javascript
thisComp.layer("Point Light 1").lightOption.shadowDiffusion;
```

Process finished with exit code 0

Updated on 2021 年 9 月 9 日

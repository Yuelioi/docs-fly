---
title: Parameters-Floating-Point-Values
order: 17
---

# 参数和浮点值

有件事必须承认；多年来，尽管给了 8 位的颜色值，但我们在内部却偷偷使用浮点表示法。

没错，即使是过亮的颜色，我们也只告诉你是 '255，255，255'。

好吧，我们不能再活在谎言中了! 给定一个颜色参数(由 After Effects 在你的效果参数数组中传递给你)，此函数返回一个浮点表示，包括任何高动态范围的组件。

## PF_ColorParamSuite1

### PF_GetFloatingPoint ColorFromColorDef

```cpp
PF_ErrPF_GetFloatingPointColorFromColorDef(
PF_ProgPtreffect_ref,
constPF_ParamDef*color_defP,
PF_PixelFloat*fp_colorP);
```

## PF_PointParamSuite1

我们还提供了一种方法来获取点参数的浮点值。

### PF_GetFloatingPoint ValueFromPointDef

```cpp
PF_ErrPF_GetFloatingPointValueFromPointDef(
PF_ProgPtreffect_ref,
constPF_ParamDef*point_defP,
A_FloatPoint*fp_pointP);
```

## PF_AngleParamSuite1

在 CS6.0.2 中，提供了一种方法来获取角度参数的浮点值。

### PF_GetFloatingPoint ValueFromAngleDef

```cpp
PF_ErrPF_GetFloatingPointValueFromAngleDef(
PF_ProgPtreffect_ref,
constPF_ParamDef*angle_defP,
A_FloatLong*fp_valueP);
```

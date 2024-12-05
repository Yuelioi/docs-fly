---
title: Working-With-Paths
order: 21
---

# 与路径一同工作

## 访问路径数据

路径与其他参数类型不同，因为它们的值不能直接访问。除了将它们签出和输入(像层参数一样)，你必须使用我们的路径数据函数套件来获取给定时间路径的详细信息。参见`PF_PathQuerySuite1`和`PF_PathDataSuite`。当路径参数传递给您时，在没签出之前, 不要使用它里面存在的值；虽然删除的路径不可用，进一步的更新是“懒惰地”完成（稍后）；除非您的效果检查出路径，否则它不会看到这些更改。

## 操作路径数据

你也可以使用[AEGP_MaskOutlineSuite3](../aegps/aegp-suites.html)来操纵路径。参见[AEGP套件的作弊效果使用](../aegps/cheating effect-usage-of-aegp-suites.html) 。路径参数被视为不透明的数据块；必须使用 get 和 set 函数来访问和操作它们。像图层参数一样，它们必须被访问的效果签出(和签入！)。

## 顶点

路径顶点比简单点更复杂。所有成员变量都PF_FpLongs（双精度），并且在层的坐标空间中。

## PF_PathVertex

| 成员  | 描述     |
| ----------- | --------------------------- |
| `x`         | 顶点位置 |
| `y`         |            |
| `tan_in_x`  | 传入切点 |
| `tan_in_y`  |                |
| `tan_out_x` | 传出切点 |
| `tan_out_y` |                 |

## PF_PathDataSuite

这个套件提供关于路径(顶点序列)的信息。

### PF_PathIsOpen

如果路径未关闭（如果开始和结束顶点不相同），则返回TRUE'。

```cpp
PF_PathIsOpen(
PF_ProgPtreffect_ref0,
PF_PathOutlinePtrpathP,
PF_Boolean*openPB);
```

### PF_PathNumSegments

检索路径中的段数。N段表示有段`[0,N-1]`；段`J`由顶点`J`和`J+1`定义。

```cpp
PF_PathNumSegments(
PF_ProgPtreffect_ref0,
PF_PathOutlinePtrpathP,
A_long*num_segmentsPL);
```

### PF_PathVertexInfo

检索指定路径的`PF_PathVertex`. 点的范围是`[0,num_segments];`对于封闭路径则是`vertex[0]==vertex[num_segments]`.

```cpp
PF_PathVertexInfo(
PF_ProgPtreffect_ref0,
PF_PathOutlinePtrpathP,
A_longwhich_pointL,
PF_PathVertex*vertexP);
```

### PF_PathPrepareSegLength

这个相当反直觉的函数通知AE你要询问一段的长度（使用下面的`PF_PathGetSegLength`), 最好做好准备。 `frequencyL` 表示希望对长度进行多少次采样；我们的内部效果使用100。

```cpp
PF_PathPrepareSegLength(
PF_ProgPtreffect_ref0,
PF_PathOutlinePtrpathP,
A_longwhich_segL,
A_longfrequencyL,
PF_PathSegPrepPtr*lengthPrepPP);
```

### PF_PathGetSegLength

检索给定段的长度。

```cpp
PF_PathGetSegLength(
PF_ProgPtreffect_ref0,
PF_PathOutlinePtrpathP,
A_longwhich_segL,
PF_PathSegPrepPtr*lengthPrepP0,
PF_FpLong*lengthPF);
```

### PF_PathEvalSegLength

检索沿给定路径段的点lengthF的位置。

```cpp
PF_PathEvalSegLength(
PF_ProgPtreffect_ref0,
PF_PathOutlinePtrpathP,
PF_PathSegPrepPtr*lengthPrepPP0,
A_longwhich_segL,
PF_FpLonglengthF,
PF_FpLong*x,
PF_FpLong*y);
```

### PF_PathEvalSegLengthDeriv1

检索沿给定路径段的点`lengthF`的位置和一阶导数。如果不确定是否需要，请不要使用。数学好难。

```cpp
PF_PathEvalSegLengthDeriv1(
PF_ProgPtreffect_ref0,
PF_PathOutlinePtrpathP,
PF_PathSegPrepPtr*lengthPrepPP0,
A_longwhich_segL,
PF_FpLonglengthF,
PF_FpLong*x,
PF_FpLong*y,
PF_FpLong*deriv1x,
PF_FpLong*deriv1y);
```

### PF_PathCleanupSegLength

当完成计算该段长度时, 请调用此命令，以便AE可以正确清理 `PF_PathSegPrepPtr`.

```cpp
PF_PathCleanupSegLength(
PF_ProgPtreffect_ref0,
PF_PathOutlinePtrpathP,
A_longwhich_segL,
PF_PathSegPrepPtr*lengthPrepPP);
```

### PF_PathIsInverted

如果路径反转，则返回`TRUE`。

```cpp
PF_PathIsInverted(
PF_ProgPtreffect_ref,
PF_PathIDunique_id,
PF_Boolean*invertedB);
```

### PF_PathGetMaskMode

检索给定路径的蒙版模式。

```cpp
PF_PathGetMaskMode(
PF_ProgPtreffect_ref,
PF_PathIDunique_id,
PF_MaskMode*modeP);
```

蒙版模式, 以下之一。

> - PF_MaskMode_NONE`
> - PF_MaskMode_ADD`
> - PF_MaskMode_SUBTRACT`
> - PF_MaskMode_INTERSECT`
> - PF_MaskMode_LIGHTEN`
> - PF_MaskMode_DARKEN`
> - PF_MaskMode_DIFFERENCE`
> - PF_MaskMode_ACCUM`

### PF_PathGetName

检索路径的名称(最长为 `PF_MAX_PATH_NAME_LEN`)。

```cpp
PF_PathGetName(
PF_ProgPtreffect_ref,
PF_PathIDunique_id,
A_char*nameZ);
```

## PF_PathQuerySuite1

这个套件用于识别和访问与效果源层相关的路径。

### PF_NumPaths

检索与效果源层相关的路径数量。

```cpp
PF_NumPaths(
PF_ProgPtreffect_ref,
A_long*num_pathsPL);
```

### PF_PathInfo

检索指定路径的 `PF_PathID`。

```cpp
PF_PathInfo(
PF_ProgPtreffect_ref,
A_longindexL,
PF_PathID*unique_idP);
```

### PF_CheckoutPath

在指定的时间获取路径的 `PF_PathOutlinePtr`。

```cpp
PF_CheckoutPath(
PF_ProgPtreffect_ref,
PF_PathIDunique_id,
A_longwhat_time,
A_longtime_step,
A_u_longtime_scale,
PF_PathOutlinePtr*pathPP);
```

### PF_CheckinPath

释放返回AE的路径。无论遇到什么错误情况，始终这样做。每次签出都必须通过签入来平衡，否则后果不堪社想。

```cpp
PF_CheckinPath(
PF_ProgPtreffect_ref,
PF_PathIDunique_id,
PF_BooleanchangedB,
PF_PathOutlinePtrpathP);
```

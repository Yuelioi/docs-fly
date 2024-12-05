---
title: upush
order: 18
---

# upush

## 描述

Adds a uniform item to an array.

```c
void  upush(<type>&array[], <type>value)
```

Pushes a uniform value onto an array for all SIMD processors (regardless of
the processor enabled state).

将一个统一的值推送到所有 SIMD 处理器的数组上（不管处理器的启用状态如何）。

This is a very specialized function, primarily used to manage light exports in
PBR lighting.

这是一个非常专业的函数，主要用于管理 PBR 照明中的光线输出。

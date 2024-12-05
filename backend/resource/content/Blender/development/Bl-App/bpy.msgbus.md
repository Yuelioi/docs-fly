---
title: 消息中枢(bpy.msgbus)
order: 8
---

# 消息中枢(bpy.msgbus)

### Description

**这章不懂，日后再校正一下**

当通过数据（data） API 更改 Blender 数据块属性时，消息中枢系统用于接收通知。

### 限制

消息中枢系统通过更新 RNA 系统触发。意味着以下更新将导致消息总线上的通知：

- 通过 Python API 进行更改，例如.`some_object.location.x += 3`
- 通过用户界面中的滑块、字段和按钮进行更改。

以下更新 **不会** 触发消息中枢通知：

- 在 3D 视窗中移动对象。
- 通过动画系统更改。

### 示例使用

下面是订阅活动对象位置更改的示例。

```python
import bpy

# 任何Python对象都可以作为订阅的拥有者.
owner = object()

subscribe_to = bpy.context.object.location


def msgbus_callback(*args):
    # This will print:
    # Something changed! (1, 2, 3)
    print("Something changed!", args)


bpy.msgbus.subscribe_rna(
    key=subscribe_to,
    owner=owner,
    args=(1, 2, 3),
    notify=msgbus_callback,
)

```

检索某些属性时，它们会转换为 Python 对象。为了创建订阅，需要避免这种情况，使用 datablock.path_resolve("property_name", False)

```python
subscribe_to = bpy.context.object.path_resolve("name", False)

```

也可以在某个类型所有实例的属性上创建订阅：

```python
subscribe_to = (bpy.types.Object, "location")
```

#### clear_by_owner

全名：bpy.msgbus.clear_by_owner(owner)

说明：清除所有使用此拥有者的订阅。

#### publish_rna()

全名：bpy.msgbus.publish_rna(key)

说明：通知订阅者对此属性的更改（通常不需要显式调用，因为更改会自动发布更新）。在某些情况下，使用更通用的 key，明确发布更改可能更好。

参数：

- key（多个）：表示订阅的数据类型。参数包括 [`bpy.types.Property`](https://docs.blender.org/api/current/bpy.types.Property.html#bpy.types.Property)实例。[`bpy.types.Struct`](https://docs.blender.org/api/current/bpy.types.Struct.html#bpy.types.Struct)类型。- ( [`bpy.types.Struct`](https://docs.blender.org/api/current/bpy.types.Struct.html#bpy.types.Struct), str) 类型和属性名称。

#### subscribe_rna()

全名：py.msgbus.subscribe_rna(key, owner, args, notify, options=set())

说明：注册消息中枢订阅。它会在加载另一个 blend 文件时被清除，或者可以通过[`bpy.msgbus.clear_by_owner()`](https://docs.blender.org/api/current/bpy.msgbus.html#bpy.msgbus.clear_by_owner"bpy.msgbus.clear_by_owner")清除.

参数：

- key：多个，表示订阅的数据类型。参数包括 -[`bpy.types.Property`](https://docs.blender.org/api/current/bpy.types.Property.html#bpy.types.Property)实例。-[`bpy.types.Struct`](https://docs.blender.org/api/current/bpy.types.Struct.html#bpy.types.Struct)类型。- ( [`bpy.types.Struct`](https://docs.blender.org/api/current/bpy.types.Struct.html#bpy.types.Struct), str) 类型和属性名称。
- owner：任意类型，处理此订阅（与 ID 进行对比）
- options：字符串集，更改订阅者的行为
  - `PERSISTENT` ，设置时，订阅者将持续保持 ID 数据映射

::: tip
所有订阅均会在文件重载时清除，订阅者可以在加载时重新注册，详见[`bpy.app.handlers.load_post`](https://docs.blender.org/api/current/bpy.app.handlers.html#bpy.app.handlers.load_post"bpy.app.handlers.load_post").
:::

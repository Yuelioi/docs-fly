---
title: 属性对象 properties
order: 6
---

---
title: 属性对象 properties
order: 9
category:

- PR脚本

---

## 属性对象

app.properties

**描述**：_在此处添加说明_

## 属性

无。

## 方法

### clearProperty()

app.properties.clearProperty()

**描述**：_在此处添加说明_

**参数**：_在这里添加参数_

**返回**：_在此处添加返回值/类型_

### doesPropertyExist()

app.properties.doesPropertyExist(property)

**描述**：检查首选项中是否存在给定的属性。

**参数**

| 参数     | 类型   | 描述         |
| -------- | ------ | ------------ |
| property | String | 要检查的属性 |

**返回**：Boolean.

**示例**：检查首选项中是否存在索引为 10 和 99 的标签：

```javascript
var property = "BE.Prefs.LabelNames.10";

var exists = app.properties.doesPropertyExist(property);

alert('Property "' + property + '" exists: ' + exists.toString());

property = "BE.Prefs.LabelNames.99";

exists = app.properties.doesPropertyExist(property);

alert('Property "' + property + '" exists: ' + exists.toString());
```

### getProperty()

app.properties.getProperty(property)

**描述**：返回一个属性值。

**参数**

| 参数     | 类型   | 描述         |
| -------- | ------ | ------------ |
| property | String | 获取值的属性 |

**返回**：String.

**示例**：获取给定索引处的标签名称：

```javascript
var labelIndex = 0;

var property = "BE.Prefs.LabelNames." + labelIndex;

if (app.properties.doesPropertyExist(property)) {
  alert(app.properties.getProperty(property));
} else {
  alert('Property "' + property + '" does not exist');
}
```

### isPropertyReadOnly()

app.properties.isPropertyReadOnly(property)

**描述**：检查给定的属性是否可以被用户覆盖。如果此类属性不存在，则返回 false。

**参数**

| 参数     | 类型   | 描述           |
| -------- | ------ | -------------- |
| property | String | 要检查的属性。 |

**返回**：Boolean.

### setProperty()

app.properties.setProperty(property, value, persistent, createIfNotExist)

**描述**：设置属性值。

**参数**

| 参数             | 类型    | 描述                           |
| ---------------- | ------- | ------------------------------ |
| property         | String  | 要创建的属性                   |
| value            | Any     | 属性值                         |
| persistent       | Boolean | 是否应该在会话之间保持不变     |
| createIfNotExist | Boolean | 应该创建，如果这样的属性不存在 |

**返回**：null.

**示例**：更改标签名称：

```javascript
var labelIndex = 0;

var property = "BE.Prefs.LabelNamesX." + labelIndex;

var newValue = "Changed via Script";

var persistent = true;

var createIfNotExist = true;

if (app.properties.doesPropertyExist(property)) {
  if (app.properties.isPropertyReadOnly(property)) {
    alert('Could not rename property "' + property + '" because it is read-only.');
  } else {
    var oldValue = app.properties.getProperty(property);

    app.properties.setProperty(property, newValue, persistent, createIfNotExist);

    alert('Value changed from "' + oldValue + '" to "' + newValue + '"');
  }
} else {
  app.properties.setProperty(property, newValue, persistent, createIfNotExist);

  alert('Created new property "' + property + '" with value "' + newValue + '"');
}
```

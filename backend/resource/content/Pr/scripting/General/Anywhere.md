---
title: Anywhere
order: 2
---

---
title: Anywhere
order: 2
category:

- PR脚本

---

## Anywhere 对象

app.anywhere

**描述**：在 anywhere 对象表示任何 Adobe Anywhere 或团队项目服务器可用。

## 属性

无。

## 方法

### getAuthenticationToken()

app.anywhere.getAuthenticationToken()

**描述**：检索身份验证令牌。

**参数**：无。

**返回**：字符串。包含的登录标记，如果不成功则为 0。

### getCurrentEditingSessionActiveSequenceURL()

app.anywhere.getCurrentEditingSessionActiveSequenceURL()

**描述**：检索作品中当前活动序列的 URL。

**参数**：无。

**返回**：字符串。返回一个包含资产的 URL，如果不成功则为 0（包括没有活动序列，或者没有打开编辑会话）。

### getCurrentEditingSessionSelectionURL()

app.anywhere.getCurrentEditingSessionSelectionURL()

**描述**：检索当前选定的单个资产的 URL。如果选择的项目多于或少于一项，则会失败。

**参数**：无。

**返回**：字符串。返回一个包含资产的 URL，如果不成功则为 0（包括如果多个项目或 fewre 被选中）。

### getCurrentEditingSessionURL()

app.anywhere.getCurrentEditingSessionURL()

**描述**：检索当前正在编辑的作品的 URL。

**参数**：无。

**返回**：返回一个包含生产 URL 的**字符串**，如果不成功则返回**0**。

### isProductionOpen()

app.anywhere.isProductionOpen()

**描述**：检索 Anywhere 或 Team Projects 作品当前是否处于打开状态。

**参数**：无。

**返回**：如果制作已打开，则返回 true；否则为 false。

### listProductions()

app.anywhere.listProductions()

**描述**：在当前服务器上检索当前用户可用的产品名称。

**参数**：无。

**返回**：返回包含可用产品名称的**字符串**数组，如果不成功则返回 0。

### openProduction()

app.anywhere.openProduction(productionURL)

**描述**：在指定的 URL 打开生产。

**参数**

| 参数          | 类型   | 描述                 |
| ------------- | ------ | -------------------- |
| productionURL | String | 要打开的产品的 url。 |

**返回**：如果成功则返回**0**。

### setAuthenticationToken()

app.anywhere.setAuthenticationToken(token, emailAddress)

**描述**：使用提供的令牌将指定的电子邮件地址记录到服务器中。

**参数**

| 参数         | 类型   | 描述                 |
| ------------ | ------ | -------------------- |
| token        | String | 授权令牌。           |
| emailAddress | String | 关联的电子邮件地址。 |

**返回**：如果成功则返回**0**。

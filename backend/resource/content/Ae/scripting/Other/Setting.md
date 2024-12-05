---
title: Setting
order: 6
---

# Setting 脚本设置

## 脚本设置

app.settings

描述：设置对象提供了一种简便的方法来管理第三方脚本。这些设置保存在 After Effects 首选项主文件中，并且在应用程序会话之间保持不变。

设置由文件中的节（section）和键（key）标识，并且每个键名都与一个值相关联。

在设置文件中，节名括在方括号和引号中，键名称列在节名下方的引号中。所有值都是字符串。

从版本 12 / CC 开始，如果 Section /
Key 不在主首选项文件中，则首选项和设置方法使用第三个参数来指定目标首选项文件。有关更多信息，请参见“首选项”对象。

::: info 提示

- 这些设置在 AE 各版本之间不共享。每次新安装 AE，设置文件也会更新，因此这些设置将失效。
- 在内部，所有保存设置的节名都以 "Settings\_"开头
- 获取或设置内部 AE 首选项，请参阅“首选项”对象
  :::

技巧

- 每个脚本使用一个 sectionName。这样可以井井有条，易于查找和使用。
- 将设置保存到特定的首选项文件中并无任何好处。省略第三个参数（prefType），使用默认值即可。

## 方法

### getSetting()

app.settings.getSetting(sectionName, keyName[, prefType])

描述：从首选项文件中检索脚本设置项的值。

::: danger
如果大于 1999 个字节，则报错(在 AE 15.0.1)
:::

参数：

- sectionName：节名称，字符串。
- keyName ：键名称，字符串。
- prefType ：可选，枚举，使用哪个首选项文件。详见首选项对象

返回：字符串。

示例

```javascript
alert(app.settings.getSetting("test_section", "test_key"));
```

### haveSetting()

app.settings.haveSetting(sectionName, keyName[, prefType])

描述：如果指定的脚本存在设置项并且有值，则返回 true。

参数：

- sectionName：设节名称，字符串。
- keyName ：设键名称，字符串。
- prefType ：可选，枚举，使用哪个首选项文件。详见首选项对象

返回：布尔值。

示例：查找 my_section 节有没有 my_key 键

```javascript
app.settings.haveSetting("test_section", "test_key");
```

### saveSetting()

app.settings.saveSetting(sectionName, keyName, value[, prefType])

描述：保存脚本设置项的值。

::: danger
:::

如果大于 1999 个字节，则报错(在 AE 15.0.1)

参数：

- sectionName：节名称，字符串。
- keyName ：键名称，字符串。
- value：新值，字符串。
- prefType ：可选，枚举，使用哪个首选项文件。详见首选项对象

返回：无。

示例：设置 my_section 的 my_key 为 false

```javascript
app.settings.saveSetting("my_section", "my_key", false);
```

注意！虽然写的是 false，但是保存会编程字符串。想要获得真正的布尔值，读取的时候需要自己解析下

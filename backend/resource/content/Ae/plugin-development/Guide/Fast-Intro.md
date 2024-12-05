---
title: Fast-Intro
order: 1
---

# 快速入门

## 示例文件

使用 vs 打开 SDK 文件夹的 `Template > Skeleton > Win > Skeleton.vcxproj`

::: tip
请先备份 Template 文件夹 方便后续创建新模板
:::

## 配置

<https://youtu.be/uHbbW3qtTJQ>

### 环境变量设置

AE_PLUGIN_BUILD_DIR(插件生成目录) : C:\Program Files\Adobe\Common\Plug-ins\7.0\MediaCore\YLL(最后面的文件夹自己定义)

![环境变量](/images\environment.png)

TargetPath(调试器目录,可选) : 你的 ae 路径

![1665287136892](/images/1665287136892.png)

![环境变量](/images\property.png)

### VS其他配置项

C++ > 常规: 将警告视为错误 改为 `否`.

C++ > 代码生成: 结构成员对齐 改为 `默认`.

连接器 > 附加依赖项: 自己加依赖路径即可.

## 编译

右键这玩意 点编译

![1665220885176](/images/1665220885176.png)

## 报错

### 无法打开源文件 stdint.h

vs 组件安少了, 在 Visual Studio Installer 安装: 使用 c++的桌面开发

![1665286881975](/images/1665286881975.png)

### 无法找到 VS 2015 生成工具

改成自己工具版本即可

![img](/images\tool.png)

### 没有访问权限

原来是我还没有新建 YLL 文件夹,新建一下即可,或者自己设置管理员权限

![img](/images\folder.png)

### C2338 报错

![img](/images\C2338.png)

或者

预处理器定义前追加 `WINDOWS_IGNORE_PACKING_MISMATCH;`

![1665286593203](/images/1665286593203.png)

### 无法打开文件 Sketon.aex

可能是 C 盘文件读取限制, 使用管理员打开 VS 即可,或者给文件夹设置管理员权限

`C:\Program Files\Adobe\Common\Plug-ins\7.0\MediaCore`

![1665286787845](/images/1665286787845.png)

### 其他报错

干脆把警告视为报错改为否

![1665286453915](/images/1665286453915.png)

## 调试

还是属性设置, 设置为 AE 路径

![img](/images\ae.png)

或者在配置 > 环境变量配置里, 追加环境变量 TargetPath : 你的 ae 路径

![1665287136892](/images/1665287136892.png)

F5 运行,或者自己点 `本地 Windows 调试器`

![1665220525248](/images/1665220525248.png)

## 其他

修改项目名称,

Sketon 改为 LetterBox

![1665285887802](/images/1665285887802.png)

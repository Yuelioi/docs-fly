---
title: VS Code环境配置
order: 1
---

# 【BL开发】VS Code环境配置

## Description

<iframe src="https://player.bilibili.com/player.html?bvid=BV1DV411z7vp&page=2&high_quality=1" width="100%" allowfullscreen="allowfullscreen" frameborder="0"></iframe>

### 1.安装插件

搜 blender，就 4 个。自己找下

![](https://cdn.yuelili.com/20220110165253.png)

### 2.插件配置

在 settings.json 里加入下面

![](https://cdn.yuelili.com/20220110165508.png)

```json
        "blender.executables": [

            {

                "path": "D:\\software\\Blender\\blender.exe",

                "name":"",

                "isDebug": false

            }

        ]

```

## 3.自动补全插件

[github 下载](https://github.com/nutti/fake-bpy-module)

然后在上面同个配置下添加你的刚才下载的文件路径。或者设置里搜索 Auto Complete

"python.autoComplete.extraPaths": [

"C:\\Users\\YL\\Downloads\\fake_bpy_modules3.0"

]

## 4.创建新插件

F1 找到新建插件，然后按要求输入插件名、作者、插件路径

![](https://cdn.yuelili.com/20220110172841.png)

然后自动创建的文件如下

![](https://cdn.yuelili.com/20220110173012.png)

## 5.在 BL 中运行

按 F1（或者 Ctrl shift P）: 输入 blender

![](https://cdn.yuelili.com/20220110173050.png)

## 6.如何打包安装插件

直接把整个目录弄成 zip 即可

## 可能出现的问题（重要！）

不要在 BL 首选项删除插件，这样会删除源码！请去 bl 的插件目录删除（本 vsc 插件是链接一个快捷方式到 bl 插件目录），删除快捷方式即可

## 其他

如果第一次启动失败，也许是因为挂了 tt，导致 host failed，没法下载插件

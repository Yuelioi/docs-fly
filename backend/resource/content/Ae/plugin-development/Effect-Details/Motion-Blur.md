---
title: Motion-Blur
order: 14
---

# 动态模糊

效果处理它们自己的运动模糊，使用[PF_InData>shutter_angle](../effect-basics/PF_InData.html)以及[PF_InData>shutter_phase](../effect-basics/PF_InData.html) 。

插件必须设置[PF_OutFlag_I_USE_SHUTTER_ANGLE](../effect-basics/PF_OutData.html) ，以便 After Effects 知道它需要这个信息。

它们必须在其他时间[check out](interaction-callback-functions.html) 自己的参数，以检查它们在快门间隔时间的变化。

如果插件在这个时间间隔之外检查参数，请设置[PF_OutFlag_WIDE_TIME_INPUT](../effect-basics/PF_OutData.html) 。

这样做可以让 After Effects 比较采样间隔内的参数，并确定它们是否有变化。

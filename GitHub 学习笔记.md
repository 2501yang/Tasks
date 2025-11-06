# **GitHub 学习笔记**
**我的地址**https://github.com/2501yang/Tasks.git
## 一、GitHub 是什么

GitHub 是全球最大的「程序员代码协作平台」+「开源项目托管库」

1. 代码云盘：存项目代码，不怕丢，支持版本回溯；

2. 协作工具：和队友一起写代码，解决冲突，合并成果；

3. 技术社区：看大佬项目学技术，分享自己的作品。

## 二、GitHub 核心概念

1. 仓库（Repo）像一个项目的“文件夹”，包含代码、文档、图片等所有文件。
2. 提交（Commit）对仓库的一次修改记录。像文档的“保存”+“备注”，记录每一步改动。
3. 分支（Branch）从主代码“分叉”出的独立修改空间，用于开发新功能或修复 bug，不影响主代码。像“草稿纸”，写完再合并到“正式文档”里。
4. 拉取请求（PR）当你在分支完成修改后，向主代码发起“合并申请”，队友可审核代码再决定是否合并。像“交作业”——你改完了，让组长检查后合并到班级作业里。
5. 合并（Merge）审核通过后，将分支的修改整合到主代码中。组长把你的“草稿”内容抄到“正式文档”里。 

## 三、GitHub 基础操作

访问官网：https://github.com；

1. 创建第一个仓库（Repo）

点击右上角「+」→「New repository」；

Repository name：仓库名（建议英文/拼音，比如 my-first-repo）；

Description：简单描述（比如“我的第一个 GitHub 项目”）；

Public/Private：公开（免费）或私有（需付费，萌新选 Public）；

勾选 Initialize this repository with a README：自动生成 README 文件（推荐！）；

点击「Create repository」→ 仓库创建成功！

2. 上传本地文件到 GitHub

你需要把电脑里的代码/文档上传到 GitHub 仓库

进入你的仓库页面；

点击「Add file」→「Upload files」；

选中本地文件（可多选），拖入上传框；

填写提交说明（比如“添加项目介绍文档”）；

点击「Commit changes」→ 文件上传成功！

3. 查看提交记录 & 回退版本

目标：找回误删的代码，或回到之前的版本。

步骤：

进入仓库页面，进入文档，点击「history」标签；

看到所有提交记录（时间、提交人、说明）；

可查看该版本的文件详情；

4. 创建分支 & 合并分支

创建分支：

仓库页面点击「Branch: main」→ 输入新分支名→ 点击「Create branch」；

或命令行输入 git checkout -b feature/login → 创建并切换到新分支。

在分支上修改代码：

本地切换到分支：git checkout feature/login；

修改文件后，提交到分支：git add . → git commit -m "新增登录功能"。

合并分支到主代码：

先切换回主分支：git checkout main；

拉取最新主代码（避免冲突）：git pull origin main；

合并分支：git merge feature/login；

推送到 GitHub：git push origin main。

## 四、GitHub 进阶技巧

1. 必写 README.md（项目的“说明书”）

作用：别人点进仓库，第一眼看到的就是它！用来介绍项目功能、使用方法、贡献指南等。

内容建议：

1. 使用 Issues 跟踪问题

作用：记录项目 bug、需求或讨论（比如“登录页面样式错位”）。

操作：仓库页面点击「Issues」→「New issue」→ 填写标题和描述 → 提交。

2. 添加 LICENSE（保护代码版权）

作用：声明代码的使用权限（比如“MIT 协议”允许他人免费使用）。

操作：仓库页面点击「Add file」→「Create new file」→ 文件名填 LICENSE → 选协议模板（MIT）。

## 五、个人主页学习模板
**我的个人主页参考**
1. https://blog.csdn.net/m0_73784704/article/details/151047622
2. https://www.cnblogs.com/PeterJXL/p/18437094 <br>
 1. ![使用语言和活动统计]![阶段一参考截图1](https://github.com/user-attachments/assets/59f16f53-3e23-4ec4-9025-60e387e728c9)


 2. ![打字特效]![阶段一参考截图2](https://github.com/user-attachments/assets/567b244e-5118-48da-bc0e-65cc140e8866)


 3. ![github徽章]![阶段一参考截图3](https://github.com/user-attachments/assets/40efe1a5-dfb1-4fd2-9694-23c4adf5ce4c)

## 六、总结
GitHub 是工具，核心是“用它解决问题”。遇到报错别慌，复制错误信息去 Google 等搜索引擎搜。
（附：GitHub 官方文档 → https://docs.github.com/zh）






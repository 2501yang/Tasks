# **Git 学习笔记**

***



在进行任何 Git 操作之前，都要先切换到 Git 仓库目录，也就是先要先切换到项目的文件夹目录下。

并在文件夹夹里创建一个文件

然后进入这个文件夹，把它当作Git 仓库，右击选择**Git Bash Here**：

![](C:\Users\ASUS-jxy\Pictures\Screenshots\屏幕截图 2025-11-09 170704.png)

打开的 GitBash 可以看到自动定位到仓库位置，如果在其他位置打开，还得再次定位，这就是在仓库/文件夹打开的原因。

**2.1 git status**

输入 **git status** 命令，查看仓库状态。

**2.2 git init**

若它还不是 Git 仓库，就用 **git init** 初始化。

再用 **git status** 命令检查一下仓库，发现 Git 已经成为一个 Git 仓库了，并且默认进入 Git 仓库的 main分支

返回的条语句：

- 进入 Git 仓库的主分支
- 删除未暂存的变更
- 这时最主要的是提示文件是 Untracked files ，就是说这个文件还没有被跟踪，还没有提交在 git 仓库里呢，而且提示你可以使用 git add 去操作你想要提交的文件。

![image-20251109172551008](C:\Users\ASUS-jxy\AppData\Roaming\Typora\typora-user-images\image-20251109172551008.png)

**2.3 git add**

接下来往仓库里添加文件，根据提示，输入 **git add +提示文件**

虽然仓库没有提交历史，但提示以下文件 Changes to be committed ，意思就是文件等待被提交。

可以使用 git rm --cached 这个命令去移除这个缓存。

**2.4 git commit**

把文件送到缓冲区，接下来用 **git commit** 命令将缓存区里的改动给提交到本地的版本库。

每次使用 **git commit** 命令都会在本地版本库生成一个 40 位的哈希值，这个哈希值也叫 commit-id（这个版本的编号），commit-id 在版本回退的时候是非常有用的，它相当于一个索引，可以在未来的任何时候通过与 git reset 的组合命令回到这里。

**git commit** 不是单独使用的，一般需要参数指定提交方式，输入 **git commit --help** 查看使用帮助（会跳转到相关网页）。

输入 **git commit -m ” commit“**，成功将文件提交到了 Git 仓库。其中， -m 代表提交信息， commit 是本次提交的信息，需要写在双引号内，提交信息你也可以写成 my commit，text commit 等任意句子。

仓库状态结果显示 nothing to commit, working tree clean，这表示已经没有内容可以提交了，即全部内容已经提交完毕。

补充**git commit**常见的用法

- **git commit --amend**，追加提交，它可以在不增加一个新的commit-id的情况下将新修改的代码追加到前一次的commit-id中

### 查看作者、时间

**2.5 git log**

输入**git log**命令，打印 Git 仓库提交日志，会显示作者、时间和你写的提交信息

### 创建、切换、合并分支

**2.6 git branch**

git init 会初始化一个主分支,也是现在的默认分支。

输入 **git branch**，可以看到在 main/master 这里：

在此基础上新建分支用**git branch a**，其中 a 是任意起的分支名字。然后再用**git branch**查看，可以看到 master 有了一个分支，这时候分支 a 跟分支 master 是一模一样的内容。

**2.7 git checkout**

虽然建立分支，但可以看到 * 还在 main 上，说明我们还处于主分支。这时候，输入 **git checkout a** 命令，切换到分支 a 上，再次查看分支，* 在 a 前，说明我们处于分支 a ，可以修改分支 a 的内容了。

> 先新建再切换很麻烦，有一步到位的方法**git checkout -b a**，意思是建立分支后自动切换到该分支。

**2.8 git merge**

你在 a 分支完成了你的部分，但是主分支还是原来那样，这个时候就需要把你的代码合并到主分支 master上来，**git merge** 就是合并分支用到的命令。

**注意，合并分两步**

1. 切换到 main分支，如果你已经在了就不用切换了;
2. 执行 git merge a ，意思就是把 a 分支的代码合并过来，不出意外，
   这个时候 a 分支的代码就顺利合并到 master 分支来了。

**删除分支**：有添加就有删除，假如分支新建错了，或者 a 分支的代码已经**顺利合并**到 master 分支来了，那么 a 分支没用了，用 **git branch -d** 删除。删除分支 a 的语句为 git branch -d a

**强制删除分支**：上面说了 **git branch -d** 是删除分支的意思，有些时候可能会删除失败，比如如果 a 分支的代码还**没有合并**到master，你执行 git branch -d a 是删除不了的，它会智能的提示你a分支还有未合并的代码，但是如果你非要删除，那就执行 **git branch -D a** 就可以强制删除 a 分支。

以上在本地环境进行操作的，最后总结一下基本 Git 操作：

- git status ：查看仓库状态
- git init：创建一个空仓库，或者重新初始化一个已有仓库
- git add：把文件添加到可提交列表（临时缓冲区）
- git commit：提交改动（增删改）至仓库
- git log：打印提交日志
- git branch：查看、添加、删除分支
- git checkout：切换分支、标签
- git merge：合并分支

## 本地与远程仓库的交互

### 1. SSH 配置

要想本地与远程仓库交互，必须有一个安全机制防止数据泄漏，这个安全机制就是 SSH

**1.1 生成 SSH key**

打开 GitBash ，输入 ssh，如果返回如下界面说明 SSH 已安装

紧接着输入 **ssh-keygen -t rsa** （指定 rsa 算法生成密钥），接着连续三个回
车键（不需要输入密码），然后就会生成两个文件 id_rsa （密钥）和 id_rsa.pub （公钥）。

**1.2 添加 SSH key**

接下来要做的是把 id_rsa.pub 的内容添加到 GitHub 上，这样你本地的 id_rsa 密钥跟 GitHub 上的 id_rsa.pub 公钥进行配对，授权成功才可以提交代码。

打开 GitHub，点击右上角的你的头像，点击设置 **settings**：

点击左侧的**SSH and GPG keys**，点击右上角的**New SSH key**：

进入到下面的界面，把刚刚生成的公钥 （id_rsa.pub 文件中的内容）复制到 Key 所在框，复制好点击下方的 **Add SSH key** 按钮即可。

提示：id_rsa.pub 可以选择用 Notepad++ 或 Sublime 打开~

添加完返回 SSH 页面就会出现你的本地信息，如下图所示：![](C:\Users\ASUS-jxy\Pictures\Screenshots\屏幕截图 2025-11-09 174227.png)

**1.3 验证绑定**

最后验证一下本地的 git 与 GitHub 是否绑定成功，在 GitBash 输入**ssh -Tgit@github.com**

到此为止，SSH 配置完成

**2.1 git clone with HTTPS**

注意上面写的是 ` HTTPS`，复制下面的 URL：

打开文件要 clone 的目录，右击选择`Git Bash`Here:

打开 GitBash 窗口，输入 **git clone**，后面跟刚刚复制的 URL，回车，等它下载完，你会看到你的文件夹里直接就是仓库里的文件。

语句长这样：

```text
git clone https://github.com/project/repo.git
```

**2.2 git clone with SSH**

在 ` HTTPS`那里，点击 ` SSH `，就会切换成第二种 git clone 方式。

其实步骤与上面的步骤一样，只不过复制的 URL 是 SSH 机制下的：

接着也是在目录下打开 Git Bash，输入 **git clone**， 后面是刚刚复制的 URL。

语句长这样， SSH 下的 URL 和 HTTPs 下的 URL 不一样。

```text
git clone git@github.com:project/repo.git
```

![](C:\Users\ASUS-jxy\Pictures\Screenshots\屏幕截图 2025-11-09 174429.png)

**2.3 异同**

同：HTTPs 与 SSH 下的 git 都可以直接进行 git clone 操作

异：

- HTTPs git clone 到本地，进行了一些文件的修改，当再次提交到 GitHub远程服务器的时候，都会进行账号与密码的输入
- SSH git clone 到本地之后，由于已有 SSH Keys 授权，就不需要用户名和密码进行授权了。

### 3. 通过 Git 提交代码

- git push：翻译为推，当你的代码更新，需要把本地的推到远程仓库
- git pull：翻译为拉，当远程仓库有更新，你需要把远程的拉到本地进行合并

前者在本地无仓库时使用，后者是本地已有仓库时使用。

使用举例：

```text
git push origin main # 把本地代码推到远程 master 分支
git pull origin main # 把远程最新的代码更新到本地
```

**一般在 push 之前先 pull ，这样不容易冲突。**

**所有的改动都需要通过 `git add`和`git commit` 提交到本地的仓库。**

**在这个文件夹/本地仓库**，右击打开 **Git Bash**，输入 `git status` 查看状态，再用 `git add`和`git commit` 提交新加的文件。

然后刷新github远程仓库，可以看到有更新

#### 显示改动

 `git add` 会把文件放入暂存区，**在`git add`** **之前**使用 `git diff` 可以显示对某一文件的改动。

红色的部分前面有个 - 代表我删除的，绿色的部分前面有个 + 代表我增加的，从这里一目了然的知道这个文件做了哪些改动。

除此之外，`git diff`还有其他用法：

```text
git diff <$id1> <$id2> # 比较两次提交之间的差异
git diff <branch1>..<branch2> # 在两个分支之间比较
git diff --staged # 比较暂存区和版本库差异
```

#### 版本回滚

想要回滚到某次提交之前的场景

首先用 `git log `查看版本号（commit 后面那一长串就是版本号）：

找到想要回退的版本号之后，在本地 Git 仓库执行如下命令，选一个即可：

```text
git reset --hard 版本号 (抛弃当前工作区的修改)
git reset --soft 版本号 (回退到之前的版本，但保留当前工作区的修改，可以重新提交)
```

如果文件只在本地，执行上面的步骤就可以

如果你的文件提交到 GitHub 上，还需要执行如下命令同步到远端：

```text
git push origin main
```

本地的版本落后于远端的版本，因此我们还需要在上述命令中加上`--force`参数：

```text
git push origin main --force
```

# Git 提交文件夹总结

### 创建一个版本库

1. 选择一个合适的地方，创建一个空目录：

```plain
$ mkdir learngit
$ cd learngit
$ pwd
/Users/michael/learngit
```

**`pwd`命令用于显示当前目录**，位于`/Users/michael/learngit`。

2. 通过**`git init`命令把这个目录变成Git可以管理的仓库**：

```plain
$ git init
Initialized empty Git repository in /Users/michael/learngit/.git/
```

### 把文件添加到版本库

例如编写一个`readme.txt`文件，内容如下：

```plain
Git is a version control system.
Git is free software.
```

1. 用命令`git add`告诉Git，把文件添加到仓库：

```plain
$ git add readme.txt
```

2. 用命令`git commit`告诉Git，把文件提交到仓库：

```plain
$ git commit -m "wrote a readme file"
[master (root-commit) eaadf4e] wrote a readme file
 1 file changed, 2 insertions(+)
 create mode 100644 readme.txt
```

解释一下`git commit`命令，`-m`后面输入的是本次提交的说明，可以输入任意内容，这样就能从历史记录里方便地找到改动记录。

`git commit`命令执行成功后会告诉，`n file changed`：n个文件被改动；`n insertions`：插入了n行内容（取决于有输入几行内容）。

Git添加文件需要`add`，`commit`一共两步，因为`commit`可以一次提交很多文件，所以你可以多次`add`不同的文件，比如：

```plain
$ git add file1.txt
$ git add file2.txt file3.txt
$ git commit -m "add 3 files."
```

# 快速版

首先进入这个文件夹，右击点击 Git Bash Here。

![](C:\Users\ASUS-jxy\Pictures\Screenshots\屏幕截图 2025-11-09 170704.png)

依次输入下面的语句：

```text
git init   // 初始化版本库

git add .   // 添加文件到版本库（只是添加到缓存区），.代表添加文件夹下所有文件 

git commit -m "first commit" // 把添加的文件提交到版本库，并填写提交备注

git remote add origin 你的远程库地址  // 把本地库与远程库关联

git push -u origin main    // 第一次推送时

git push origin main  // 第一次推送后，直接使用该命令即可推送修改
```
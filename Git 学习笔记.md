# **Git 学习笔记**

***

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

简单解释一下`git commit`命令，`-m`后面输入的是本次提交的说明，可以输入任意内容，这样就能从历史记录里方便地找到改动记录。

`git commit`命令执行成功后会告诉你，`1 file changed`：1个文件被改动（我们新添加的readme.txt文件）；`2 insertions`：插入了两行内容（取决于readme.txt有输入几行内容）。

Git添加文件需要`add`，`commit`一共两步，因为`commit`可以一次提交很多文件，所以你可以多次`add`不同的文件，比如：

```plain
$ git add file1.txt
$ git add file2.txt file3.txt
$ git commit -m "add 3 files."
```





在进行任何 Git 操作之前，都要先切换到 Git 仓库目录，也就是先要先切换到项目的文件夹目录下。

![img](https://picx.zhimg.com/v2-28f9d91a77bfe6d3f612a94f2471869b_1440w.jpg)

并在文件夹夹里创建一个文件如 test.txt。

![img](https://pic1.zhimg.com/v2-55de4cc3b93f5a0e3ddf83a0f149fd92_1440w.jpg)

然后我们进入这个文件夹，把它当作我们的 Git 仓库，鼠标右击选择**Git Bash Here**：

![img](https://pica.zhimg.com/v2-7624ff752d5674c0855fda89014f9e1a_1440w.jpg)

打开的 GitBash 可以看到自动定位到我们的仓库位置，如果在其他位置打开，还得再次定位，这就是我们在仓库/文件夹打开的原因。

![img](https://pic4.zhimg.com/v2-5da73cc924cec00c451f02d9735a2477_1440w.jpg)

**2.1 git status**

首先输入 **git status** 命令，查看仓库状态。可以看到提示，当前目录还不是一个 Git 仓库，它还只是个普通的文件夹。

> 以后我们可以经常使用**git status**命令查看仓库状态

![img](https://pic3.zhimg.com/v2-a98089ef883d25faca200c5e43b0ed74_1440w.jpg)

**2.2 git init**

既然它还不是 Git 仓库，我们就用 **git init** 初始化，可以看到初始化完成，它是一个空的 Git 仓库。

![img](https://pica.zhimg.com/v2-f0f7c6870677f2e907eca02b20b7a6c2_1440w.jpg)

我们再用 **git status** 命令检查一下仓库，发现 Git Demo 目录已经成为一个 Git 仓库了，并且默认进入 Git 仓库的 master分支，即主分支。

返回的三条语句：

- 进入 Git 仓库的主分支
- 目前还没有提交过文件（git 是版本控制工具，要记录提交历史的以及每一个版本）
- 这时最主要的是提示 test.txt 文件是 Untracked files ，就是说 test.txt 这个文件还没有被跟踪，还没有提交在 git 仓库里呢，而且提示你可以使用 git add 去操作你想要提交的文件。

![img](https://pic4.zhimg.com/v2-aea30cfecf603975de6e3708277c07bf_1440w.jpg)

**2.3 git add**

接下来我们往仓库里添加文件，根据提示，输入 **git add test.txt**，再输入 **git status** 看看仓库状态：

虽然仓库没有提交历史，但提示以下文件 Changes to be committed ，意思就是 test.txt 文件等待被提交。

![img](https://pic2.zhimg.com/v2-cf3746adab0f0cf347abb38d8d2af5f7_1440w.jpg)

可以看到刚刚 git “不承认” 我们的文件，现在通过 **git add** 命令”承认“了，并且把它加到一个可提交的列表（临时缓冲区）。

当然你可以使用 git rm --cached 这个命令去移除这个缓存。

**2.4 git commit**

刚刚我们把文件送到缓冲区，接下来用 **git commit** 命令将缓存区里的改动给提交到本地的版本库。

每次使用 **git commit** 命令我们都会在本地版本库生成一个 40 位的哈希值，这个哈希值也叫 commit-id（这个版本的编号），commit-id 在版本回退的时候是非常有用的，它相当于一个索引，可以在未来的任何时候通过与 git reset 的组合命令回到这里。

**git commit** 不是单独使用的，一般需要参数指定提交方式，我们可以输入 **git commit --help** 查看使用帮助（会跳转到相关网页）。

输入 **git commit -m ”first commit“**，我们成功将文件 `text.txt` 提交到了 Git 仓库。其中， -m 代表提交信息， first commit 是本次提交的信息，需要写在双引号内，提交信息你也可以写成 my commit，text commit 等任意句子。

![img](https://pic2.zhimg.com/v2-34facfbe0a3aaa3b1291dc48fd38a30d_1440w.jpg)

我们再看看仓库状态，结果显示 nothing to commit, working tree clean，这表示已经没有内容可以提交了，即全部内容已经提交完毕。

![img](https://pic3.zhimg.com/v2-3f1a56e4c0f5c0c4025e49fd44de710a_1440w.jpg)

![img](https://pic3.zhimg.com/v2-3f1a56e4c0f5c0c4025e49fd44de710a_1440w.jpg)



补充**git commit**常见的用法

- **git commit --amend**，追加提交，它可以在不增加一个新的commit-id的情况下将新修改的代码追加到前一次的commit-id中

### 查看作者、时间

**2.5 git log**

输入**git log**命令，打印 Git 仓库提交日志，会显示作者、时间和你写的提交信息：

![img](https://pic3.zhimg.com/v2-ecbd2fce1a60b9c4807a9cf2645eded8_1440w.jpg)

### 创建、切换、合并分支

**2.6 git branch**

branch 是分支的意思，多人合作时，你一个模块，别人一个模块，你们各建一个分支就能保证改动互不干扰，等最后做完，把你们的分支合并起来就好。

git init 会初始化一个主分支 master，也是我们现在的默认分支。

输入 **git branch**，可以看到我们就在 master 这里：

![img](https://pic2.zhimg.com/v2-bb0e635fc0725edb23332c91569f9d45_1440w.jpg)

在此基础上新建分支用**git branch a**，其中 a 是我起的分支名字。然后再用**git branch**查看，可以看到 master 有了一个分支，这时候分支 a 跟分支 master 是一模一样的内容。

![img](https://pic3.zhimg.com/v2-89d43fdb142b64b265024cbcf6ba0a5a_1440w.jpg)

**2.7 git checkout**

虽然建立分支，但可以看到 * 还在 master 上，说明我们还处于主分支。这时候，输入 **git checkout a** 命令，切换到分支 a 上，再次查看分支，* 在 a 前，说明我们处于分支 a ，可以修改分支 a 的内容了。

![img](https://pic1.zhimg.com/v2-54fcb512de84dd36b3755972118c0c36_1440w.jpg)

> 先新建再切换很麻烦，有一步到位的方法**git checkout -b a**，意思是建立分支后自动切换到该分支。

**2.8 git merge**

你在 a 分支完成了你的部分，但是主分支还是原来那样，这个时候就需要把你的代码合并到主分支 master上来，**git merge** 就是合并分支用到的命令。

**注意，合并分两步**

1. 切换到 master 分支，如果你已经在了就不用切换了;
2. 执行 git merge a ，意思就是把 a 分支的代码合并过来，不出意外，
   这个时候 a 分支的代码就顺利合并到 master 分支来了。

![img](https://pic3.zhimg.com/v2-79bc3e8783cd1eb86f7dbc20995a8eae_1440w.jpg)

**删除分支**：有添加就有删除，假如分支新建错了，或者 a 分支的代码已经**顺利合并**到 master 分支来了，那么 a 分支没用了，用 **git branch -d** 删除。删除分支 a 的语句为 git branch -d a，这里不再演示。

**强制删除分支**：上面说了 **git branch -d** 是删除分支的意思，有些时候可能会删除失败，比如如果 a 分支的代码还**没有合并**到master，你执行 git branch -d a 是删除不了的，它会智能的提示你a分支还有未合并的代码，但是如果你非要删除，那就执行 **git branch -D a** 就可以强制删除 a 分支。

**2.9 git tag**

这是今天学习的最后一个操作，tag 是标签的意思，写报告、写代码都有版本的迭代，比如 1.0，1.1 。git 可以给每个版本打上这样的标签，这样你忽然想看看第一版的报告是什么内容，就可以通过标签切换回去。

**git tag v1** 代表新建 v1 tag，你可以换成其他名字，**git tag** 代表查看标签。

![img](https://pic3.zhimg.com/v2-a005764dcc3f9a90419adb0da97846b0_1440w.jpg)

如果想切换至某一标签，还是用刚刚学的**git checkout v1**即可，v1 是某版本标签名。



以上全是一些最基本的 git 操作，而且全是在本地环境进行操作的，完全没有涉及到远程仓库，最后总结一下基本 Git 操作：

- git status ：查看仓库状态
- git init：创建一个空仓库，或者重新初始化一个已有仓库
- git add：把文件添加到可提交列表（临时缓冲区）
- git commit：提交改动（增删改）至仓库
- git log：打印提交日志
- git branch：查看、添加、删除分支
- git checkout：切换分支、标签
- git merge：合并分支
- git tag：新建、查看标签

## 本地与远程仓库的交互


下载仓库文件
把本地的文件上传到远程仓库
仓库文件有更新时保持本地与远程一致

若主分支是 **main**，不是 **master**，下面的语句中要把 master 换成 main。

### 1. SSH 配置

要想本地与远程仓库交互，必须有一个安全机制防止数据泄漏，这个安全机制就是 SSH

**1.1 生成 SSH key**

打开 GitBash ，输入 ssh，如果返回如下界面说明 SSH 已安装![img](https://pic4.zhimg.com/v2-39d5c0b5ee6080e84b14c312d647ce99_1440w.jpg)

紧接着输入 **ssh-keygen -t rsa** （指定 rsa 算法生成密钥），接着连续三个回
车键（不需要输入密码），然后就会生成两个文件 id_rsa （密钥）和 id_rsa.pub （公钥）。

划线的地方就是密钥和公钥存储的位置，等会儿需要打开文件复制密钥。

![img](https://pic1.zhimg.com/v2-61e214b5581ed8a1f5ed0b2c4bfe41a4_1440w.jpg)

**1.2 添加 SSH key**

接下来要做的是把 id_rsa.pub 的内容添加到 GitHub 上，这样你本地的 id_rsa 密钥跟 GitHub 上的 id_rsa.pub 公钥进行配对，授权成功才可以提交代码。

打开 GitHub，点击右上角的你的头像，点击设置 **settings**：

![img](https://pic1.zhimg.com/v2-52a703b129c727cab6908e9024ba2426_1440w.jpg)

点击左侧的**SSH and GPG keys**，点击右上角的**New SSH key**：

![img](https://picx.zhimg.com/v2-ace642dfc3baf9049166364be8f8663f_1440w.jpg)

进入到下面的界面，把刚刚生成的公钥 （id_rsa.pub 文件中的内容）复制到 Key 所在框， Title 不用填，复制好点击下方的 **Add SSH key** 按钮即可。

提示：id_rsa.pub 可以选择用 Notepad++ 或 Sublime 打开~

![img](https://pica.zhimg.com/v2-151e48875d93a41d71f24bc3b1aa7e0c_1440w.jpg)

添加完返回 SSH 页面就会出现你的本地信息，如下图所示：

![img](https://pic2.zhimg.com/v2-8c11b6bf381d49816661a346f1829149_1440w.jpg)

**1.3 验证绑定**

最后我们验证一下本地的 git 与 GitHub 是否绑定成功，在 GitBash 输入**ssh -Tgit@github.com**，如果返回下面的提示，说明绑定成功，不过有时候返回有些慢呀

![img](https://pica.zhimg.com/v2-153f6e613952e22951b3ed3ee24bf45e_1440w.jpg)

到此为止，我们的 SSH 配置完成，接下来我们可以与远程仓库交互啦！

### 2. 通过 Git 下载代码

最简单的就是点击 **Download ZIP**，这是常规的下载操作，下载一个压缩包到指定目录。

另一种方法就是使用 git clone，这种方法又分为 cloning with HTTPS 和 clone with SSH，下面我们分别演示一下。
当然如果你只是为了下载别人的仓库看看，不需要版本控制这些功能，直接download 就行；但如果需要用到版本管理，就得学习 git clone

![img](https://pic1.zhimg.com/v2-95c3610a4ebeaf1e1e90317b37fd73c4_1440w.jpg)

**2.1 git clone with HTTPS**

注意上面写的是 `Clone with HTTPS`，复制下面的 URL：

![img](https://pic1.zhimg.com/v2-adbe78b782bd466740464ce6fa2963d8_1440w.jpg)

打开文件要 clone 的目录，右击选择`Git Bash`Here:

![img](https://pic1.zhimg.com/v2-4d44fb10ff024a8da70b11cde8c2d468_1440w.jpg)

打开 GitBash 窗口，输入 **git clone**，后面跟刚刚复制的 URL，回车，等它下载完，你会看到你的文件夹里直接就是仓库里的文件。

语句长这样：

```text
git clone https://github.com/project/repo.git
```

![img](https://pic4.zhimg.com/v2-8022d93bbb4521a69f6079e1cf4a55cd_1440w.jpg)

**2.2 git clone with SSH**

在 `Clone with HTTPS`那里，点击 `Use SSH `，就会切换成第二种 git clone 方式。

其实步骤与上面的步骤一样，只不过复制的 URL 是 SSH 机制下的：

![img](https://pica.zhimg.com/v2-985b62405745b1b6f30ab7a3e5e3a6f0_1440w.jpg)

接着也是在目录下打开 Git Bash，输入 **git clone**， 后面是刚刚复制的 URL。

语句长这样，可以看出 SSH 下的 URL 和 HTTPs 下的 URL 不一样。

```text
git clone git@github.com:project/repo.git
```

**2.3 异同**

同：HTTPs 与 SSH 下的 git 都可以直接进行 git clone 操作

异：

- HTTPs git clone 到本地，进行了一些文件的修改，当再次提交到 GitHub远程服务器的时候，都会进行账号与密码的输入
- SSH git clone 到本地之后，由于已有 SSH Keys 授权，就不需要用户名和密码进行授权了。

### 3. 通过 Git 提交代码

**3.1 两个命令**

上一节一些基本的 git 命令，不过都只涉及本地，想本地与远程仓库交互，就需要两个新命令：

- git push：翻译为推，当你的代码更新，需要把本地的推到远程仓库
- git pull：翻译为拉，当远程仓库有更新，你需要把远程的拉到本地进行合并

前者在本地无仓库时使用，后者是本地已有仓库时使用。

使用举例：

```text
git push origin master # 把本地代码推到远程 master 分支
git pull origin master # 把远程最新的代码更新到本地
```

**一般我们在 push 之前都会先 pull ，这样不容易冲突。**

**3.2 提交代码**

通过 git 提交代码的前提是**已有仓库**，然后把仓库 clone 到本地，修改后再pull。

相信通过前面的教程，你已经学会如何新建和删除一个仓库，我这里用一个临时建的仓库演示一下提交功能。

![img](https://pic2.zhimg.com/v2-c7956be39ab2bfce559d0f1d8df0a15f_1440w.jpg)

然后我们把仓库 clone 到本地，为了不输入账号密码，我这里用 SSH 的方法 clone。

![img](https://pic2.zhimg.com/v2-437794d64cef130248294e2d34745b0b_1440w.jpg)

如图只是新加一个 python 文件，也可以修改一个已有文件。

![img](https://picx.zhimg.com/v2-00820d40a6ab5bb777d5a5b8174708bf_1440w.jpg)

**所有的改动都需要通过 `git add`和`git commit` 提交到本地的仓库。**

**在这个文件夹/本地仓库**，右击打开 **Git Bash**，输入 `git status` 查看状态，再用 `git add`和`git commit` 提交新加的文件。

注意 first commit 是为这次 commit 起的名字

![img](https://pic3.zhimg.com/v2-a32c17281f81c89c722e506fc7c7a61e_1440w.jpg)

然后输入`git log`查看提交历史，可以看到一共有两次 commit，第一次是在 GitHub 上新建仓库，第二次是刚刚在本地新提交文件。

![img](https://pic3.zhimg.com/v2-60a1f5d595701f54193bbb10eab21476_1440w.jpg)

**到现在我们才完成本地更新的工作，接下来要把本地仓库 push 到远程仓库**

在 **Git Bash** 输入 `git push origin master` 即可

默认向 GitHub 上的 test 目录提交了代码，而这个代码是在 master 分支。

![img](https://pica.zhimg.com/v2-d985d9dffb76e3928b5030bf3e2dea24_1440w.jpg)

然后刷新远程仓库，可以看到有更新：

![img](https://pica.zhimg.com/v2-3d0e3405cd274198783e8a7bc66e5306_1440w.jpg)

####  设置别名

上面提到的基本操作，每次都要输入完整的单词。

使用 `alias`` 可以给这些操作七个简单的别名，就能简化输入。

基本语法是：

```text
git config --global alias.别名 git命令
```

比如给 commit 和 status 起别名：

```text
git config --global alias.ci commit
git config --global alias.st status
```

下次再用这两个命令就可以写：

```text
git ci
git st
```

除了简单命令还可以设置组合命令，给组合命令加引号即可：

```text
git config --global alias.pullom 'pull origin master'
git config --global alias.pushom 'push origin master'
```

![img](https://picx.zhimg.com/v2-46c5524da1b9500760158e582f9f08f3_1440w.jpg)



起别名后，可以看到 `.gitconfig` 文件里有 `alias` 的配置，如果不需要某个别名，删掉即可。

#### 显示改动

 `git add` 会把文件放入暂存区，**在`git add`** **之前**使用 `git diff` 可以显示对某一文件的改动。

红色的部分前面有个 - 代表我删除的，绿色的部分前面有个 + 代表我增加的，所以从这里你们很一目了然的知道我到底对这个文件做了哪些改动。

![img](https://pica.zhimg.com/v2-e805c9acef3c928513e62305482d3054_1440w.jpg)

除此之外，`git diff`还有其他用法：

```text
git diff <$id1> <$id2> # 比较两次提交之间的差异
git diff <branch1>..<branch2> # 在两个分支之间比较
git diff --staged # 比较暂存区和版本库差异
```

#### 版本回滚

想要回滚到某次提交之前的场景

首先用 `git log `查看版本号（commit 后面那一长串就是版本号）：

![img](https://pic1.zhimg.com/v2-cfe8a833f87e0cc8ce010b91086abc0c_1440w.jpg)

找到想要回退的版本号之后，在本地 Git 仓库执行如下命令，选一个即可：

```text
git reset --hard 版本号 (抛弃当前工作区的修改)
git reset --soft 版本号 (回退到之前的版本，但保留当前工作区的修改，可以重新提交)
```

如果你的文件只在本地，执行上面的步骤就可以啦。

如果你的文件提交到 GitHub 上，还需要执行如下命令同步到远端：

```text
git push origin 分支名
```

本地的版本落后于远端的版本，因此我们还需要在上述命令中加上`--force`参数：

```text
git push origin 分支名 --force
```

## GitHub 提交文件夹快速版

首先进入这个文件夹，右击点击 Git Bash Here。

![img](https://pica.zhimg.com/v2-11c67cca339ab57d3afe547c86d1896a_1440w.jpg)

依次输入下面的语句：

```text
git init   // 初始化版本库

git add .   // 添加文件到版本库（只是添加到缓存区），.代表添加文件夹下所有文件 

git commit -m "first commit" // 把添加的文件提交到版本库，并填写提交备注

git remote add origin 你的远程库地址  // 把本地库与远程库关联

git push -u origin main    // 第一次推送时

git push origin main  // 第一次推送后，直接使用该命令即可推送修改
```
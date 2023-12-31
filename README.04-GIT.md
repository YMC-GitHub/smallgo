## basic do

- [x] set git default branch
```bash
git config --global init.defaultBranch main
```

- [x] ini git repo
```bash
[ ! -e .git ] && git init;
```

- [x] set user name and email
```bash
git config --global user.name "yemiancheng" ;git config --global user.email "ymc.github@gmail.com";
```

- [ ] change the master bransh to "main" from "master"
```bash
git branch -M main
```

- [x] add github repo (in plan)
```bash
# copy token files to current project
cp /d/code/shell/github-cli/secrets/* ./secrets

shfile=/d/code/shell/github-cli/index.sh
# repourl=ymc-github/detect-os;repodesc="detect operating system in shell script";

repourl=ymc-github/smallgo;repodesc="tiny programs in go";

# add repo with default desc
$shfile $repourl add
# put repo desc
# $shfile $repourl put $repodesc
# $shfile $repourl del
```

- [x] put github repo desc (process)
```bash
$shfile $repourl put $repodesc
# $shfile $repourl del
```

- [ ] del github repo
```bash
# $shfile $repourl del
```



- [x] add remote repo url
```bash
# git remote add origin https://github.com/YMC-GitHub/smallgo.git
git remote add origin git@github.com:YMC-GitHub/smallgo.git
```

- [x] set gitignore and commit it
```bash
git add .gitignore;git commit -m "build(core): add gitignore";
git add .gitignore;git commit -m "build(core): set note for go allow list";
git add .gitignore;git commit -m "build(core): put gitignore to design project dirs";
```

[fefer: some gitignore template](https://github.com/github/gitignore)


- [x] set edifotconfig and commit it
```bash
git add .editorconfig ; git commit -m "build(core): add editor config";
```

- [x] set license and commit it
```bash
git add LICENSE ; git commit -m "build(core): add license";
```

- [x] set readmd and commit it
```bash
git add README.md ; git commit -m "docs(core): add readme";
```

- [x] get time with format
```bash
format="+%Y-%m-%d %H:%M:%S";time="2021-05-18 13:12:19";date -d "$time" "$format";
```

- [x] git commit files with custom date
```bash
git commit -m "chore(core): update readme in github action" --date "$(date "+%Y-%m-%d %H:%M:%S" -d "+8 hour") +0800"
```

- [x] design project dirs
```bash
git add release_win/.gitignore;git commit -m "build(core): add gitignore to design project dirs";
git add release_win/README.md;git commit -m "docs(core): output *.exe files here";
git add release_win/*.ico;git commit -m "build(core): output icon files here";


git add icons/.gitignore;git commit -m "build(core): add gitignore to design project dirs";
git add icons/README.md;git commit -m "docs(core): put icon files here as resources";
git add icons/*.ico;git commit -m "build(core): add resources files icon";

git add .vscode/.gitignore;git commit -m "build(core): add gitignore to design project dirs";
git add .vscode/README.md;git commit -m "docs(core): put vscode editor config files of project level here";
```

- [x] add docs about how to setup this
```bash
git add README.01-SETUP.md ; git commit -m "docs(core): add readme to setup this";
git add README.01-SETUP.md ; git commit -m "docs(core): put readme to setup this";
git add README.01-SETUP-02.md ; git commit -m "docs(core): put readme to setup this";
git add README.01-SETUP*.md ; git commit -m "docs(core): put readme to setup this";
```

- [x] add docs about how to code before coding
```bash
git add README.02-BEFORE-CODE.md ; git commit -m "docs(core): add readme to code this";
```

- [x] add docs about some libs i noted
```bash
git add README.03-LIB-NOTE.md ; git commit -m "docs(core): add readme to note libs";
git add README.03-LIB-NOTE.md ; git commit -m "docs(core): put readme to note libs";
```

- [x] add go code snippets in project level
```bash
git add .vscode/.gitignore;git commit -m "build(core): add go code snippets in project level";
git add .vscode/go.code-snippets;git commit -m "docs(core): add go code snippets";
git add .vscode/go.code-snippets;git commit -m "build(core): put go code snippets to build for different os";
```

- [x] add hello world in root dirs
```bash
git add main.go;git commit -m "build(core): add hello world in root dirs";
```

- [x] add go mod and workspace
```bash
git add go.mod go.sum go.work.sum;git commit -m "build(core): add mod and workspace";
git add go.mod go.sum go.work.sum;git commit -m "build(core): put mod and workspace";
```

- [x] add mod chapter01
```bash
git add src/chapter01/higo;git commit -m "build(core): add mod chapter01";
```

- [x] push to github repo (the first)
```bash
git push -u origin main
```

- [ ] replace git protocol with https when failed
```bash
# set
git config --global url."https://".insteadof git://
# unset
# git config --global --unset url."https://".insteadof 
```

- [x] rename github repo remote name from "origin" to "github"
```bash
git remote rename origin github
```

- [ ] remove github repo remote
```bash
git remote remove github
```

- [ ] set remote repo url from https to git
```bash
git remote set-url origin git@github.com:YMC-GitHub/smallgo.git
# git remote set-url github git@github.com:YMC-GitHub/smallgo.git
```


- [x] speedup github for `git clone`
```bash
# set
git config --global url."https://ghproxy.com/https://github.com".insteadOf "https://github.com"
# unset
# git config --global --unset url."https://ghproxy.com/https://github.com".insteadOf 
# get
# git config --global url."https://ghproxy.com/https://github.com".insteadOf
```

- [x] unset speedup github for `git clone` when pushing
```bash
# unset
git config --global --unset url."https://ghproxy.com/https://github.com".insteadOf 
```


- [x] mgnt git credential
```powershell
# git config --global credential.helper manager
# git config --global credential.helper cache
# git config --global --unset credential.helper
# https://zhuanlan.zhihu.com/p/157751660
```

- [x] add docs about using git for this repo i noted
```bash
git add README.04-GIT.md ; git commit -m "docs(core): add readme to note git for this repo";
git add README.04-GIT.md ; git commit -m "docs(core): put readme to note git for this repo";

```


- [x] add docs about using ssh in github i noted
```bash
git add README.05-SSH.md ; git commit -m "docs(core): add readme to note ssh";
```

- [x] set pull action for this repo
```bash
# zero:knowledge:s:pull-action
# you must know .
# git config pull.rebase false  # merge
git config pull.rebase true   # rebase
# git config pull.ff only       # fast-forward only
# zero:knowledge:e:pull-action
```

- [x] keep workspace clean before pulling
```bash
# git stash
```

- [x] pull remote github repo before pushing again
```bash
git pull origin main
# git pull github main
```

- [x] push remote github repo (no the first time)
```bash
git push origin main
```

- [x] add chapter02
```bash
git add src/chapter02/higo/*.go; git commit -m "feat(core): use fyne to make higo gui";
git add src/chapter02/higo/readme.md; git commit -m "docs(core): use fyne to make higo gui";
git add go.mod go.sum go.work.sum;git commit -m "build(core): put mod and workspace";
git add src/chapter02/higo/*.go; git commit -m "feat(core): export some func";
```
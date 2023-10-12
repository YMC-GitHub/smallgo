##  install go  
- [x] install go in window with scoop (please with admin right)
```powershell
# scoop search go;
scoop install go;

# GOROOT
# scoop prefix go
```

- [x] make proxy in china
```powershell
go env -w GOPROXY=https://goproxy.cn,direct
```
[refer: use goproxy.io](https://goproxy.io/zh/)

- [x] enable the go modules feature
```powershell
go env -w GO111MODULE=on
```

- [ ] set go path in window for powershell
```powershell
# GOROOT VS GOPATH
# $env:GOPATH = "D:\code\go\smallgo.zero.com"
# or:
# go env -w GOPATH="D:\code\go\smallgo.zero.com"

# or:
# M:\zero\scripts\idman\edit-env.ps1 set -envkey "GOPATH" -value "C:\Users\zero\go"

# go env -w GOPATH="D:\code\go"
# M:\zero\scripts\idman\edit-env.ps1 set -envkey "GOPATH" -value "D:\code\go"

# go env
```

- [ ] set go path in window for bash
```powershell
# export GOPATH=/d/code/go/smallgo.zero.com
```

- [x] manage go modules and deps with go mod
[refer:1 manage go modules and deps](https://zhuanlan.zhihu.com/p/482014524)

- [x] define go project dirs construtor
```powershell
# src -- save source code,save files that `go GET`
# pkg -- save compiled files,save files that `go install`
# bin -- save compiled files that can be exec file  ,save files that `go install`
```


## code go in vscode editor
- [x] install or update vscode in window with scoop
```powershell
# scoop install vscode;
```

- [x] update vscode in window with scoop
```powershell
# scoop upgrade vscode;
```

- [x] install go exts in vscode
```powershell
code --install-extension  golang.go
```

- [x] install go dev tool kit in vscode ?
```powershell
# in vscode
# ctrl + shift +p
# go: install/update
```

- [x] set auto saving go code in vscode

- [x] define or use go code snippets in vscode
```powershell
# command + shift + p
# config user snippet
# sletect one level: project,global,language
# .vscode\go.code-snippets

# or:
# write snippet at https://snippet-generator.app/

# pkg_import
# pkg_import_one
# pkg_import_many
# pkg_define
```

- [x] code go code in vscode


- [x] design and add micro code
```powershell
# src/chapter01/higo

$name="chapter01/higo" ; $null=mkdir -p "src/$name" -ErrorAction SilentlyContinue ; 
New-item -Path "src/$name" -Name "readme.md" -ItemType "file" -ErrorAction SilentlyContinue ;
```

## manage deps
- [x] where the deps that `go get` save
```powershell
# when GO111MODULE="on" or "off"
# GOPATHE/src

# when GO111MODULE="on"
# GOPATHE/pkg/mod/
```
[refer:1 where download deps to](https://www.cnblogs.com/davis12/p/14987615.htm)

- [x] main flow
```powershell
# init,tidy,download,vendor,edit,grapth,verifi

# go help mod 
go mod init pkgname
go mod tidy
# go help mod download
go mod download
# go help mod vendor
go mod vendor
go mod edit
# go help mod graph
go mod graph
# go help mod verify

go mod verify
```

## build source code to output

- [x] install pkg
```powershell
# bin,pkg,src
go install chapter11/goinstall
```

- [x] build source files to binay files
```powershell
go build
# go help build
# usage: go build [-o output] [build flags] [packages]

# use package name if it exsits
# go build
# ./zerosmgo.exe
# 
# go build -o main.exe  main.go
# ./main.exe

# to mini files:
# go build -ldflags="-s -w -H windowsgui" main.go
```

- [x] build pakcage with package name in subdirs with gid mod (uc 1)
```powershell
# uc: ini a subdir as a package
$workspace=get-location;cd src\chapter01\higo;go mod init chapter01;go mod tidy;cd $workspace;

# uc: clean all *.exe files for package xx in a subdir
# del ./src/chapter01/higo/*.exe -recurse
# uc: clean all *.exe files for package xx in project root dirs
# del ./*.exe -recurse

# uc: build pakcage with package name in subdirs
# use package name if it exsits
go build chapter01
# ./chapter01.exe


# uc: use the main.exe as output
# go clean -x -n chapter01
# go build -o main.exe chapter01
# ./main.exe

# uc: switch location then build
# $workspace=get-location;cd src\chapter01\higo;
# go build
# ./chapter01.exe

# go build -o main.exe
# ./main.exe

# go build -o mainmin.exe  -ldflags="-s -w -H windowsgui"
# ./mainmin.exe
```

- [x] build pakcage with package name in subdirs with gid mod (uc 2)
```powershell
cd $workspace;
go build -C "src/chapter01/higo" -o main.exe 
src/chapter01/higo/main.exe

# go build -C "src/chapter01/higo" -o mainmin.exe -ldflags="-s -w -H windowsgui"
go build -C "src/chapter01/higo" -o mainmin.exe -ldflags="-s -w"
# -H windowsgui -- without dos window
# -s -- Remove sign information
# -W -- Remove DWARF debug information
src/chapter01/higo/mainmin.exe
```

- [x] build pakcage with package name in subdirs with gid mod (uc 3)
```powershell
cd $workspace;
go build -o release_win/main.exe chapter01
release_win/main.exe


go build -o release_win/mainmin.exe -ldflags="-s -w" chapter01
release_win/mainmin.exe

# go build -o release_win/mainmincli.exe -ldflags="-s -w" chapter01
# release_win/mainmincli.exe


# go build -o release_win/maingui.exe -ldflags="-s -w -H windowsgui" chapter01 
# release_win/maingui.exe # no thing output in console window
```

- [x] build for different platform (when CGO_ENABLED=0 )
```powershell
go env GOOS GOARCH
# for win 64
go env -w GOOS=windows;go env -w GOARCH=amd64;

# for win 32
go env -w GOOS=windows;go env -w GOARCH=386;
 

# for mac (Apple Silicon)
go env -w GOOS=darwin;go env -w GOARCH=amd64;

# for linux 64
go env -w GOOS=linux;go env -w GOARCH=arm64;go env -w GOARM=7;

# then:
# go build -o xxx *.go
```
[refer:1 go cross-compiling ](https://juejin.cn/post/6844903817071296525#heading-13)

- [x] build for different platform (when CGO_ENABLED=1 )
```powershell
go env GOOS GOARCH
# for win 64
go env -w GOOS=windows;go env -w GOARCH=amd64;
# TODO ...

# for win 32
go env -w GOOS=windows;go env -w GOARCH=386;
# TODO ...

# for mac (Apple Silicon)
go env -w GOOS=darwin;go env -w GOARCH=amd64;
# TODO ...

# for linux 64
go env -w GOOS=linux;go env -w GOARCH=arm64;go env -w GOARM=7;
go env -w CC=aarch64-linux-gnu-gcc;go env -w CXX=aarch64-linux-gnu-g++;go env -w AR=aarch64-linux-gnu-ar;


```


- [x] build source files to binay files and run
```powershell
go run
# usage: go run [build flags] [-exec xprog] package [arguments...]
```

- [x] run binay files
```powershell
./zerosmgo.exe
```

- [x] clean build/install/test files
```powershell
go clean

go clean -i -n
# go help clean
# usage: go clean [clean flags] [build flags] [packages]
# -n print the remove commands it would execute
# -x print remove commands as it executes them
```



- [x] get remote pkg
```powershell
# cellnet is not an exec pkg, it is a net lib
go get github.com/davyxu/cellnet

# tabtoy is  an exec pkg
go get github.com/davyxu/tabtoy
```

- [x] use go module or use gopath ?
```powershell
# USE GO MOD AFTER 1.13 (RECOMMEND)
go env -w GOPATH="D:\code\go"
go env -w GO111MODULE=on;go mod init zerosmgo;go mod tidy;
# go mod install zerosmgo;  #?
# go install zerosmgo; #?
# go install # 
# go install .
# go install chapter01/higo # package chapter01/higo is not in std (D:\CustomSoftware\windows\Scoop\apps\go\current\src\chapter01\higo)
# use package name if it exsits
# go build
# ./zerosmgo.exe
# 
# go build -o main.exe  main.go
# ./main.exe


# go help mod
# init        initialize new module in current directory
# tidy        add missing and remove unused modules
# download    download modules to local cache
# edit        edit go.mod from tools or scripts
# graph       print module requirement graph
# why         explain why packages or modules are needed
# verify      verify dependencies have expected content



# go env -w GO111MODULE=on
# go mod init test.com
# code ...
# go mod tidy
# go build;go test;
# go list -m all


# go list -m -versions github.com/gogf/gf
# go get github.com/gogf/gf@master
# go get github.com/gogf/gf@v1.16.9
```

```powershell
# SET GOPATH TO PROJECT DIR ?
go env -w GO111MODULE=off;go env -w GOPATH="D:\code\go\smallgo.zero.com"

# ALL THE PKS SAVE IN GOPATH/SRC/;PKGS THE DIFFERENT PROJECTS USED IN A GIT REPO!
# PROJECT A AND PROJECT B USE PKG A AND B;
```


[refer:1 about  go mod and use gopath and choose one](https://blog.csdn.net/qq_15386973/article/details/107172344)


[refer:2 manage dep from gopath to go mod](https://www.cnblogs.com/hiyang/p/12695155.html)

[refer:3 go module manage](https://juejin.cn/post/6947666847147917320?from=search-suggest)



## make xx.exe files 
- [x] mini my go exe files (uc 1)
```powershell
# go build -o release_win/mainmincli.exe -ldflags="-s -w" chapter01
# release_win/mainmincli.exe


# go build -o release_win/maingui.exe -ldflags="-s -w -H windowsgui" chapter01 
# release_win/maingui.exe # no thing output in console window
```
[refer:1](https://blog.csdn.net/qq_32394351/article/details/93468119)


- [x] prepare xx.exe icon
```powershell
# magick --version
$null=mkdir -p icons;
$icon_burn_src="G:\code\cdn\data\2022-11-27";magick convert -density 256x256 -background transparent $icon_burn_src/favicon-256X256.png -define icon:auto-resize=256,48,32,16 -colors 256 icons/window.ico

# $icon_burn_src="M:\LORA-TRAIN-FACE\2023-09-07\xiaoai_3_preview";magick convert -density 256x256 -background transparent $icon_burn_src/hyx_0.8_00108_.png -define icon:auto-resize=256,48,32,16 -colors 256 icons/window.ico

```

- [x] prepare rcedit cli in window with scoop
```powershell
# scoop search rcedit;
scoop install rcedit;
# scoop update rcedit;
```


- [x] edit xx.exe icon (uc rcedit)
```powershell
copy-item icons/window.ico release_win/logo.ico
# may fails
rcedit release_win/mainmin.exe --set-icon "./release_win/logo.ico"
```

- [x] edit xx.exe infomation (uc rcedit)
```powershell
# may fails
$version="1.0.0";$ProductName="smgo";$ProductDesc="small promgram";$CompanyName="zero.com";
rcedit release_win/mainmin.exe  --set-version-string "ProductVersion" "$version"
rcedit release_win/mainmin.exe  --set-version-string "ProductName" "$ProductName"
rcedit release_win/mainmin.exe  --set-version-string "FileDescription" "$ProductDesc"
rcedit release_win/mainmin.exe  --set-version-string "CompanyName" "$CompanyName"
rcedit release_win/mainmin.exe  --set-version-string "LegalCopyright" "Copyright (c) 2022-2023 $CompanyName"
rcedit release_win/mainmin.exe  --set-file-version "$version"
rcedit release_win/mainmin.exe  --set-product-version "$version"
rcedit release_win/mainmin.exe  --set-product-version "$version"
rcedit release_win/mainmin.exe --set-version-string OriginalFilename "$ProductName"
# or:
# $file="D:\ts-project\scripts\set-exe-info.ts";tsx "$file";
```
[refer:1 edit xx.exe or xx.dll file infomation](https://www.cnblogs.com/webjlwang/p/16461209.html)


## run xx.exe files in production
- [x] in cli terminal powershell
```powershell
release_win/mainmin.exe
```

- [x] double click program in window


## elegant shutdown
- [x] elegant shutdown
[refer:1 elegant shutdown program](https://www.cnblogs.com/wishFreedom/p/16573575.html)


## fix fags for newer when setting up
- [x] fix `go.mod file not found in current directory or any parent directory`
```powershell
# in project dir
# go mod init
# vs go mod init xx
# ini modules  in your main.go files dir
# go mod init youpkgname

go env -w GO111MODULE=on;go mod init zerosmgo;go mod tidy;

# 
# or :
# go env -w GO111MODULE=off ?

```

[refers:1](https://www.jianshu.com/p/19c11ba25a15)

[refers:2.go mod init syantc](https://blog.csdn.net/mmiaoChong/article/details/121133265)

[refers:3.about your project name and pkg name](https://www.zhihu.com/question/458797990)

- [x] fix `$GOPATH/go.mod exists but should not`
```powershell
# your code project path vs GOPATH
# del os env GOPATH
M:\zero\scripts\idman\edit-env.ps1 set -envkey "GOPATH" -value "D:\code\go"
# set GOPATH in go env
go env -w GOPATH="D:\code\go"
```
[refer:1](https://blog.csdn.net/AlexanderRon/article/details/120656720)

[refer:2 del project path in gopath after mod on ?](https://www.jianshu.com/p/fdebd9af571c)


- [x] fix `package chapter01/higo is not in std`
```powershell
# when: go build -o main chapter01/higo

# done:build without args
# cd src\chapter01\higo;go mod init chapter01;go mod tidy;go build;

# ./chapter01.exe

# done:build + files
# cd src\chapter01\higo;go mod init chapter01;go mod tidy;go build main.go lib.go;
# ./main.exe

# done:build + exec + files
#  cd src\chapter01\higo;go mod init chapter01;go mod tidy;go build -o higo.exe  main.go lib.go
# ./higo.exe

# done:build + pakcage
# when using gopath
# go env -w GOPATH="D:\code\go\smallgo.zero.com"
# go build -o main chapter01/higo


# go build -o main.exe src/chapter01/higo
```

- [x] fix `warning: go env -w GOPATH=... does not override conflicting OS environment variable`
```powershell
$env:GOPATH=""
```
[refer:1](https://www.cnblogs.com/wangchengxu/p/15892993.html)

- [x] fix `no Go files in xx dir`
```powershell
# add xx.go to that dirs
```


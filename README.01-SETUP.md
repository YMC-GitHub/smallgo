##  install go  
- [x] install go in window with scoop (please with admin right)
```powershell
# scoop search go;
scoop install go;

# GOROOT
# scoop prefix go
```

- [x] check go version
```bash
go version
# go version go1.21.3 windows/amd64
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

- [x] get/set the important config
```powershell
# GET
go env GOOS GOARCH CGO_ENABLED CC GOROOT GOPATH

# SET
go env GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=gcc


# UNSET
go env -u GOOS GOARCH CGO_ENABLED CC GOROOT GOPATH


# build on linux ? 
# go env -w GOOS=linux;go env -w GOARCH=arm64;go env -w CGO_ENABLED=1;

# go env -w GOOS=linux;go env -w GOARCH=arm64;go env -w CGO_ENABLED=1;go env -w CC=arch64-linux-gnu-gcc;
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
# unset them:
# go env -u GOOS GOARCH CGO_ENABLED CC GOROOT GOPATH

# get them:
# go env GOOS GOARCH CGO_ENABLED CC GOROOT GOPATH

go build .
# ./zerosmgo.exe

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
# https://developer.fyne.io/started/cross-compiling
go env GOOS GOARCH
# for win 64
go env -w GOOS=windows GOARCH=amd64;

# for win 32
go env -w GOOS=windows GOARCH=386;

# for mac (Apple Silicon)
go env -w GOOS=darwin GOARCH=amd64;
go env -w GOOS=darwin GOARCH=386;


# for linux
go env -w GOOS=linux GOARCH=amd64;
go env -w GOOS=linux GOARCH=386;

go env -w GOOS=linux GOARCH=arm64;
go env -w GOOS=linux GOARCH=arm;

go env -w GOOS=linux;go env -w GOARCH=arm64;go env -w GOARM=7;


# for android
go env -w GOOS=android GOARCH=amd64;
go env -w GOOS=android GOARCH=386;

go env -w GOOS=android GOARCH=arm64;
go env -w GOOS=android GOARCH=arm;

# then:
# go build -o xxx *.go


# go install github.com/fyne-io/fyne-cross@latest
# fyne-cross <command> [options]
# fyne-cross windows -arch=amd64,386
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
# go env -w GOPATH="D:\code\go"
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

# list all mod
# go list -m all


# go list -m -versions github.com/gogf/gf
# go get github.com/gogf/gf@master
# go get github.com/gogf/gf@v1.16.9
```

```powershell
# SET GOPATH TO PROJECT DIR ?
go env -w GO111MODULE=off;go env -w GOPATH="D:\code\go\smallgo.zero.com"

# GO ENV UNSET GOPATH
# go env -u GOPATH

# ALL THE PKS SAVE IN GOPATH/SRC/;PKGS THE DIFFERENT PROJECTS USED IN A GIT REPO!
# PROJECT A AND PROJECT B USE PKG A AND B;
```


[refer:1 about  go mod and use gopath and choose one](https://blog.csdn.net/qq_15386973/article/details/107172344)


[refer:2 manage dep from gopath to go mod](https://www.cnblogs.com/hiyang/p/12695155.html)

[refer:3 go module manage](https://juejin.cn/post/6947666847147917320?from=search-suggest)



## make xx.exe files 
- [x] mini my go exe files (uc 1)
```powershell
# build
go build -o release_win/mainmincli.exe -ldflags="-s -w" chapter01
# run
release_win/mainmincli.exe

# build
go build -o release_win/maingui.exe -ldflags="-s -w -H windowsgui" chapter01 
# run
release_win/maingui.exe # no thing output in console window
```
[refer:1](https://blog.csdn.net/qq_32394351/article/details/93468119)


- [x] prepare xx.exe icon (use imagemagick)
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
# copy to release dir
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
# solv: add your path (ps: ./src/chapter02/higo ) to go.work in project root dir   
# why: 1. *.go files in your project root mod do not use any code from ./src/chapter02/higo


# after that:
# done:build without args
# cd src\chapter01\higo;go mod init chapter01;go mod tidy;go build;./chapter01.exe
# $ws=get-location;cd src\chapter02\higo;go mod init chapter02;go mod tidy;go build;./chapter02.exe;set-location $ws;
# main module (zerosmgo) does not contain package zerosmgo/src/chapter02/higo

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


## use fyne to make gui
- [x] download fyne lib
```powershell
go get fyne.io/fyne/v2@latest
go install fyne.io/fyne/v2/cmd/fyne@latest
go get -u github.com/fyne-io/fyne-cross

# download the require pkg
go mod tidy
```

- [x] code your first app for fyne
```go
// ...
```

- [x] run *.go files for dev
```powershell
go run main.go
go run chapter02
```

- [x] build *.go files to *.exe files
```powershell
# go build -o release_win/fyne.exe .
# release_win/fyne.exe


go build -o release_win/main.exe .
release_win/main.exe

go build -o release_win/mainmin.exe -ldflags="-s -w" .
release_win/mainmin.exe

# go build -o release_win/mainmincli.exe -ldflags="-s -w" .
# release_win/mainmincli.exe


# go build -o release_win/maingui.exe -ldflags="-s -w -H windowsgui" . 
# release_win/maingui.exe # no thing output in console window
```

- [x] build *.exe files for dev
```powershell
# go mod tidy
go build -o release_win/fyne.exe .;release_win/fyne.exe;
```

- [x] fix `error while importing fyne.io/fyne/v2/app: build constraints exclude all Go files in xx\gl`
```go
// error while importing fyne.io/fyne/v2/app: build constraints exclude all Go files in D:\code\go\pkg\mod\github.com\go-gl\gl@v0.0.0-20211210172815-726fda9656d6\v3.2-core\gl
// https://blog.csdn.net/booskz/article/details/126941067
// go get -u github.com/fyne-io/fyne-cross
// fyne-cross windows -arch=amd64,386

// go env GOOS GOARCH CGO_ENABLED GOPATH

// enable CGO_ENABLED
go env -w GOOS=windows;go env -w GOARCH=amd64;go env -w CGO_ENABLED=1;

go build .
```

- [x] fix `cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%` 
```powershell
# choose one of them ? msys2,tdm-gcc,Cygwin,gcc
# scoop search gcc;
# scoop install gcc;
# 11.2.0,96MB

# scoop install tdm-gcc;
# 10.3.0,76MB

#  scoop search msys2
scoop install msys2
# 2023-0718,55MB
```

- [x] Once installed do not use the MSYS terminal that opens
- [x] Open “MSYS2 MinGW 64-bit” from the start menu
- [x] Execute the following commands (if asked for install options be sure to choose “all”)
```bash
# in msys2

# check prepare:
gcc --version;g++ --version;go version;

# copy:ctrl + c; 
# paste to msys2 :shift + ins;

pacman -Syu
pacman -S git mingw-w64-x86_64-toolchain

# pacman -Syu
# pacman -S git mingw-w64-x86_64-toolchain
# pacman -S mingw-w64-ucrt-x86_64-gcc
# gcc --version

```
- [x]  add /c/Program\ Files/Go/bin and ~/Go/bin to your $PATH, for MSYS2
```bash
# add your GOPATH/bin to OS PATH

# eg.
# echo "export PATH=\$PATH:/c/Program\ Files/Go/bin:~/Go/bin" >> ~/.bashrc


# add
# use go installed by  scoop
# echo "export PATH=/d/CustomSoftware/windows/Scoop/apps/go/current/bin:\$PATH" >> ~/.bashrc
# use go path that has set as GOPATH
echo "export PATH=/d/code/go/bin:\$PATH" >> ~/.bashrc

source ~/.bashrc
go --version

# del
# sed -i /Scoop/d ~/.bashrc
# sed -i /code/d ~/.bashrc

# get
# cat ~/.bashrc

# use
# source ~/.bashrc

```

```powershell
# M:\zero\scripts\idman\edit-env-path.ps1 get
# $p=scoop prefix go;$p=$p -replace "\\","/" -replace ":","";$p="/$p";$p

```


- [x] For the compiler to work on other terminals you will need to set up the windows %PATH% variable to find these tools. 
```powershell
# “C:\msys64\mingw64\bin”
```

- [x] build
```bash
cd /d/code/go/smallgo.zero.com

go env GOOS GOARCH CGO_ENABLED CC GOPATH GOROOT


# GO ENV UNSET GOROOT
# go env -u GOROOT

# GO ENV UNSET GOPATH
# go env -u GOPATH

# step CC,CGO,GOARCH,GOOS in bash
CC=gcc;CGO_ENABLED=1;GOARCH=amd64;GOOS=windows
# step CC,CGO,GOARCH,GOOS in with go env
go env -w GOOS=linux;go env -w GOARCH=arm64;go env -w CGO_ENABLED=1


go build .
```

[refer:1 Install one of the available C compilers for windows for fyne](https://developer.fyne.io/started/)

[refer:2 Build Fyne in Windos 10 Pro ](https://cloud.tencent.com/developer/article/2329625?areaId=106001)

[refer:3 Windows builds MSYS2 and MINGW64 environments](https://www.jianshu.com/p/96031565dafb)

- [ ] fix `gcc_arm64.S:30: Error: no such instruction:` 
```bash
# gcc_arm64.S:30: Error: no such instruction: 

# CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC=aarch64-linux-gnu-gcc go build -o baseline_client-arm64  baseline_client.go

# go env GOOS GOARCH CGO_ENABLED CC GOROOT GOPATH


# build on linux ? 
# go env -w GOOS=linux;go env -w GOARCH=arm64;go env -w CGO_ENABLED=1;

# go env -w GOOS=linux;go env -w GOARCH=arm64;go env -w CGO_ENABLED=1;go env -w CC=arch64-linux-gnu-gcc;
# expect:done
# but: cgo: C compiler "aarch64-linux-gnu-gcc" not found: exec: "aarch64-linux-gnu-gcc": executable file not found in %PATH%

# build on window ? 
# go env -w GOOS=windows;go env -w GOARCH=amd64;go env -w CGO_ENABLED=1;go env -w CC=gcc;

```

- [x] fix `fyne: command not found when you have installed in it`
```powershell
#  
# https://developer.fyne.io/faq/troubleshoot#compiler-issues
# go env GOPATH

$value=go env GOPATH;$value="$value\bin";
M:\zero\scripts\idman\edit-env-path.ps1 get
M:\zero\scripts\idman\edit-env-path.ps1 del -value $value
M:\zero\scripts\idman\edit-env-path.ps1 add -value $value
```


- [x] MSYS2 uses windows environment variables
```powershell
M:\zero\scripts\idman\edit-env.ps1 set -envkey "MSYS2_PATH_TYPE" -value "inherit"

# refer: https://blog.csdn.net/a2824256/article/details/118599404
```


- [x] restore msys2 default os path
```bash
export PATH=/mingw64/bin:/usr/local/bin:/usr/bin:/bin:/c/Windows/System32:/c/Windows:/c/Windows/System32/Wbem:/c/Windows/System32/WindowsPowerShell/v1.0/:/usr/bin/site_perl:/usr/bin/vendor_perl:/usr/bin/core_perl
```

- [x] check your fyne installation
```bash
# https://geoffrey-artefacts.fynelabs.com/github/andydotxyz/fyne-io/setup/latest/
```

- [x] fix `could not fine gcc or clang compilers`
```powershell
# when check your fyne installation

# D:\CustomSoftware\windows\Scoop\apps\msys2\current\mingw64\bin
# [fyne & msy2 - add mingw64/bin to OS PATH](https://blog.csdn.net/sywdebug/article/details/126459784)

# get which go is 
# gcm go


# GOTOOLDIR=D:\CustomSoftware\windows\Scoop\apps\go\current\pkg\tool\windows_amd64

$value="D:\CustomSoftware\windows\Scoop\apps\msys2\current\mingw64\bin"
# M:\zero\scripts\idman\edit-env-path.ps1 get
# M:\zero\scripts\idman\edit-env-path.ps1 del -value $value
M:\zero\scripts\idman\edit-env-path.ps1 add -value $value

# D:\CustomSoftware\windows\Scoop\apps\msys2\current\ucrt64\bin

# from powershell to mysy2
# ming264
```

- [x] fix `fyne ` not found in powershell
```bash
# get where fyne is in bash
which fyne
# /d/code/go/bin/fyne
```

```powershell
$value="D:\code\go\bin";M:\zero\scripts\idman\edit-env-path.ps1 add -value $value;

# M:\zero\scripts\idman\edit-env-path.ps1 get
# M:\zero\scripts\idman\edit-env-path.ps1 del -value $value
M:\zero\scripts\idman\edit-env-path.ps1 add -value $value
```

- [x] the size of gui *.exe file of hello world with fyne is `30+ mb` .

- [x] design and add micro code
```powershell
# add dirs
$name="chapter02/higo" ; $null=mkdir -p "src/$name" -ErrorAction SilentlyContinue ; 
# add readme.md file
$null=New-item -Path "src/$name" -Name "readme.md" -ItemType "file" -ErrorAction SilentlyContinue ;

# add code files
$null=copy-item main.go "src/$name/main.go"

# add a package , as a module, and download require deps
$workspace=get-location;cd "src/$name";$pkgName=$name -replace "/.*","";go mod init $pkgName;go mod tidy;cd $workspace;


# uc: clean all *.exe files for package xx in a subdir
del "src/$name/*.exe" -recurse

# uc: clean all *.exe files for package xx in project root dirs
del ./*.exe -recurse

# uc: build pakcage with package name in subdirs
$workspace=get-location;cd "src/$name";go build $pkgName;& ./"$pkgName".exe;cd $workspace;
# $workspace=pwd;cd "src/$name";go build;& ./"$pkgName".exe;cd $workspace;


# uc: use the main.exe as output
$workspace=pwd;cd "src/$name";go build -o main.exe $pkgName;./main.exe;cd $workspace;

# go build -o mainmin.exe  -ldflags="-s -w -H windowsgui"
# ./mainmin.exe

# uc:
go build -C "src/$name" -o main.exe ;& "src/$name/main.exe";

# uc:
go build -C "src/$name" -o mainmin.exe  -ldflags="-s -w" ;& "src/$name/mainmin.exe";
# -H windowsgui -- without dos window
# -s -- Remove sign information
# -W -- Remove DWARF debug information

# uc:
go build -o release_win/main.exe $pkgName;release_win/main.exe;

go build -o release_win/mainmin.exe -ldflags="-s -w" $pkgName;release_win/mainmin.exe;

go build -o release_win/mainmincli.exe -ldflags="-s -w" $pkgName;release_win/mainmincli.exe;

go build -o release_win/maingui.exe -ldflags="-s -w -H windowsgui" $pkgName 
# release_win/maingui.exe # no thing output in console window

# uc:
go build -o release_win/${pkgName}.exe $pkgName;& release_win/${pkgName}.exe;

go build -o release_win/${pkgName}min.exe -ldflags="-s -w" $pkgName;& release_win/${pkgName}min.exe;

go build -o release_win/${pkgName}cli.exe -ldflags="-s -w" $pkgName;& release_win/${pkgName}cli.exe;

go build -o release_win/${pkgName}gui.exe -ldflags="-s -w -H windowsgui" $pkgName ;& release_win/${pkgName}gui.exe;

del release_win/*.exe -recurse
```

```bash

## prepare *.png files

- [x] copy file to project as background
```powershell
# copy-item "M:\bg-design\01\blue-sea-blue-sky_00003.png" icons/background.png
# del *.exe
```


- [x] fix `package chapter01/higo is not in std `
```powershell
# when: go build -o main chapter01/higo
# solv: add your path (ps: ./src/chapter02/higo ) to go.work in project root dir   
# why: 1. *.go files in your project root mod do not use any code from ./src/chapter02/higo
```
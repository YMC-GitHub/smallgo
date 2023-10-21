
# use fyne to make gui

## the baisc workflow

- [x] download fyne lib and cli
```powershell
go get fyne.io/fyne/v2@latest
# 
go install fyne.io/fyne/v2/cmd/fyne@latest
# download fyne-cross
go get -u github.com/fyne-io/fyne-cross

# or download the require pkg afer import them to your *.go file 
# go mod tidy
```

- [x] code your first app for fyne
```go
// ...
```

- [x] download the require pkg afer import them to your *.go file 
```powershell
# do these before go run xx, go build xx.
go mod tidy
```

- [x] run *.go files for dev
```powershell
# for file
go run main.go
# for some pkg-name
go run chapter02
```

- [x] build *.go files to *.exe files (when coding in project root dir)
```powershell
# go build -o release_win/fyne.exe .;release_win/fyne.exe
go build -o release_win/main.exe .;release_win/main.exe;
go build -o release_win/mainmin.exe -ldflags="-s -w" .;release_win/mainmin.exe;
# go build -o release_win/mainmincli.exe -ldflags="-s -w" .;release_win/mainmincli.exe
# no thing output in console window
go build -o release_win/maingui.exe -ldflags="-s -w -H windowsgui" . ;release_win/maingui.exe;
```

- [x] build *.exe files for dev
```powershell
# go mod tidy
go build -o release_win/fyne.exe .;release_win/fyne.exe;
```




- [x] fix `build constraints exclude all Go files in xx\gl`
```powershell
# when: go build xx.
# why:
# https://blog.csdn.net/booskz/article/details/126941067

# solution:  please enable cgo
go env -w CGO_ENABLED=1;


# go get -u github.com/fyne-io/fyne-cross
# fyne-cross windows -arch=amd64,386

# step 1. get the main config to analyze 
go env GOOS GOARCH CGO_ENABLED GOPATH

# step 2. enable CGO_ENABLED
go env -w GOOS=windows GOARCH=amd64 CGO_ENABLED=1;

# step 3. try build again
go build .
```

- [x] fix `cgo: C compiler "gcc" not found: exec: "gcc": executable file not found in %PATH%` 
```powershell
# step 1. have you installed gcc? if not, install one.
# choose one of them ? msys2,tdm-gcc,Cygwin,gcc
# scoop search gcc;
# scoop install gcc;
# 11.2.0,96MB

# scoop install tdm-gcc;
# 10.3.0,76MB

#  scoop search msys2
# scoop install msys2
# 2023-0718,55MB

# step 2. have you added it to os path? if not, add it.
# ...
```

## install msys2
- [x] install msys2 ( to as c/c++ compiler) in window with scoop 
```powershell
scoop install msys2
# now, youe can use its cli now.
# eg. gcc --version
# eg. msys2
```

warn: `Once installed do not use the MSYS terminal that opens` !!!
- [x] Open `“MSYS2 MinGW 64-bit”` from the start menu
- [x] install toolchain for `msys2 mingw 64`
```bash
# in msys2

# check prepare:
gcc --version;g++ --version;go version;

# copy in msys2 : ctrl + ins;
# paste in msys2 : shift + ins;

# Execute the following commands (if asked for install options be sure to choose “all”)
pacman -Syu
pacman -S git mingw-w64-x86_64-toolchain
# pacman -S mingw-w64-ucrt-x86_64-gcc
# gcc --version
```

- [x] MSYS2 uses windows environment variables
```powershell
M:\zero\scripts\idman\edit-env.ps1 set -envkey "MSYS2_PATH_TYPE" -value "inherit"
# refer: https://blog.csdn.net/a2824256/article/details/118599404
```

- [ ] add the go path to your os path 
```bash
# reading more:
# get your go version in msys2 bash
go --version
# get where your go cli is in msys2 bash
which go
# /d/CustomSoftware/windows/Scoop/shims/go

# you can get GO cli path of in bash as below (if install it with scoop and msys2 uses windows environment variables)
value=$(scoop prefix go);value=$(echo "/$value" | sed "s/://g" | sed 's,\\,\/,g' | sed 's,\/+,\/,g' | tr "[A-Z]" "[a-z]");echo "$value"
```


- [x] add the GOPATH/bin  to your os path 
```bash
# (the GOPATH is the value of `go env GOPATH`)
# add
# eg.
# echo "export PATH=\$PATH:/c/Program\ Files/Go/bin:~/Go/bin" >> ~/.bashrc
# eg.
value="/d/code/go/bin"; echo "export PATH=$value:\$PATH" >> ~/.bashrc

# use it in current bash
source ~/.bashrc

# del
# del it if you do not use it.
# eg.
# sed -i /Scoop/d ~/.bashrc
# sed -i /code/d ~/.bashrc
# sed -i /$value/d ~/.bashrc

# get
# cat ~/.bashrc

# use it in current bash
# source ~/.bashrc

# reading more:
# you can get GOPATH of go env in bash as below
value=$(go env GOPATH);value=$(echo "/$value" | sed "s/://g" | sed 's,\\,\/,g' | sed 's,\/+,\/,g' | tr "[A-Z]" "[a-z]");echo "$value"
```

- [x] For the compiler to work on other terminals you will need to set up the windows %PATH% variable to find these tools. 
```powershell
# $value="C:\msys64\mingw64\bin"
$value="D:\CustomSoftware\windows\Scoop\apps\msys2\current\mingw64\bin"
M:\zero\scripts\idman\edit-env-path.ps1 add -value "$value"

# reading more:
# you can get the value if install msys2 with scoop
# eg. for mingw64/bin
$value=scoop prefix msys2;$value="$value/mingw64/bin";$value=$value -replace "/","\";$value;
# eg. for mingw32/bin
$value=scoop prefix msys2;$value="$value/mingw32/bin";$value=$value -replace "/","\";$value;
# eg. for ucrt64/bin
$value=scoop prefix msys2;$value="$value/ucrt64/bin";$value=$value -replace "/","\";$value;

# reading more:
# now, you can switch your terminal to mingw64 in powershell
# mingw64
# and , you can mingw64 back to powershell
# exit
```



- [x] restore msys2 default os path  when you have made it to dirty!
```bash
export PATH=/mingw64/bin:/usr/local/bin:/usr/bin:/bin:/c/Windows/System32:/c/Windows:/c/Windows/System32/Wbem:/c/Windows/System32/WindowsPowerShell/v1.0/:/usr/bin/site_perl:/usr/bin/vendor_perl:/usr/bin/core_perl
```


- [x] get the current os path in msys2
```BASH
echo $PATH
# /d/code/go/bin:/mingw64/bin;XXXXXX

# /d/code/go/bin IS GOPATH. 
# XXXXXX IS INHERIT FROM WINDOWS PATH, I USE XX TO REPALCE IT HERE

# check if it is in .bashrc file
cat ~/.bashrc | grep "\$PATH"

# reading more:
# you can check your GOPATH FOR THESE
go env GOPATH
# C:\Users\zero\go ?

# get where your fyne is
which fyne
# /d/code/go/bin/fyne
```

- [x] get/reset the build env in msys2
```bash
go env GOOS GOARCH CGO_ENABLED CC GOPATH GOROOT
# windows
# amd64
# 1
# gcc
# C:\Users\zero\go
# D:\CustomSoftware\windows\Scoop\apps\go\current

# reset them if you need the default values
go env -u GOOS GOARCH CGO_ENABLED CC GOPATH GOROOT
```



- [x] set env values to build *.exe files in msys2
```bash
cd /d/code/go/smallgo.zero.com

# get current env values for buiding
go env GOOS GOARCH CGO_ENABLED CC GOPATH GOROOT

# to unset them if you need the default values
# go env -u GOOS GOARCH CGO_ENABLED CC GOPATH GOROOT

# set env values to build for windows,amd64 (enable cgo)
go env -w GOOS=windows;go env -w GOARCH=amd64;go env -w CGO_ENABLED=1

# run it
go run .
# go run chapter02

# build it
go build .
# go build chapter02

# reading more:
# you can also set build env in shell var or os var in mysy2
CC=gcc;CGO_ENABLED=1;GOARCH=amd64;GOOS=windows
# set env values to build for linux,arm64 (enable cgo)
go env -w GOOS=linux;go env -w GOARCH=arm64;go env -w CGO_ENABLED=1
```

[refer:1 Install one of the available C compilers for windows for fyne](https://developer.fyne.io/started/)

[refer:2 Build Fyne in Windos 10 Pro ](https://cloud.tencent.com/developer/article/2329625?areaId=106001)

[refer:3 Windows builds MSYS2 and MINGW64 environments](https://www.jianshu.com/p/96031565dafb)


## install fyne

- [x] check your fyne installation with fyne-setup.exe tool
```bash
# https://geoffrey-artefacts.fynelabs.com/github/andydotxyz/fyne-io/setup/latest/
```


## note more for fyne

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


## some fixs when using fyne

- [x] fix `fyne ` not found in powershell
```bash
# step 1. analyze where fyne is in bash
which fyne
# eg. /d/code/go/bin/fyne

# step 2. analyze why it is in that location
# eg. /d/code/go/bin/fyne
# i have set the GOPATH to there

# step 3. try for my idea
```

```powershell
# try to get/del/add value to os path in powershell script
$value="D:\code\go\bin";M:\zero\scripts\idman\edit-env-path.ps1 add -value $value;

# M:\zero\scripts\idman\edit-env-path.ps1 get
# M:\zero\scripts\idman\edit-env-path.ps1 del -value $value
M:\zero\scripts\idman\edit-env-path.ps1 add -value $value
```


- [x] fix `fyne: command not found when you have installed in it`
```powershell
# ADD GOPATH/BIN TO  OS PATH
# https://developer.fyne.io/faq/troubleshoot#compiler-issues
# go env GOPATH

# GET CURRENT OS PATH VALUE
M:\zero\scripts\idman\edit-env-path.ps1 get

# PREPARE YOUR GOPATH/BIN
$value=go env GOPATH;$value="$value\bin";

# ADD IT
M:\zero\scripts\idman\edit-env-path.ps1 add -value $value
# DEL IT
M:\zero\scripts\idman\edit-env-path.ps1 del -value $value
```

- [x] fix `could not fine gcc or clang compilers`
```powershell
# when: check your fyne installation with fyne-setuo.exe tools

# why:
# D:\CustomSoftware\windows\Scoop\apps\msys2\current\mingw64\bin
# [fyne & msy2 - add mingw64/bin to OS PATH](https://blog.csdn.net/sywdebug/article/details/126459784)

# step 1. get the value of the current GOTOOLDIR
go env GOTOOLDIR
# GOTOOLDIR=D:\CustomSoftware\windows\Scoop\apps\go\current\pkg\tool\windows_amd64
# it is the correct value ? change it?

# step 2. add the correct gcc bin path to os path
# eg. use the msys2 mingw64 bin
$value="D:\CustomSoftware\windows\Scoop\apps\msys2\current\mingw64\bin";
M:\zero\scripts\idman\edit-env-path.ps1 add -value $value;

# reading more:
# get which go is in powershell
gcm go

# from powershell to mysy2 mingw64 bash
mingw64

# from  mysy2 mingw64 bash to powershell
# exit
# get the current os path value in powershell with script file
M:\zero\scripts\idman\edit-env-path.ps1 get

# set/del some item of os path value in powershell with script file
# eg. use the msys2 mingw32 bin
$value="D:\CustomSoftware\windows\Scoop\apps\msys2\current\mingw32\bin";M:\zero\scripts\idman\edit-env-path.ps1 add -value $value;
# eg. del the msys2 mingw32 bin
$value="D:\CustomSoftware\windows\Scoop\apps\msys2\current\mingw32\bin";M:\zero\scripts\idman\edit-env-path.ps1 del -value $value;
# eg. use the msys2 mingw64 bin
$value="D:\CustomSoftware\windows\Scoop\apps\msys2\current\mingw64\bin";M:\zero\scripts\idman\edit-env-path.ps1 add -value $value;
# eg. del the msys2 mingw64 bin
$value="D:\CustomSoftware\windows\Scoop\apps\msys2\current\mingw64\bin";M:\zero\scripts\idman\edit-env-path.ps1 del -value $value;
# eg. use the msys2 ucrt64 bin
$value="D:\CustomSoftware\windows\Scoop\apps\msys2\current\ucrt64\bin";M:\zero\scripts\idman\edit-env-path.ps1 add -value $value;
# eg. del the msys2 ucrt64 bin
$value="D:\CustomSoftware\windows\Scoop\apps\msys2\current\ucrt64\bin";M:\zero\scripts\idman\edit-env-path.ps1 del -value $value;
```

- [x] fix `package xx is not in std `
```powershell
# eg. xx is chapter02/higo

when: go build -o main chapter01/higo
solution: add your path (ps: ./src/chapter02/higo ) to go.work in project root dir   
why: 1. *.go files in your project root mod do not use any code from ./src/chapter02/higo
```

- [x] fix `gcc_arm64.S:30: Error: no such instruction:` 
```bash
# when: go run xx ; go build xx;
# why: set  an error GOARCH to GOOS
# solution: go env -w GOOS=windows GOARCH=amd64 CGO_ENABLED=1;

# analyze:
# step 1. get current env values for builing
go env GOOS GOARCH CGO_ENABLED CC GOROOT GOPATH

# step 2. set the correct env values for different os.
# eg. for windows,amd64
go env -w GOOS=windows GOARCH=amd64 CGO_ENABLED=1;
```
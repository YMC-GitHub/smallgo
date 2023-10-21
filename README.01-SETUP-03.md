## use wails to make gui

- [x]install wails
```powershell
go install github.com/wailsapp/wails/v2/cmd/wails@latest


```

- [x] os check
```powershell
wails doctor
```

- [x] install WebView2 
```powershell
# choco list webview

```
[Wails requires that the WebView2 runtime is installed](https://wails.io/docs/gettingstarted/installation)


- [ ] Generate a Svelte project using TypeScript
```powershell
wails init -n myproject -t svelte-ts
```

- [x] Generate a react project using TypeScript
```powershell
wails init -n myproject -t react-ts
```

- [x] Generate a vue project using TypeScript
```powershell
wails init -n myproject -t vue-ts
```

- [x]  run application in development mode in your project directory
```
wails dev 
```

- [x] Build your application and run it

- [x] Bind your Go code to the frontend so it can be called from JavaScript

- [x] build *.go files, h5 files to *.exe files
```powershell
# https://wails.io/docs/reference/cli/#build
```

- [x] install upx in window with scoop
```powershell
# scoop search upx;
scoop install upx;
# 4.1.0,592kb
```

[upx  -- compressing your applications](https://upx.github.io/).

- [ ] install upx in window with choco
```powershell
choco install nsis
```

- [x] install nsis in window with scoop
```powershell
# scoop search nsis;
scoop install nsis;
```

[nsis -- generating Windows installers](https://wails.io/docs/guides/windows-installer/)


- [x] generate an installer for your application
```powershell
wails build -nsis
```
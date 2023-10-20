package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	fmt.Println("[info] Hello, Zero!")
	// use the same package methon
	// pkgFunc()

	hiye()
}

func GetAppThisFilePath() (thisfilepath string) {
	_, thisfilepath, _, _ = runtime.Caller(0)
	// fmt.Printf("[info] app file location: %s\n", thisfilepath)
	return thisfilepath
}
func GetExeThisFilePath() (thisfilepath string) {
	thisfilepath, err := filepath.Abs(os.Args[0])
	// dir, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("[info] app file location: %s\n", thisfilepath)
	return thisfilepath
}

func GetCurrentDir() (dir string) {
	// get the current working directory

	// using the function
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Printf("[info] app run in dir: %s\n", dir)
	return dir
}

func GetAppInDir() (dir string) {
	dir = filepath.Dir(GetAppThisFilePath())
	// using the function
	// dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	// dir, err := os.Executable()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("[info] app in dir: %s\n", dir)
	return dir
}

// func getAppInDirExeDiff() (dir string) {
// 	// dir = filepath.Dir(getAppThisFilePath())
// 	// using the function
// 	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
// 	// dir, err := os.Executable()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// fmt.Printf("[info] app in dir: %s\n", dir)
// 	return dir
// }

func AppDoThing() {
}
func AppLoadIconFromUrl(iconUri string) (iconResource fyne.Resource) {
	// iconUri := "https://liblibai-online.vibrou.com/web/image/13cc10f392b0f42ab2057f8558f123ee58c8812c49d12c1606751aa0b5cdddfc.png"
	iconResource, err := fyne.LoadResourceFromURLString(iconUri)
	if err != nil {
		AppDoThing()
		fmt.Println(err)
	}
	return iconResource
}
func AppLoadIconFromPath(iconUri string) (iconResource fyne.Resource) {
	// iconUri := "https://liblibai-online.vibrou.com/web/image/13cc10f392b0f42ab2057f8558f123ee58c8812c49d12c1606751aa0b5cdddfc.png"
	iconResource, err := fyne.LoadResourceFromPath(iconUri)
	if err != nil {
		AppDoThing()
		fmt.Println(err)
	}
	return iconResource
}

func hiye() {
	APP_THIS_FILE_PATH := GetAppThisFilePath()
	fmt.Printf("[info] app this file location: %s\n", APP_THIS_FILE_PATH)

	APP_IN_DIR_EXE_DIFF := GetExeThisFilePath()
	fmt.Printf("[info] app this file location (exe diff): %s\n", APP_IN_DIR_EXE_DIFF)

	APP_IN_DIR := GetAppInDir()
	fmt.Printf("[info] app in dir: %s\n", APP_IN_DIR)

	APP_RUN_IN_DIR := GetCurrentDir()
	fmt.Printf("[info] app run in dir: %s\n", APP_RUN_IN_DIR)

	// USE APP_RUN_IN_DIR AS APP_ROOT_DIR
	APP_ROOT_DIR := APP_RUN_IN_DIR

	// use APP_THIS_FILE_PATH/../../../../ as APP_ROOT_DIR
	APP_ROOT_DIR = path.Join(APP_THIS_FILE_PATH, "../../../../")
	fmt.Printf("[info] app root dir: %s\n", APP_ROOT_DIR)

	APP_ICON_DIR := path.Join(APP_ROOT_DIR, "icons")
	fmt.Printf("[info] app icon dir: %s\n", APP_ICON_DIR)

	// cahce resource uri in no-static-length array
	var resUri []string
	resUri = append(resUri, path.Join(APP_ICON_DIR, "background.png"), "https://liblibai-online.vibrou.com/web/image/13cc10f392b0f42ab2057f8558f123ee58c8812c49d12c1606751aa0b5cdddfc.png", path.Join(APP_ICON_DIR, "window.ico"))
	resUri = append(resUri, "https://ghproxy.com/https://raw.githubusercontent.com/ymc-github/cdn/master/data/2022-11-28/favicon-32X32.png")

	EN_REMOTE_ICON := false
	EN_REMOTE_BG := true

	a := app.New()
	w := a.NewWindow("launcher")

	iconUri := ""
	// iconUri = "D:\\code\\go\\smallgo.zero.com\\icons\\window.ico" //cause image: unknown format for icon ?

	// load icon from local path and set icon to app
	// iconUri = "D:\\code\\go\\smallgo.zero.com\\icons\\background.png"
	// iconUri = path.Join(APP_ICON_DIR, "background.png")
	// iconUri = path.Join(APP_ICON_DIR, "logo.ico")
	// iconUri = path.Join(APP_ICON_DIR, "window.ico")
	iconUri = resUri[2]
	// iconUri = resUri[3]
	fmt.Printf("[info] app icon uri: %s\n", iconUri)

	icon := AppLoadIconFromPath(iconUri)
	if icon != nil {
		w.SetIcon(icon)
	}

	// load remote icon switcher
	if EN_REMOTE_ICON {
		// load icon from remote url with http  and set icon to app
		// iconUri = "https://ghproxy.com/https://raw.githubusercontent.com/ymc-github/cdn/master/data/2022-11-28/favicon-32X32.png"
		// iconUri = "https://liblibai-online.vibrou.com/web/image/13cc10f392b0f42ab2057f8558f123ee58c8812c49d12c1606751aa0b5cdddfc.png"
		iconUri = resUri[1]
		// iconUri = resUri[3]
		icon = AppLoadIconFromUrl(iconUri)
		if icon != nil {
			w.SetIcon(icon)
		}
	}

	// set the window to fixed size
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(512, 768))

	// places the window at the center of screen
	w.CenterOnScreen()
	// w.SetTitle("")

	// w.SetFullScreen(false)

	// hello := widget.NewLabel("Hello Fyne!")
	// w.SetContent(container.NewVBox(
	// 	hello,
	// 	widget.NewButton("Hi!", func() {
	// 		hello.SetText("Welcome :)")
	// 	}),
	// ))

	imageUriBg := ""
	// imageUriBg = "D:\\code\\go\\smallgo.zero.com\\icons\\background.png"
	// imageUriBg = path.Join(APP_ICON_DIR, "background.png")
	imageUriBg = resUri[1]
	// imageUriBg = "G:\\code\\cdn\\data\\2022-11-28\\favicon-1024X1024.png"
	// load image from file
	image := canvas.NewImageFromFile(imageUriBg)

	// load remote image as bg switcher
	if EN_REMOTE_BG {
		// load image from remote url with http protocol
		// imageUriBg = "https://liblibai-online.vibrou.com/web/image/13cc10f392b0f42ab2057f8558f123ee58c8812c49d12c1606751aa0b5cdddfc.png"
		imageUriBg = resUri[1]
		resource, err := fyne.LoadResourceFromURLString(imageUriBg)
		if err != nil {
			fmt.Println(err)
		} else {
			image = canvas.NewImageFromResource(resource)
		}
	}

	// todo:
	// more: loading image from remote url to show in list (loop,random)

	// image := canvas.NewImageFromURI(imageUriBg)
	// image.FillMode = canvas.ImageFillOriginal

	// set the image size to match its contaner size
	image.FillMode = canvas.ImageFillStretch

	w.SetContent(image)

	// show window and run
	w.ShowAndRun()

}

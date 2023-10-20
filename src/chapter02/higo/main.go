package main

import (
	"fmt"
	"os"
	"path/filepath"

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

func getCurrentDir() (dir string) {
	// get the current working directory

	// using the function
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("[info] app run in dir: %s\n", dir)
	return dir
}

func getAppInDir() (dir string) {
	// using the function
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("[info] app in dir: %s\n", dir)
	return dir
}

func appDoThing() {
}
func appLoadIconFromUrl(iconUri string) (iconResource fyne.Resource) {
	// iconUri := "https://liblibai-online.vibrou.com/web/image/13cc10f392b0f42ab2057f8558f123ee58c8812c49d12c1606751aa0b5cdddfc.png"
	iconResource, err := fyne.LoadResourceFromURLString(iconUri)
	if err != nil {
		appDoThing()
		fmt.Println(err)
	}
	return iconResource
}
func appLoadIconFromPath(iconUri string) (iconResource fyne.Resource) {
	// iconUri := "https://liblibai-online.vibrou.com/web/image/13cc10f392b0f42ab2057f8558f123ee58c8812c49d12c1606751aa0b5cdddfc.png"
	iconResource, err := fyne.LoadResourceFromPath(iconUri)
	if err != nil {
		appDoThing()
		fmt.Println(err)
	}
	return iconResource
}

func hiye() {
	a := app.New()
	w := a.NewWindow("launcher")

	iconUri := ""
	// iconUri = "D:\\code\\go\\smallgo.zero.com\\icons\\window.ico" //cause image: unknown format for icon ?

	// load icon from local path and set icon to app
	iconUri = "D:\\code\\go\\smallgo.zero.com\\icons\\background.png"
	icon := appLoadIconFromPath(iconUri)
	if icon != nil {
		w.SetIcon(icon)
	}

	// load icon from remote url with http  and set icon to app
	// iconUri = "https://ghproxy.com/https://raw.githubusercontent.com/ymc-github/cdn/master/data/2022-11-28/favicon-32X32.png"
	iconUri = "https://liblibai-online.vibrou.com/web/image/13cc10f392b0f42ab2057f8558f123ee58c8812c49d12c1606751aa0b5cdddfc.png"
	icon = appLoadIconFromUrl(iconUri)
	if icon != nil {
		w.SetIcon(icon)
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

	getCurrentDir()
	getAppInDir()
	//
	// ~/AppData\Local\Temp\go-build2802241422\b001\exe\main.exe
	// [D:\code\go\smallgo.zero.com\release_win\fyne.exe]

	path, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}
	dir := filepath.Dir(path)
	fmt.Printf("[info] app in dir: %s\n", dir)

	filename := "D:\\code\\go\\smallgo.zero.com\\icons\\background.png"
	// filename = "G:\\code\\cdn\\data\\2022-11-28\\favicon-1024X1024.png"
	// load image from file
	image := canvas.NewImageFromFile(filename)

	// load image from remote url with http protocol
	uri := "https://liblibai-online.vibrou.com/web/image/13cc10f392b0f42ab2057f8558f123ee58c8812c49d12c1606751aa0b5cdddfc.png"
	r, err := fyne.LoadResourceFromURLString(uri)
	if err != nil {
		fmt.Println(err)
	} else {
		image = canvas.NewImageFromResource(r)
	}

	// todo:
	// more: loading image from remote url to show in list (loop,random)

	// image := canvas.NewImageFromURI(uri)
	// image.FillMode = canvas.ImageFillOriginal

	// set the image size to match its contaner size
	image.FillMode = canvas.ImageFillStretch

	w.SetContent(image)

	// show window and run
	w.ShowAndRun()

}

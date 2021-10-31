package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	//	"fyne.io/fyne/v2/container"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"fyne.io/fyne/v2"
	//"fyne.io/fyne/v2/widget"

	"fyne.io/fyne/v2/canvas"
)
func main() {
	a := app.New()
	w := a.NewWindow("Picture Gallery")
	w.Resize(fyne.NewSize(1280,720));
	root_src:="C:\\Users\\BIDHUBHUSAN\\Desktop\\Pepcoding lecture notes\\hackathon2\\images";
	files, err := ioutil.ReadDir(root_src);
	if err != nil {
        log.Fatal(err)
    }
	var picsArr []string;
	for _, file := range files {
        fmt.Println(file.Name(), file.IsDir())
		if file.IsDir() ==false{
			extension:= strings.Split(file.Name(), ".")[1];
			if extension == "jpg"|| extension == "jpeg"{
				picsArr = append(picsArr,root_src+"\\"+file.Name());
			}
		}
    }
	
	tabs := container.NewAppTabs(
		container.NewTabItem("Image1", canvas.NewImageFromFile(picsArr[0])),
		container.NewTabItem("Image2", canvas.NewImageFromFile(picsArr[1])),
		container.NewTabItem("Image3", canvas.NewImageFromFile(picsArr[2])),
		container.NewTabItem("Image4", canvas.NewImageFromFile(picsArr[3])),
	)
	w.SetContent(tabs);
	w.ShowAndRun()
}


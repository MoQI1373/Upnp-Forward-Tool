package main

import (
	"fmt"
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"github.com/fatih/color"
	"gitlab.com/NebulousLabs/go-upnp"
	"log"
	"os"
	"strconv"
)

var (
	Windows *ui.Window
	Status *ui.Entry
	Progress *ui.ProgressBar
)

func stringToInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}

func SubForm() ui.Control {
	box := ui.NewVerticalBox()
	box.SetPadded(true)

	forward, err := upnp.Discover()
	if err != nil {
		log.Fatal(err)
	}

	ip, err := forward.ExternalIP()
	if err != nil {
		log.Fatal(err)
	}

	EmptyNothing := ui.NewLabel("")
	box.Append(EmptyNothing, false)

	YourIP := ui.NewLabel("Your public ip: " + ip)
	box.Append(YourIP, false)

	group := ui.NewGroup("")
	group.SetMargined(true)
	box.Append(group, false)

	Form := ui.NewForm()
	Form.SetPadded(true)
	group.SetChild(Form)

	Empty := ui.NewLabel("")
	Form.Append("", Empty, false)

	Desc := ui.NewEntry()
	Form.Append("Description:", Desc, false)

	Port := ui.NewEntry()
	Form.Append("Port:", Port, false)

	EmptyText := ui.NewLabel("")
	Form.Append("", EmptyText, false)

	Status := ui.NewEntry()
	Form.Append("Status: ", Status, false)
	Status.SetReadOnly(true)

	Progress := ui.NewProgressBar()

	Clear := ui.NewButton("Clear All")
	Form.Append("", Clear, false)
	Clear.OnClicked(func(button *ui.Button){
		Desc.SetText("")
		Port.SetText("")
		Status.SetText("")
		Progress.SetValue(0)
	})


	Nothing := ui.NewLabel("")
	box.Append(Nothing, true)

	box.Append(Progress, false)

	NothingXD := ui.NewLabel("")
	box.Append(NothingXD, false)

	ForwardButton := ui.NewButton("Proceed")
	box.Append(ForwardButton, false)
	ForwardButton.OnClicked(func(button *ui.Button){
		DescText := Desc.Text()
		WrittenPort := Port.Text()

		PortNum, _ := strconv.ParseUint(WrittenPort, 0, 0)
		err = forward.Forward(uint16(PortNum), DescText)
		Status.SetText("Sucessful")
		Progress.SetValue(100)
		if err != nil {
			log.Fatal(err)
		}
	})

	return box
}

func DeleteForm() ui.Control {
	box := ui.NewVerticalBox()
	box.SetPadded(true)

	Nothing := ui.NewLabel("")
	box.Append(Nothing, false)

	Text := ui.NewLabel(" You can delete forwarded ports here")
	box.Append(Text, false)

	group := ui.NewGroup("")
	group.SetMargined(true)
	box.Append(group, false)

	Form := ui.NewForm()
	Form.SetPadded(true)
	group.SetChild(Form)

	Empty := ui.NewLabel("")
	Form.Append("", Empty, false)

	Port := ui.NewEntry()
	Form.Append("Port: ", Port, false)

	EmptyText := ui.NewLabel("")
	Form.Append("", EmptyText, false)

	Status := ui.NewEntry()
	Form.Append("Status: ", Status, false)
	Status.SetReadOnly(true)

	Progress := ui.NewProgressBar()

	Clear := ui.NewButton("Clear All")
	Form.Append("", Clear, false)
	Clear.OnClicked(func(button *ui.Button){
		Port.SetText("")
		Progress.SetValue(0)
		Status.SetText("")
	})

	LongEmpty := ui.NewLabel("")
	box.Append(LongEmpty, true)

	Progress.SetValue(0)
	box.Append(Progress, false)

	NothingXD := ui.NewLabel("")
	box.Append(NothingXD, false)

	DeleteButton := ui.NewButton("Delete Port")
	box.Append(DeleteButton, false)
	DeleteButton.OnClicked(func(button *ui.Button){
		Ports := Port.Text()
		PortNum, _ := strconv.ParseUint(Ports, 0, 0)

		forward, err := upnp.Discover()
		if err != nil {
			log.Fatal(err)
		}
		err = forward.Clear(uint16(PortNum))
		Status.SetText("Sucessful")
		Progress.SetValue(100)
		if err != nil {
			log.Fatal(err)
		}
	})


	return box
}

func AboutForm() ui.Control {
	box := ui.NewVerticalBox()
	box.SetPadded(true)

	var Empty = ui.NewLabel("")
	box.Append(Empty, false)

	var Name = ui.NewLabel("Name: Upnp Forward")
	box.Append(Name, false)

	var Author = ui.NewLabel("Author: OwoNico")
	box.Append(Author, false)

	var Version = ui.NewLabel("Version: 1.0")
	box.Append(Version, false)

	return box
}

func Upnp() {
	Window := ui.NewWindow("Upnp Forward", 420, 550, false)
	Window.OnClosing(func(*ui.Window) bool {
		ui.Quit()
		return true
	})
	ui.OnShouldQuit(func() bool {
		Window.Destroy()
		return true
	})

	tabs := ui.NewTab()
	Window.SetChild(tabs)
	Window.SetMargined(true)

	tabs.Append("Upnp Forward",SubForm())
	tabs.SetMargined(0,true)

	tabs.Append("Delete Ports", DeleteForm())
	tabs.SetMargined(0, true)

	tabs.Append("About", AboutForm())
	tabs.SetMargined(0, true)

	Window.Show()
}

func main()  {
	color.Cyan("Upnp Forwarding tool by OwoNico")
	color.Blue("#https://github.com/OwoNico")
	fmt.Println("")

	w := color.New(color.FgWhite).Add(color.Bold)
	y := color.New(color.FgYellow).Add(color.Bold)
	g := color.New(color.FgGreen).Add(color.Bold)

	w.Println("Loading Upnp Tool...")
	w.Println("Connecting to router...")
	w.Println("Loading UI...")
	w.Println("Loading API...")
	w.Println("Loaded all threads")
	fmt.Println("")
	y.Println("Launching Upnp Forward...")
	g.Println("Launched")
	ui.Main(Upnp)

}

package watchgoggy

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
)

// App indicate config infos of single monitor app
type App struct {
	AppName string   `json:"AppName"`
	Cmd     string   `json:"Cmd"`
	Params  []string `json:"Params"`
	Pid     int
}

var (
	//Applications all monitor apps
	Applications = make(map[string]App)
)

var goggies = "goggies.json"

//init read config infos from goggies
func init() {
	if os.Getenv("goggies") != "" {
		reloadGoggies(os.Getenv("goggies"))
	} else {
		inArgs, s := goggiesInArgs()
		if inArgs {
			reloadGoggies(s)
		} else {
			reloadGoggies(goggies)
		}
	}

}

//goggiesInArgs if params goggies in args，return true and value，or else return false and ""
func goggiesInArgs() (bool, string) {
	var file string
	flag.StringVar(&file, "goggies", "", "plz input -goggies xxx to make sure the config files")
	flag.Parse()
	if file == "" {
		return false, ""
	}
	return true, file
}

func GetApp(name string) *App {
	app, exist := Applications[name]
	if exist {
		return &app
	}
	return nil
}

func RemoveApp(name string) {
	delete(Applications, name)
}

//reloadGoggies reload app config infos from goggies
func reloadGoggies(file string) {
	Applications = make(map[string]App)
	bytes, err1 := ioutil.ReadFile(file)
	if err1 != nil {
		panic("cant read config file, plz check")
	}
	apps := make([]App, 0)
	err2 := json.Unmarshal(bytes, &apps)
	if err2 != nil {
		panic("cant deserialize config file, plz check")
	}
	for _, app := range apps {
		Applications[app.AppName] = app
	}
}

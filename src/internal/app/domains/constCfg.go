package domains

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

type Cfg struct {
	WorkDir     string `json:"WorkDir"`
	AssetsDir   string `json:"AssetsDir"`
	LogDir      string `json:"LogDir"`
	TempFileDir string `json:"TempFileDir"`

	TempGraph   string `json:"TempGraph"`
	HistoryFile string `json:"HistoryFile"`

	XWindowGraph uint32 `json:"XWindowGraph"`
	YWindowGraph uint32 `json:"YWindowGraph"`

	DarkTheme string `json:"DarkTheme"`
	IconPath  string `json:"IconPath"`
	TypePath  string `json:"TypePath"`

	Debug bool `json:"Debug"`
}

const (
	configPathMain = "/Applications/smartCalc.app/Contents/Resources/config/smartCalc.cfg"
	configPathOpt  = "config/smartCalc.cfg"
)

var Config *Cfg = InitConfig("")

func createNewConfig() *Cfg {
	var c Cfg

	c.WorkDir = "./"
	c.AssetsDir = c.WorkDir + "assets/"
	c.LogDir = c.WorkDir + "log/"
	c.TempFileDir = c.WorkDir + "temp_file/"
	c.TempGraph = "tempGraph.png"
	c.HistoryFile = "history.json"
	c.XWindowGraph = 600
	c.YWindowGraph = 600
	c.DarkTheme = "no"
	c.IconPath = c.AssetsDir + "Icon.png"
	c.TypePath = c.AssetsDir + "protosans56.ttf"
	c.Debug = false

	data, _ := json.MarshalIndent(c, "", "    ")
	err := os.WriteFile(configPathOpt, data, 0777)

	os := runtime.GOOS
	switch os {
	case "windows":
		fmt.Println("Windows")
	case "darwin":
		fmt.Println("MAC operating system")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}

	if err == nil {
		fmt.Println("Create and write config to:", configPathOpt)
	} else {
		fmt.Println("Cannot write config to:", configPathOpt)
	}
	return &c
}

func readConfig(fileName string, c *Cfg) error {
	var err error
	if dataFromFile, err := os.ReadFile(fileName); err == nil {
		if err = json.Unmarshal(dataFromFile, c); err == nil {
			fmt.Println("Load config from:", fileName)
			return err
		}
	} else {
		return err
	}
	return err
}

func InitConfig(fileName string) *Cfg {
	var c Cfg

	if fileName != "" {
		if err := readConfig(fileName, &c); err == nil {
			return &c
		}
	}

	if err := readConfig(configPathMain, &c); err == nil {
		return &c
	}

	if err := readConfig(configPathOpt, &c); err == nil {
		return &c
	}

	return createNewConfig()
}

func (c *Cfg) GetWorkDir() string     { return c.WorkDir }
func (c *Cfg) GetAssetsDir() string   { return c.AssetsDir }
func (c *Cfg) GetLogDir() string      { return c.LogDir }
func (c *Cfg) GetTempFileDir() string { return c.TempFileDir }
func (c *Cfg) GetTempGraph() string   { return c.TempGraph }
func (c *Cfg) GetHistoryFile() string { return c.HistoryFile }
func (c *Cfg) GetXWindowGraph() int   { return int(c.XWindowGraph) }
func (c *Cfg) GetYWindowGraph() int   { return int(c.YWindowGraph) }
func (c *Cfg) GetDarkTheme() string   { return c.DarkTheme }
func (c *Cfg) GetIconPath() string    { return c.IconPath }
func (c *Cfg) GetTypePath() string    { return c.TypePath }

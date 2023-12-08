package domains

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
)

type Cfg struct {
	WorkDir     string `json:"WorkDir"`     // Directory with programm
	AssetsDir   string `json:"AssetsDir"`   // Directory with help.txt and type
	LogDir      string `json:"LogDir"`      // Directory for log
	TempFileDir string `json:"TempFileDir"` // Directory for Tempfile

	TempGraph   string `json:"TempGraph"`   // tempfile name
	HistoryFile string `json:"HistoryFile"` // historyfile path (with directory)

	XWindowGraph uint32 `json:"XWindowGraph"` // Graph window size X
	YWindowGraph uint32 `json:"YWindowGraph"` // Graph window size Y

	DarkTheme string `json:"DarkTheme"` //Dark mode "yes" or "no"
	IconPath  string `json:"IconPath"`  //iconfile name
	TypePath  string `json:"TypePath"`  //Typefile name

	Debug bool `json:"Debug"` //debug mode with output additional info to terminal -  true or false
}

var (
	// Path for config in linux system
	// [0] - main; [1] - optional
	ConfigLinuxPath = []string{
		"/etc/smartCalc/smartCalcLinux.cfg",
		"config/smartCalcLinux.cfg",
	}

	// Path for config in Mac system
	userDir, _ = os.Getwd()
	testDir    = userDir[:len(userDir)-19]

	ConfigMacPath = []string{
		"/Applications/smartCalc.app/Contents/Resources/config/smartCalcMacOS.cfg",
		userDir + "/config/smartCalcMacOSIn.cfg",
		testDir + "/config/smartCalcMacOSIn.cfg",
	}

	Os          = runtime.GOOS   // "windows", "darwin", "linux"
	Arch        = runtime.GOARCH // "amd64", "386", "arm"
	Config *Cfg = InitConfig("") // Handling config path by type config name in quotes (but this way not recommend)
)

// Create and write new config for Mac
func createNewMacConfig() *Cfg {
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
	err := os.WriteFile(ConfigMacPath[0], data, 0777)
	if err == nil {
		fmt.Println("Create and write config to:", ConfigMacPath[0])
	} else {
		fmt.Println("Cannot write config to:", ConfigMacPath[0])
	}
	return &c
}

// Create and write new config for Linux
func createNewLinuxConfig() *Cfg {
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
	err := os.WriteFile(ConfigLinuxPath[1], data, 0777)
	if err == nil {
		fmt.Println("Create and write config to:", ConfigLinuxPath[1])
	} else {
		fmt.Println("Cannot write config to:", ConfigLinuxPath[1])
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

// Inicialize config
func InitConfig(fileName string) *Cfg {
	var c Cfg

	// fmt.Println("testDir:", userDir)

	if fileName != "" {
		if err := readConfig(fileName, &c); err == nil {
			return &c
		}
	}

	switch Os {
	case "linux":
		for _, path := range ConfigLinuxPath {
			if err := readConfig(path, &c); err == nil {
				return &c
			}
		}
		return createNewLinuxConfig()

	case "darwin":
		for _, path := range ConfigMacPath {
			if err := readConfig(path, &c); err == nil {
				return &c
			}
		}
		return createNewMacConfig()
	}

	return createNewLinuxConfig()
}

// func (c *Cfg) GetWorkDir() string     { return c.WorkDir }
// func (c *Cfg) GetAssetsDir() string   { return c.AssetsDir }
// func (c *Cfg) GetLogDir() string      { return c.LogDir }
// func (c *Cfg) GetTempFileDir() string { return c.TempFileDir }
// func (c *Cfg) GetTempGraph() string   { return c.TempGraph }
// func (c *Cfg) GetHistoryFile() string { return c.HistoryFile }
// func (c *Cfg) GetXWindowGraph() int   { return int(c.XWindowGraph) }
// func (c *Cfg) GetYWindowGraph() int   { return int(c.YWindowGraph) }
// func (c *Cfg) GetDarkTheme() string   { return c.DarkTheme }
// func (c *Cfg) GetIconPath() string    { return c.IconPath }
// func (c *Cfg) GetTypePath() string    { return c.TypePath }

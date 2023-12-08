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
	Os              = runtime.GOOS   // "windows", "darwin", "linux"
	Arch            = runtime.GOARCH // "amd64", "386", "arm"
	userDir, _      = os.Getwd()
	testDir         = testDirFind(userDir)
	ConfigLinuxPath = []string{
		"/etc/smartCalc/smartCalcLinux.cfg",
		userDir + "/config/smartCalcLinuxIn.cfg",
		testDir + "/config/smartCalcTest.cfg",
	}

	// Path for config in Mac system

	ConfigMacPath = []string{
		"/Applications/smartCalc.app/Contents/Resources/config/smartCalcMacOS.cfg",
		userDir + "/config/smartCalcMacOSIn.cfg",
		testDir + "/config/smartCalcTest.cfg",
	}

	Config *Cfg = InitConfig("") // Handling config path by type config name in quotes (but this way not recommend)
)

// // Create and write new config for Mac
// func createNewMacConfig() *Cfg {
// 	var c Cfg

// 	c.WorkDir = userDir
// 	c.AssetsDir = c.WorkDir + "/assets/"
// 	c.LogDir = c.WorkDir + "/log/"
// 	c.TempFileDir = c.WorkDir + "/tmp/"
// 	c.TempGraph = "tempGraph.png"
// 	c.HistoryFile = c.WorkDir + "/var/history.json"
// 	c.XWindowGraph = 600
// 	c.YWindowGraph = 600
// 	c.DarkTheme = "no"
// 	c.IconPath = c.AssetsDir + "Icon.png"
// 	c.TypePath = c.AssetsDir + "protosans56.ttf"
// 	c.Debug = false

// 	fmt.Println("New config file created")

// 	return &c
// }

// // Create and write new config for Linux
// func createNewLinuxConfig() *Cfg {
// 	var c Cfg

// 	c.WorkDir = userDir
// 	c.AssetsDir = c.WorkDir + "/assets/"
// 	c.LogDir = c.WorkDir + "/log/"
// 	c.TempFileDir = c.WorkDir + "/tmp/"
// 	c.TempGraph = "tempGraph.png"
// 	c.HistoryFile = c.WorkDir + "/var/history.json"
// 	c.XWindowGraph = 600
// 	c.YWindowGraph = 600
// 	c.DarkTheme = "no"
// 	c.IconPath = c.AssetsDir + "Icon.png"
// 	c.TypePath = c.AssetsDir + "protosans56.ttf"
// 	c.Debug = false

// 	fmt.Println("New config file created")

// 	return &c
// }

// Create and write new config file
func createNewConfig() *Cfg {
	var c Cfg

	c.WorkDir = userDir
	c.AssetsDir = c.WorkDir + "/assets/"
	c.LogDir = c.WorkDir + "/log/"
	c.TempFileDir = c.WorkDir + "/tmp/"
	c.TempGraph = "tempGraph.png"
	c.HistoryFile = c.WorkDir + "/var/history.json"
	c.XWindowGraph = 600
	c.YWindowGraph = 600
	c.DarkTheme = "no"
	c.IconPath = c.AssetsDir + "Icon.png"
	c.TypePath = c.AssetsDir + "protosans56.ttf"
	c.Debug = false

	fmt.Println("New config file created")

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
		// return createNewLinuxConfig()
		return createNewConfig()

	case "darwin":
		for _, path := range ConfigMacPath {
			if err := readConfig(path, &c); err == nil {
				return &c
			}
		}
		// return createNewMacConfig()
		return createNewConfig()
	}

	// return createNewLinuxConfig()
	return createNewConfig()
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

func testDirFind(userDir string) string {
	if len(userDir) > 20 {
		return userDir[:len(userDir)-19]
	}
	return userDir
}

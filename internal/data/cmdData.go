package data

import (
	"log"

	"github.com/spf13/cobra"
)

var DataPathFlagName = "data"

// TODO (ask): is this a good way to do inheritance by composition that go promotes?
type CmdData struct {
	cmd *cobra.Command
	Data
}

func getCmdDataPath(cmd *cobra.Command) string {
	// TODO: handle errors
	dataPath, _ := cmd.Flags().GetString(DataPathFlagName)
	return dataPath
}

func (c *CmdData) Load(cmd *cobra.Command) *CmdData {
	if c == nil {
		c = &CmdData{}
	}
	c.cmd = cmd
	c.Data = LoadData(getCmdDataPath((cmd)))
	return c
}

func (c *CmdData) Save() {
	if c == nil {
		log.Fatalln("Pointer to cmdData is nil")
	}
	SaveData(getCmdDataPath(c.cmd), c.Data)
}

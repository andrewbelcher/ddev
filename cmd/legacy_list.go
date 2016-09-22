package cmd

import (
	"github.com/drud/bootstrap/cli/local"
	"github.com/fsouza/go-dockerclient"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var LegacyListCmd = &cobra.Command{
	Use:   "list",
	Short: "List applications that exist locally",
	Long:  `List applications that exist locally.`,
	Run: func(cmd *cobra.Command, args []string) {

		client, _ := GetDockerClient()

		containers, _ := client.ListContainers(docker.ListContainersOptions{All: true})
		containers = local.FilterNonLegacy(containers)

		local.RenderContainerList(containers)
	},
}

func init() {
	LegacyCmd.AddCommand(LegacyListCmd)

}

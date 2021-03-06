package cmd

import (
	"github.com/spf13/cobra"

	"github.com/joelanford/kubectl-operator/internal/pkg/action"
)

func newCatalogCmd(cfg *action.Configuration) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "catalog",
		Short: "Manage operator catalogs",
	}
	cmd.AddCommand(
		newCatalogAddCmd(cfg),
		newCatalogListCmd(cfg),
		newCatalogRemoveCmd(cfg),
	)
	return cmd
}

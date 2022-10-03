package commands

import (
	"strconv"

	"github.com/gandarfh/maid-san/internal/commands/resources/repository"
	"github.com/gandarfh/maid-san/pkg/errors"
	"github.com/gandarfh/maid-san/pkg/repl"
	"github.com/gandarfh/maid-san/pkg/table"
	"github.com/gandarfh/maid-san/pkg/utils"
)

type List struct {
	envs *[]repository.Resources
}

func (c *List) Read(args ...string) error {
	return nil
}

func (c *List) Eval() error {
	repo, err := repository.NewResourcesRepo()
	if err != nil {
		return errors.InternalServer("Error when connect to database!")
	}

	c.envs = repo.List()

	return nil
}

func (c *List) Print() error {
	tbl := table.NewTable([]string{"id", "parent", "name", "endpoint", "method"})
	rows := []table.Row{}

	for _, item := range *c.envs {
		parent := item.Parent()

		row := table.Row{
			strconv.FormatUint(uint64(item.ID), 10),
			utils.ReplaceByEnv(parent.Name),
			utils.ReplaceByEnv(item.Name),
			utils.ReplaceByEnv(item.Endpoint),
			utils.ReplaceByEnv(item.Method),
		}
		rows = append(rows, row)
	}

	tbl.SetRows(rows)

	return nil
}

func (w *List) Run(args ...string) error {
	if err := w.Read(args...); err != nil {
		return err
	}

	if err := w.Eval(); err != nil {
		return err
	}

	if err := w.Print(); err != nil {
		return err
	}

	return nil
}

func ListInit() repl.Repl {
	return &List{}
}

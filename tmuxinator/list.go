package tmuxinator

import (
	"fmt"
	"strings"

	"github.com/joshmedeski/sesh/model"
)

func (t *RealTmuxinator) List() ([]*model.TmuxinatorConfig, error) {
	res, err := t.shell.ListCmd("tmuxinator", "list")
	if err != nil {
		return []*model.TmuxinatorConfig{}, err
	}
	return parseTmuxinatorConfigsOutput(res)
}

func parseTmuxinatorConfigsOutput(rawList []string) ([]*model.TmuxinatorConfig, error) {
	cleanedList := strings.Split(rawList[1], "  ")
	sessions := make([]*model.TmuxinatorConfig, 0, len(cleanedList))
	for _, line := range cleanedList {
		if len(line) > 0 {
			fmt.Println("line: ", line)
			session := &model.TmuxinatorConfig{
				Name: strings.TrimSpace(line),
			}
			sessions = append(sessions, session)
		}
	}

	return sessions, nil
}
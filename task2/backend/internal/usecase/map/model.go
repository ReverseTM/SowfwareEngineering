package _map

import (
	"encoding/json"
	"fmt"
	"software-engineering-2/internal/model"
	"software-engineering-2/internal/usecase/map/general/command"
	"software-engineering-2/internal/usecase/map/general/command/commands"
)

type MapData struct {
	Map     *model.RoadMap `json:"map"`
	History *History       `json:"history"`
}

type History struct {
	UndoCommands []command.Command `json:"undo"`
	RedoCommands []command.Command `json:"redo"`
}

func (h *History) UnmarshalJSON(data []byte) error {
	var raw struct {
		Undo []json.RawMessage `json:"undo"`
		Redo []json.RawMessage `json:"redo"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	parseCommands := func(rawList []json.RawMessage) ([]command.Command, error) {
		cmds := make([]command.Command, 0, len(rawList))
		for _, rawCmd := range rawList {
			var typeHolder struct {
				Action string `json:"action"`
			}

			if err := json.Unmarshal(rawCmd, &typeHolder); err != nil {
				return nil, err
			}

			var cmd command.Command
			switch typeHolder.Action {
			case "ADD-CITY":
				var c commands.AddCityCommand
				if err := json.Unmarshal(rawCmd, &c); err != nil {
					return nil, err
				}
				cmd = &c
			case "UPDATE-CITY":
				var c commands.UpdateCityCommand
				if err := json.Unmarshal(rawCmd, &c); err != nil {
					return nil, err
				}
				cmd = &c
			case "DELETE-CITY":
				var c commands.DeleteCityCommand
				if err := json.Unmarshal(rawCmd, &c); err != nil {
					return nil, err
				}
				cmd = &c
			case "ADD-ROAD":
				var c commands.AddRoadCommand
				if err := json.Unmarshal(rawCmd, &c); err != nil {
					return nil, err
				}
				cmd = &c
			case "UPDATE-ROAD":
				var c commands.UpdateRoadCommand
				if err := json.Unmarshal(rawCmd, &c); err != nil {
					return nil, err
				}
				cmd = &c
			case "DELETE-ROAD":
				var c commands.DeleteRoadCommand
				if err := json.Unmarshal(rawCmd, &c); err != nil {
					return nil, err
				}
				cmd = &c
			default:
				return nil, fmt.Errorf("unknown command action: %s", typeHolder.Action)
			}

			cmds = append(cmds, cmd)
		}
		return cmds, nil
	}

	var err error
	h.UndoCommands, err = parseCommands(raw.Undo)
	if err != nil {
		return err
	}

	h.RedoCommands, err = parseCommands(raw.Redo)
	if err != nil {
		return err
	}

	return nil
}

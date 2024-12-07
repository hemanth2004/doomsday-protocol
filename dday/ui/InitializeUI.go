package ui

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
	"github.com/hemanth2004/doomsday-protocol/dday/util"
	"github.com/hemanth2004/doomsday-protocol/dday/util/tree"

	"github.com/evertras/bubble-table/table"
)

type BoxResolution [2]int

type ColumnKeyWidthPair struct {
	Key      string
	Width    int
	Flexible bool
}

var (
	idPair = ColumnKeyWidthPair{
		Key:      "id",
		Width:    3,
		Flexible: false,
	}
	namePair = ColumnKeyWidthPair{
		Key:      "element",
		Width:    3,
		Flexible: true,
	}
	progressBarPair = ColumnKeyWidthPair{
		Key:      "bar",
		Width:    4,
		Flexible: true,
	}
	statusPair = ColumnKeyWidthPair{
		Key:      "status",
		Width:    11,
		Flexible: false,
	}
	sizePair = ColumnKeyWidthPair{
		Key:      "size",
		Width:    2,
		Flexible: true,
	}
	speedPair = ColumnKeyWidthPair{
		Key:      "speed",
		Width:    13,
		Flexible: false,
	}
	etaPair = ColumnKeyWidthPair{
		Key:      "eta",
		Width:    2,
		Flexible: true,
	}
	downloadTableColumns = []table.Column{
		table.NewColumn(idPair.Key, "ID", idPair.Width),
		table.NewFlexColumn(namePair.Key, "Name", namePair.Width),
		table.NewFlexColumn(progressBarPair.Key, "Progress", progressBarPair.Width),
		table.NewColumn(statusPair.Key, "Status", statusPair.Width),
		table.NewFlexColumn(sizePair.Key, "Size", sizePair.Width),
		table.NewColumn(speedPair.Key, "Speed", speedPair.Width),
		table.NewFlexColumn(etaPair.Key, "ETA", etaPair.Width),
	}
)

func InitialTeaModel(Application *core.Application) MainModel {
	return MainModel{

		Application:  Application,
		ResourceList: &Application.ResourceList,

		CurrentState: util.NewStateHandler([]string{"guides", "downloads", "new resource"}, 1),
		Downloads: DownloadsModel{
			ResourceList: &Application.ResourceList,

			LogFunction:   &Application.LogFunction,
			CurrentWindow: util.NewStateHandler([]int{2, 1, 0}, 2),

			DownloadsTable: table.New(downloadTableColumns).
				Border(styles.CustomBorder).
				WithBaseStyle(styles.TableStyle).
				WithMultiline(true),
			//WithPageSize(2),

			ResourceTree: tree.New([]tree.Node{}),

			ConsoleModel: ConsoleModel{
				Viewport:      viewport.New(5, 5),
				ConsoleOpened: true,
				LogsContent:   [][2]string{},
			},

			HelpSet: InitDownloadsHelpSet(),
		},
		NewResource: NewResourceModel{},
		Guides:      GuidesModel{},

		HelpSet: InitMainHelpSet(),
	}
}

func InitMainHelpSet() HelpSet {
	return HelpSet{
		{"ctrl+q/e", "switch tabs"},
	}
}

func InitDownloadsHelpSet() []HelpSet {
	return []HelpSet{
		// Console
		{
			{"enter", "close/open"},
			{"↑/↓", "navigate"},
			{"tab", "switch panels"},
		},
		// Download Inspector
		{
			{"↑/↓", "navigate about"},
			{"tab", "switch panels"},
		},
		// Downloads Table
		{
			{"↑/↓", "navigate"},
			{"space", "pause selected"},
			{"tab", "switch panels"},
		},
	}
}

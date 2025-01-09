package ui

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/submodels"
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
		Width:    7,
		Flexible: false,
	}
	statusPair = ColumnKeyWidthPair{
		Key:      "status",
		Width:    11,
		Flexible: false,
	}
	sizePair = ColumnKeyWidthPair{
		Key:      "size",
		Width:    16,
		Flexible: false,
	}
	speedPair = ColumnKeyWidthPair{
		Key:      "speed",
		Width:    13,
		Flexible: false,
	}
	etaPair = ColumnKeyWidthPair{
		Key:      "eta",
		Width:    8,
		Flexible: false,
	}
	downloadTableColumns = []table.Column{
		table.NewColumn(idPair.Key, "ID", idPair.Width),
		table.NewFlexColumn(namePair.Key, "Name", namePair.Width),
		table.NewColumn(progressBarPair.Key, "Progress", progressBarPair.Width).WithStyle(lipgloss.NewStyle().Width(namePair.Width).Align(lipgloss.Right)),
		table.NewColumn(statusPair.Key, "Status", statusPair.Width),
		table.NewColumn(sizePair.Key, "Size", sizePair.Width),
		table.NewColumn(speedPair.Key, "Speed", speedPair.Width),
		table.NewColumn(etaPair.Key, "ETA", etaPair.Width),
	}
)

func InitialTeaModel(Application *core.Application) MainModel {
	return MainModel{

		Application:  Application,
		ResourceList: &Application.ResourceList,

		CurrentState: util.NewStateHandler([]string{"home", "downloads", "new resource"}, 0),
		Downloads: DownloadsModel{
			ResourceList:   &Application.ResourceList,
			LogFunction:    &Application.LogFunction,
			LogsContentRef: &Application.LogsContent,

			CurrentWindow: util.NewStateHandler([]int{2, 1, 0}, 2),

			DownloadsTable: table.New(downloadTableColumns).
				Border(styles.CustomBorder).
				WithBaseStyle(styles.TableStyle).
				WithMultiline(false),
			//WithPageSize(2),

			ResourceTree: tree.New([]tree.Node{}, 5, 5),

			ConsoleModel: submodels.ConsoleModel{
				Viewport:      viewport.New(5, 5),
				ConsoleOpened: true,
				LogsContent:   [][2]string{},
			},

			HelpSet: InitDownloadsHelpSet(),
		},
		NewResource: NewResourceModel{},
		Home: HomeModel{
			CurrentWindow: util.NewStateHandler([]int{2, 1, 0}, 0),
			TextViewer: submodels.TextViewerModel{
				Path:      "packaged/Guides/_welcome.md",
				Scrollbar: submodels.NewScrollbar(),
			},
			GuideTree: submodels.GuideTreeModel{
				GuidesPath:        Application.GuidesFolderPath,
				ReadGuideCallback: Application.GuideViewerCallback,
				Scrollbar:         submodels.NewScrollbar(),
			},
			StatusModel: submodels.StatusModel{
				ApplicationObject: Application,
				Progress:          submodels.NewMultilineProgress(5, 3, styles.StatusFillStyle, styles.StatusBackgroundStyle),
			},
			HelpSet: InitHomeHelpSet(),
		},

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
			{"ctrl+d", "show ctrl panel"},
			{"tab", "switch panels"},
		},
		// Download Inspector
		{
			{"↑/↓", "navigate about"},
			{"ctrl+d", "show ctrl panel"},
			{"tab", "switch panels"},
		},
		// Downloads Table
		{
			{"↑/↓", "navigate"},
			{"ctrl+d", "show ctrl panel"},
			{"tab", "switch panels"},
		},
	}
}

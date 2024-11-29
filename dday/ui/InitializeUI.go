package ui

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/hemanth2004/doomsday-protocol/dday/core"
	"github.com/hemanth2004/doomsday-protocol/dday/ui/styles"
	"github.com/hemanth2004/doomsday-protocol/dday/util"
	"github.com/hemanth2004/doomsday-protocol/dday/util/tree"

	"github.com/evertras/bubble-table/table"
)

const (
	columnKeyID          = "id"
	columnKeyName        = "element"
	columnKeyProgressBar = "bar"
	columnKeyStatus      = "status"
	columnKeySpeed       = "speed"
	columnKeyETA         = "eta"
)

func InitialTeaModel(Application *core.Application) MainModel {
	return MainModel{

		Application:  Application,
		ResourceList: &Application.ResourceList,

		CurrentState: util.NewStateHandler([]string{"guides", "downloads", "new resource"}),
		Downloads: DownloadsModel{
			ResourceList: &Application.ResourceList,

			CurrentWindow: util.NewStateHandler([]int{2, 1, 0}),

			DownloadsTable: table.New([]table.Column{
				table.NewColumn(columnKeyID, "ID", 3),
				table.NewFlexColumn(columnKeyName, "Name", 2),
				table.NewFlexColumn(columnKeyProgressBar, "Progress", 2),
				table.NewFlexColumn(columnKeyStatus, "Status", 1),
				table.NewFlexColumn(columnKeySpeed, "Speed", 1),
				table.NewFlexColumn(columnKeyETA, "ETA", 1),
			}).Border(styles.CustomBorder).WithBaseStyle(styles.TableStyle),

			ResourceTree: tree.New([]tree.Node{}),

			ConsoleViewport: viewport.New(5, 5),
		},
		NewResource: NewResourceModel{},
		Guides:      GuidesModel{},
	}
}

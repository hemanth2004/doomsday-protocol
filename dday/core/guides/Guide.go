package guides

type Guide struct {
	Title   string // To display on the guide viewer
	Content string // To display on the guide viewer
	Format  string // ".md" for markdown, ".txt" for plain text etc.
}

type ChangeViewingGuideMsg string

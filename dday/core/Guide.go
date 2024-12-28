package core

type Guide struct {
	Id string // For other guides to be able to link to this guide using the <a> tag

	ShortName string // To display on the guide tree
	Title     string // To display on the guide viewer
	Content   string // To display on the guide viewer
	Format    string // ".md" for markdown, ".txt" for plain text etc.
}

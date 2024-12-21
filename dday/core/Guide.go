package core

type Guide struct {
	ShortName string // To display on the guide tree
	Title     string // To display on the guide viewer
	Content   string // To display on the guide viewer
	Format    string // ".md" for markdown, ".txt" for plain text etc.
}

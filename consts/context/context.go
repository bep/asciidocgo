package context

// Symbol name for the type of content (e.g., :paragraph).
type Context int

const (
	Document Context = iota
	Section
	Paragraph
	// Used by substitutors in SubMacros()
	Kbd
	Button
	Menu
	Image
	Unknown
)

func (c Context) String() string {
	switch c {
	case Document:
		return "document"
	case Section:
		return "section"
	case Paragraph:
		return "paragraph"
	case Kbd:
		return "kbd"
	case Button:
		return "button"
	case Menu:
		return "menu"
	case Image:
		return "image"
	}
	return "unknown"
}

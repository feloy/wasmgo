package dom

const (
	HyperlinkRelationAlternate  = "alternate"
	HyperlinkRelationAuthor     = "author"
	HyperlinkRelationBookmark   = "bookmark"
	HyperlinkRelationExternal   = "external"
	HyperlinkRelationHelp       = "help"
	HyperlinkRelationLicense    = "license"
	HyperlinkRelationNext       = "next"
	HyperlinkRelationNofollow   = "nofollow"
	HyperlinkRelationNoreferrer = "noreferrer"
	HyperlinkRelationNoopener   = "noopener"
	HyperlinkRelationPrev       = "prev"
	HyperlinkRelationSearch     = "search"
	HyperlinkRelationTag        = "tag"
)

type HyperlinkOptions struct {
	Destination string
	Relation    string
}

// NewHyperlink returns a new "<a>" tag with text content
func NewHyperlink(text string, options HyperlinkOptions) *Tag {
	return &Tag{
		Name:      "a",
		InnerHTML: text,
		Attributes: map[string]string{
			"href": options.Destination,
			"rel":  options.Relation,
		},
	}
}

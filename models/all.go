package models

// PageData represents the data structure passed to templates, it holds the common data for all pages
type PageData struct {
	Title   string
	Header  string
	Content map[string]interface{}
}

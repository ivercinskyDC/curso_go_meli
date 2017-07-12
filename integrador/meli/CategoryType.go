package meli

//Category struct para el json category
type Category struct {
	ID                       string     `json:"id"`
	TotalItemsInThisCategory int        `json:"total_items_in_this_category"`
	ChildrenCategories       []Category `json:"children_categories"`
}

//CategoryNotFound it to parse the response of a notfound
type CategoryNotFound struct {
	Cause   string `json:"cause"`
	Error   string `json:"error"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

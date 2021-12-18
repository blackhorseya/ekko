package todo

type reqID struct {
	ID string `uri:"id" binding:"required"`
}

type reqTitle struct {
	Title string `uri:"title"`
}

type reqStatus struct {
	Status bool `json:"status"`
}

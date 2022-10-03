package todo

type reqID struct {
	ID uint64 `uri:"id" binding:"required"`
}

type reqTitle struct {
	Title string `uri:"title" binding:"required"`
}

type reqStatus struct {
	Status int32 `json:"status"`
}

package todo

type reqID struct {
	ID int64 `uri:"id"`
}

type reqTitle struct {
	Title string `uri:"title" binding:"required"`
}

type reqStatus struct {
	Status int32 `json:"status"`
}

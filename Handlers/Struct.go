package Handlers

type content struct {
	Authlevel     int
	U_id          int
	Posts         []use.Post
	FilteredPosts []use.Post
	LikedPosts    []use.Post
}

type RequestBody struct {
	Pid int `json:"pid"`
}

type CommentRequestBody struct {
	Pid int `json:"pid"`
	Cid int `json:"cid"`
}

type jsonResponse struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	Likes    int    `json:"likes"`
	Dislikes int    `json:"dislikes"`
}
type jsons struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Userl   int    `json:"userl"`
}

type CommentJsons struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Posts   []use.Post `json:"posts"`
}

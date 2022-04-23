package models

type QuizParam struct {
	Question string `json:"question" xml:"question" form:"question"`
}

/*
App
*/

type CreateAppParam struct {
	Name        string `json:"name" xml:"name" form:"name"`
	Type        string `json:"type" xml:"type" form:"type"`
	Description string `json:"description" xml:"description" form:"description"`
	Extend      string `json:"extend" xml:"extend" form:"extend"`
}

type GetAppParam struct {
	AppId string `json:"app_id" xml:"app_id" form:"app_id"`
}

type DeleteAppParam struct {
	AppId string `json:"app_id" xml:"app_id" form:"app_id"`
}

type UpdateAppParam struct {
	AppId       string `json:"app_id" xml:"app_id" form:"app_id"`
	Name        string `json:"name" xml:"name" form:"name"`
	Description string `json:"description" xml:"description" form:"description"`
	Extend      string `json:"extend" xml:"extend" form:"extend"`
}

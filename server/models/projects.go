package models

type Project struct {
	Id                 int      `json:"id"`
	ProjectName        string   `json:"projectName"`
	ProjectDescription string   `json:"projectDescription"`
	Link               string   `json:"link"`
	Technologies       []string `json:"technologies"`
	ImageName          string   `json:"imageName"`
}

package models

type Success struct {
	Message string `json:"message"`
}

type Error struct {
	Message string `json:"message"`
}

type CreateDoc struct {
	Title string `json:"title"`
}

type GetDoc struct {
	Title string `json:"title"`
}

type GetAllDocuments struct {
	Limit  int32  `json:"limit"`
	Page   int32  `json:"page"`
	DocsId string `json:"docs_id"`
}

type UpdateDocument struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	DocsId  string `json:"docs_id"`
}

type DeleteDocument struct {
	Title string `json:"title"`
}

type ShareDocument struct {
	Title          string `json:"title"`
	RecipientEmail string `json:"recipient_email"`
	Permissions    string `json:"permissions"`
	UserId         string `json:"user_id"`
	Id             string `json:"id"`
}

type SearchDocument struct {
	Title  string `json:"title"`
	DocsId string `json:"docs_id"`
}

type GetAllVersions struct {
	Title string `json:"title"`
}

type RestoreVersion struct {
	Title   string `json:"title"`
	Version int32  `json:"version"`
}

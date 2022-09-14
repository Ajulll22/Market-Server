package requests

type QueryProducts struct {
	Id_category string `query:"id_category"`
	Page        int    `query:"page"`
	Limit       int    `query:"limit"`
	Search      string `query:"search"`
}

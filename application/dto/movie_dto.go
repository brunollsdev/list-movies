package dto

type SearchResultDto struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ReleaseDate string `json:"release_date"`
}

type FavoriteMovieRequest struct {
	Id      int    `json:"id"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

type MovieDto struct {
	Id      int    `json:"id"`
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
}

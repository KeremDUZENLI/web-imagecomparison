package app

type VoteDTO struct {
	UserName    string `json:"userName"`
	ImageA      string `json:"imageA"`
	ImageB      string `json:"imageB"`
	ImageWinner string `json:"imageWinner"`
	ImageLoser  string `json:"imageLoser"`
}

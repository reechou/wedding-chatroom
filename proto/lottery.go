package proto

type CreateLotteryReq struct {
	WeddingId    int64  `json:"weddingId"`
	CreateUserId int64  `json:"createUserId"`
	Name         string `json:"name"`
	Num          int64  `json:"num"`
}

type DeleteLotteryReq struct {
	LotteryId int64 `json:"lotteryId"`
}

type UpdateLotteryReq struct {
	LotteryId int64  `json:"lotteryId"`
	Name      string `json:"name"`
	Num       int64  `json:"num"`
}

type GetLotteryListReq struct {
	WeddingId int64 `json:"weddingId"`
}

type StartLotteryReq struct {
	WeddingId uint64 `json:"weddingId"`
	LotteryId int64  `json:"lotteryId"`
}

type EndLotteryReq struct {
	WeddingId uint64 `json:"weddingId"`
	LotteryId int64  `json:"lotteryId"`
}

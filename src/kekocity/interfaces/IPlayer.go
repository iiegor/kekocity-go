package interfaces

type IUser interface {
  GetUsername() string
  GetUserId() int64
  GetCoins() int32
}

package repository

type Bucket string

const (
	AccessTokens Bucket = "access_tokens"
	RefreshToken Bucket = "refresh_token"
)

type TokenRepository interface {
	Save(chatId int64, token string, bucket Bucket) error
	Get(chatId int64, bucket Bucket) (string, error)
	Delete(chatId int64, bucket Bucket) error
}

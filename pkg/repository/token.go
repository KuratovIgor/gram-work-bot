package repository

type Bucket string

const (
	AccessTokens Bucket = "access_tokens"
	RefreshToken Bucket = "refresh_token"
	Resumes      Bucket = "resume"
)

type TokenRepository interface {
	Save(chatID int64, token string, bucket Bucket) error
	Get(chatId int64, bucket Bucket) (string, error)
}

type ResumeRepository interface {
	Save(chatID int64, resumes []any, bucket Bucket) error
	Get(chatID int64, bucket Bucket) (any, error)
}

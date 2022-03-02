package repository

type Secret struct {
	Key   string
	Value string
}

type TwoFactorRepository interface {
	Get(string) (*Secret, error)
}

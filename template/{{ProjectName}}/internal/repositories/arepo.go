package repositories

type ARepo interface {
	ReadAllA() ([]string, error)
}

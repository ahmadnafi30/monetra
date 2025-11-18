package bcrypt

import "golang.org/x/crypto/bcrypt"

type Interface interface {
	GenerateFromPassword(password string) (string, error)
	CompareHashAndPassword(hashedPassword string, password string) error
}

type bcryptImpl struct {
	cost int
}

func Init() Interface {
	return &bcryptImpl{
		cost: bcrypt.DefaultCost,
	}
}

func (b *bcryptImpl) GenerateFromPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (b *bcryptImpl) CompareHashAndPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err
}

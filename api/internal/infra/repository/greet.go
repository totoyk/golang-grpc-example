package repository

type GreetRepository struct{}

func NewGreetRepository() *GreetRepository {
	return &GreetRepository{}
}

func (g *GreetRepository) FindGreet() string {
	return "Hello, "
}

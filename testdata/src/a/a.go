package a

type UserSuccess struct {
	ID int
	Name string
}

type UserFailed struct {
	Id int
	Name string
}

type FromPathSuccess struct {
	URL string
}

type FromPathFailed struct {
	Url string
}
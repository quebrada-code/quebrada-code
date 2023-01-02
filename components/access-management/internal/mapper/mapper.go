package mapper

type IMapper[T, M interface{}] interface {
	ToEntity() M
	ToModel(T) M
}

package db

type IdModelInterface interface {
	GetId() int
}

type IdModel struct {
	Id int
}

func (m IdModel) GetId() int {
	return m.Id
}

func ExtractIds[T IdModelInterface](models []T) []int {
	ids := make([]int, 0, len(models))

	for _, model := range models {
		ids = append(ids, model.GetId())
	}

	return ids
}

func FindById[T IdModelInterface](id int, models []T) *T {
	for _, model := range models {
		if model.GetId() == id {
			return &model
		}
	}
	return nil
}

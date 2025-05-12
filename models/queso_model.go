package models

type Queso struct {
	ID          int    `json:"id"`
	Nombre      string `json:"nombre" validate:"required"`
	Descripcion string `json:"descripcion,omitempty"`
	Precio      int    `json:"precio" validate:"required,gt=0"`
}

/*
	quesos/
		controllers/
			queso_controller.go
		models/
			queso_model.go
		reposuitories/
			queso_repo.go

		routes/
			api.go
		utils/
			database.go
		go.mod
		go.sum
		main.go
*/

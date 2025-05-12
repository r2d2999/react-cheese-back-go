package repositories

import (
	"database/sql"
	"quesos/models"
)


type QuesoRepository struct {
	DB *sql.DB
}

//Método para crear un nuevo queso
func (r *QuesoRepository) CreateNewQueso(queso *models.Queso)(*models.Queso, error){
	query := `
		INSERT INTO quesos (id, nombre, descripcion, precio)
		VALUES ($1, $2, $3, $4)
		RETURNING id, nombre, descripcion, precio
	`;

	err := r.DB.QueryRow(
		query, queso.ID,  queso.Nombre, queso.Descripcion,
		queso.Precio,
	).Scan(
		&queso.ID, &queso.Nombre, &queso.Descripcion, &queso.Precio,
	)

	if err != nil {
		return nil, err
	}

	return queso, nil
}

//repositorio para obtener un solo registro por ID
func (r *QuesoRepository) GetQuesoByID(id int) (*models.Queso, error){
	var queso models.Queso

	err := r.DB.QueryRow("SELECT id, nombre, descripcion, precio FROM quesos WHERE id = $1", id).Scan(&queso.ID, &queso.Nombre, &queso.Descripcion, &queso.Precio)

	if err != nil{
		if err == sql.ErrNoRows{
			return nil, nil
		}

		return nil, err
	}

	return &queso, nil
}
//Repositorio para obtner todos los registros de QUESOS
func (r *QuesoRepository) GetAll() ([]models.Queso, error){
	rows, err := r.DB.Query("SELECT id, nombre, descripcion, precio FROM quesos")

	if err != nil{
		return nil, err
	}
	defer rows.Close()

	var quesos []models.Queso

	for rows.Next(){
		var queso models.Queso
		if err := rows.Scan(&queso.ID, &queso.Nombre, &queso.Descripcion, &queso.Precio); err != nil{
			return nil, err
		}
		quesos = append(quesos, queso)
	}

	if err := rows.Err(); err != nil{
		return nil, err
	}

	return quesos, nil
}
package controllers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"quesos/models"
	"quesos/repositories"
	"strconv"
)

type QuesoController struct {
	Repo *repositories.QuesoRepository
}

func NewQuesoController(db *sql.DB) *QuesoController{
	return &QuesoController{
		Repo: &repositories.QuesoRepository{DB: db},
	}
}

func (c *QuesoController) CreateQuesoCtlr (w http.ResponseWriter, r *http.Request){
	//Se valida que el método sea POST
	if r.Method != http.MethodPost{
		http.Error(w, "Metodo no permitodo", http.StatusMethodNotAllowed)
		return
	}

	//Decodificar el request del body
	var nuevoQueso models.Queso

	err := json.NewDecoder(r.Body).Decode(&nuevoQueso)
	if err != nil{
		http.Error(w, "Formato json no valido", http.StatusBadRequest)
		return
	}

	//Validaciones
	if nuevoQueso.Nombre == ""{
		http.Error(w, "El nombre es requerido", http.StatusBadRequest)
		return
	}

	if nuevoQueso.Precio <= 0 {
		http.Error(w, "El precio debe ser mayor a 0", http.StatusBadRequest)
		return
	}

	quesoCreado, err := c.Repo.CreateNewQueso(&nuevoQueso)
	if err != nil{
		http.Error(w, "Error al crear el queso: " + err.Error(), http.StatusInternalServerError)
		return
	}

	//Enviar el response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(quesoCreado)
}

func(c *QuesoController) GetQuesoByIDCtlr(w http.ResponseWriter, r *http.Request){
	idStr := r.PathValue("id")

	//Se obtiene el parametro
	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w, "ID invalido", http.StatusBadRequest)
		return
	}

	//Se obtiene el registro por ID
	queso, err := c.Repo.GetQuesoByID(id)
	if err != nil{
		http.Error(w, "Error al buscar el queso", http.StatusInternalServerError)
		return
	}

	//verificación del registro
	if queso == nil{
		http.Error(w, "Queso no encontrado", http.StatusNotFound)
		return
	}

	//Devuelve el response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(queso)
}

func (c *QuesoController) GetAllQuesos(w http.ResponseWriter, r *http.Request){
	//verificar el método http
	if r.Method != http.MethodGet{
		http.Error(w, "Método no permitodo...", http.StatusMethodNotAllowed)
		return 
	}

	//Obtener todos los quesos
	quesos, err := c.Repo.GetAll()
	if err != nil{
		http.Error(w, "Erro al obtneer los registros de quesos: " + err.Error(), http.StatusInternalServerError)
		return
	}

	//Configurar el response
	if err := json.NewEncoder(w).Encode(quesos); err != nil {
		http.Error(w, "Erro al codificar la response", http.StatusInternalServerError)
	}
}
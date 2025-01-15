package product

import (
	"fmt"
	"golang-api/service/auth"
	"golang-api/types"
	"golang-api/utils"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
	user  types.UserStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", auth.WithJWTAuth(h.handleGetProducts, h.user)).Methods(http.MethodGet)
	router.HandleFunc("/products", auth.WithJWTAuth(h.handleCreateProduct, h.user)).Methods(http.MethodPost)
	router.HandleFunc("/product/{id:[0-9]+}", h.handleGetProduct).Methods(http.MethodGet)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {

	var payload types.ProductPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	err := h.store.CreateProduct(types.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Image:       payload.Image,
		Price:       payload.Price,
		Quantity:    payload.Quantity,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, nil)
}

func (h *Handler) handleGetProduct(w http.ResponseWriter, r *http.Request) {

	q := mux.Vars(r)
	id, err := strconv.Atoi(q["id"])

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	p, err := h.store.GetDetailProduct(id)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, p)
}

// func handleUpdateProduct(w http.ResponseWriter, r *http.Request) {

// }

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {

	p, err := h.store.GetProducts()

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, p)
}

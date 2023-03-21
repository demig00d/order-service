package http

import (
	"github.com/demig00d/order-service/internal/controller/http/view"
	"github.com/demig00d/order-service/internal/usecase"
	"github.com/demig00d/order-service/pkg/logger"

	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

type orderRoutes struct {
	uc usecase.Order
	l  logger.Interface
}

// @Summary     Show single order
// @Description Show order by ID
// @ID          order
// @Tags  	    order
// @Accept      json
// @Produce     json
// @Success     200 {object} view.OrderView
// @Failure     500 {object} response
// @Router      /order/:id [get]
func (r *orderRoutes) getById(c *gin.Context) {
	orderId := c.Param("id")
	order, err := r.uc.GetById(c.Request.Context(), orderId)
	if err != nil {
		r.l.Error(err, "http - v1 - order")
		errorResponse(c, http.StatusInternalServerError, "database problems")

		return
	}

	var orderView view.OrderView
	err = json.Unmarshal(order.OrderInfo, &orderView)

	if err != nil {
		r.l.Error(err, "http - v1 - order")
		return
	}
	c.HTML(http.StatusOK, "index.html", orderView)
}

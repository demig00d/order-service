package http

import (
	"github.com/demig00d/order-service/internal/entity"
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
// @Success     200 {object} entity.Order
// @Router      /order/:id [get]
func (r *orderRoutes) getById(c *gin.Context) {
	orderId := c.Param("id")
	orderDto, err := r.uc.GetById(c.Request.Context(), orderId)
	if err != nil {
		c.HTML(http.StatusNotFound, "notfound_order.html", nil)
		return
	}

	var order entity.Order
	err = json.Unmarshal(orderDto.OrderInfo, &order)

	if err != nil {
		r.l.Error(err, "http - v1 - order")
		c.HTML(http.StatusInternalServerError, "invalid_data.html", nil)
		return
	}
	c.HTML(http.StatusOK, "index.html", order)
}

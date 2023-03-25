package http

import (
	"errors"

	"github.com/demig00d/order-service/internal/usecase"
	"github.com/demig00d/order-service/pkg/logger"
	"github.com/jackc/pgx/v4"

	"net/http"

	"github.com/gin-gonic/gin"
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
	order, err := r.uc.GetById(c.Request.Context(), orderId)
	if err != nil {
		if errors.Unwrap(err) == pgx.ErrNoRows {
			r.l.Info("http - v1 - order not found")
			c.HTML(http.StatusNotFound, "notfound_order.html", nil)
			return
		}
		r.l.Error(err, "http - v1 - order")
		c.HTML(http.StatusInternalServerError, "invalid_data.html", nil)

		return
	}

	if err != nil {
		return
	}
	c.HTML(http.StatusOK, "index.html", order)
}

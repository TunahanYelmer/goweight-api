package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"weightApi/internal/bluetooth/scanner"
)

func ScanDevicesHandler(c *gin.Context) {
	devices := scanner.ScanForDevices() // []models.BLEDevice
	c.JSON(http.StatusOK, devices)
}
func ConnectDevicesHandler(c *gin.Context) {

}
func ReadDeviceHandler(c *gin.Context) {

}

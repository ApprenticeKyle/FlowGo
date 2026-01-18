package devops

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// StreamLogs handles Server-Sent Events for logs
func (h *DevOpsHandler) StreamLogs(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

	// Subscribe to logs
	logChan := h.devopsService.Broadcaster.Subscribe()
	defer h.devopsService.Broadcaster.Unsubscribe(logChan)

	// Listen for client disconnect
	clientGone := c.Writer.CloseNotify()

	for {
		select {
		case <-clientGone:
			return
		case event := <-logChan:
			eventJSON, err := json.Marshal(event)
			if err != nil {
				continue
			}
			fmt.Fprintf(c.Writer, "data: %s\n\n", eventJSON)
			c.Writer.Flush()
		}
	}
}

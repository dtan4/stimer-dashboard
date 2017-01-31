package systemd

import (
	"strings"
	"time"

	"github.com/coreos/go-systemd/dbus"
)

const (
	timerSuffix   = ".timer"
	timerUnitType = "Timer"
)

// Client represents systemd D-Bus API client.
type Client struct {
	conn *dbus.Conn
}

// Timer represents systemd timer
type Timer struct {
	UnitName        string    `json:"unit_name"`
	LastTriggeredAt time.Time `json:"last_triggered_at"`
	NextTriggerAt   time.Time `json:"next_trigger_at"`
	Result          string    `json:"result"`
}

// NewClient creates new Client object
func NewClient(conn *dbus.Conn) *Client {
	return &Client{
		conn: conn,
	}
}

// NewConn establishes a new connection to D-Bus
func NewConn() (*dbus.Conn, error) {
	return dbus.New()
}

// ListTimers returns installed systemd timers
func (c *Client) ListTimers() ([]*Timer, error) {
	units, err := c.conn.ListUnits()
	if err != nil {
		return []*Timer{}, err
	}

	timers := []*Timer{}

	for _, unit := range units {
		if !strings.HasSuffix(unit.Name, timerSuffix) {
			continue
		}

		timer := &Timer{
			UnitName: unit.Name,
		}

		props, err := c.conn.GetUnitTypeProperties(unit.Name, timerUnitType)
		if err != nil {
			return []*Timer{}, err
		}

		if v, ok := props["LastTriggerUSec"]; ok {
			if lastTriggerUSec, ok := v.(uint64); ok {
				timer.LastTriggeredAt = time.Unix(int64(lastTriggerUSec)/1000/1000, 0)
			}
		}

		if v, ok := props["NextElapseUSecRealtime"]; ok {
			if nextElapseUSecRealtime, ok := v.(uint64); ok {
				timer.NextTriggerAt = time.Unix(int64(nextElapseUSecRealtime)/1000/1000, 0)
			}
		}

		if v, ok := props["Result"]; ok {
			if result, ok := v.(string); ok {
				timer.Result = result
			}
		}

		timers = append(timers, timer)
	}

	return timers, nil
}

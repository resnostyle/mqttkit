// Package hadisc builds Home Assistant MQTT discovery device blocks.
package hadisc

// Device returns a standard HA discovery "device" map.
func Device(identifiers []string, name, manufacturer, model string) map[string]any {
	return map[string]any{
		"identifiers":  identifiers,
		"name":         name,
		"manufacturer": manufacturer,
		"model":        model,
	}
}

package hadisc

import "github.com/resnostyle/mqttkit/mqttpub"

// LatencyReachable returns Latency sensor + Reachable binary_sensor discovery configs
// for a probe/current JSON topic that exposes latency_ms and reachable.
func LatencyReachable(uid, stateTopic string, device map[string]any) []mqttpub.Config {
	latency := map[string]any{
		"name":                  "Latency",
		"unique_id":             uid + "_latency",
		"state_topic":           stateTopic,
		"value_template":        "{{ value_json.latency_ms }}",
		"device":                device,
		"object_id":             uid + "_latency",
		"state_class":           "measurement",
		"unit_of_measurement":   "ms",
		"icon":                  "mdi:speedometer",
		"json_attributes_topic": stateTopic,
	}
	reachable := map[string]any{
		"name":                  "Reachable",
		"unique_id":             uid + "_reachable",
		"state_topic":           stateTopic,
		"value_template":        "{{ 'true' if value_json.reachable else 'false' }}",
		"device":                device,
		"object_id":             uid + "_reachable",
		"device_class":          "connectivity",
		"payload_on":            "true",
		"payload_off":           "false",
		"icon":                  "mdi:lan-connect",
		"json_attributes_topic": stateTopic,
	}
	return []mqttpub.Config{
		{ObjectID: uid + "_latency", Component: "sensor", Payload: latency},
		{ObjectID: uid + "_reachable", Component: "binary_sensor", Payload: reachable},
	}
}

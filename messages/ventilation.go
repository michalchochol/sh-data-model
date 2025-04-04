package messages

type Message struct {
	Topic string
	Unit  string
	Value interface{}
}

// const (
// 	VENT_DATA_ENERGY_USED         = "ventilation/data/EnergyUsed"
// 	VENT_DATA_ENERGY_RECOVERED    = "ventilation/data/EnergyRecovered"
// 	VENT_DATA_EXCHANGE_EFFICIENCY = "venilation/data/ExchangeEfficiency"
// 	VENT_DATA_EXCHAUST_TEMP       = "ventilation/data/ExchaustTemp"
// 	VENT_DATA_EXTERNAL_TEMP       = "ventilation/data/ExternalTemp"
// 	VENT_DATA_HEAT_POW_CONS       = "ventilation/data/HeatingPowerConsumption"
// 	VENT_DATA_HEAT_RECOVERY       = "ventilation/data/HeatRecovery"
// 	VENT_DATA_POWER_CONS          = "ventilation/data/PowerConsumption"
// 	VENT_DATA_SUPPLY_AIR_TEMP     = "ventilation/data/SupplyAirTemp"
// )

var VENT_DATA_ENERGY_USED = Message{Topic: "ventilation/data/EnergyUsed", Unit: "kWh"}
var VENT_DATA_ENERGY_RECOVERED = Message{Topic: "ventilation/data/EnergyRecovered", Unit: "kWh"}
var VENT_DATA_EXCHANGE_EFFICIENCY = Message{Topic: "venilation/data/ExchangeEfficiency", Unit: "%"}
var VENT_DATA_EXCHAUST_TEMP = Message{Topic: "ventilation/data/ExchaustTemp", Unit: "°C"}
var VENT_DATA_EXTERNAL_TEMP = Message{Topic: "ventilation/data/ExternalTemp", Unit: "°C"}
var VENT_DATA_HEAT_POW_CONS = Message{Topic: "ventilation/data/HeatingPowerConsumption", Unit: "W"}
var VENT_DATA_HEAT_RECOVERY = Message{Topic: "ventilation/data/HeatRecovery", Unit: "W"}
var VENT_DATA_POWER_CONS = Message{Topic: "ventilation/data/PowerConsumption", Unit: "W"}
var VENT_DATA_SUPPLY_AIR_TEMP = Message{Topic: "ventilation/data/SupplyAirTemp", Unit: "°C"}

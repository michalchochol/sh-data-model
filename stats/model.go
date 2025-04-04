package stats

import (
	"github.com/michalchochol/sh-data-model/messages"
)

type StatsPoint struct {
	Measurement string
	Tags        map[string]string
	Fields      map[string]interface{}
}

var statsMap = make(map[string]StatsPoint)

func initStatsMap() {
	// ************************
	// VENTILATION
	// ************************
	statsMap[messages.VENT_DATA_ENERGY_USED.Topic] = StatsPoint{
		Measurement: "power",
		Tags:        map[string]string{"device": "ventilation", "name": "EnergyUsed", "unit": messages.VENT_DATA_ENERGY_RECOVERED.Unit},
	}

	statsMap["ventilation/data/EnergyUsed"] = StatsPoint{
		Measurement: "power",
		Tags:        map[string]string{"device": "ventilation", "name": "EnergyUsed", "unit": "kWh"},
	}
	statsMap["ventilation/data/EnergyRecovered"] = StatsPoint{
		Measurement: "power",
		Tags:        map[string]string{"device": "ventilation", "name": "EnergyRecovered", "unit": "kWh"},
	}
	statsMap["ventilation/data/ExchangeEfficiency"] = StatsPoint{
		Measurement: "efficiency",
		Tags:        map[string]string{"device": "ventilation", "name": "ExchangeEfficiency", "unit": "%"},
	}
	statsMap["ventilation/data/ExhaustTemp"] = StatsPoint{
		Measurement: "temperature",
		Tags:        map[string]string{"device": "ventilation", "name": "ExhaustTemp", "unit": "°C"},
	}
	statsMap["ventilation/data/ExternalTemp"] = StatsPoint{
		Measurement: "temperature",
		Tags:        map[string]string{"device": "ventilation", "name": "ExternalTemp", "unit": "°C"},
	}
	statsMap["ventilation/data/HeatingPowerConsumption"] = StatsPoint{
		Measurement: "power",
		Tags:        map[string]string{"device": "ventilation", "name": "HeatingPowerConsumption", "unit": "W"},
	}
	statsMap["ventilation/data/HeatRecovery"] = StatsPoint{
		Measurement: "power",
		Tags:        map[string]string{"device": "ventilation", "name": "HeatRecovery", "unit": "W"},
	}
	statsMap["ventilation/data/PowerConsumption"] = StatsPoint{
		Measurement: "power",
		Tags:        map[string]string{"device": "ventilation", "name": "PowerConsumption", "unit": "W"},
	}
	statsMap["ventilation/data/SupplyAirTemp"] = StatsPoint{
		Measurement: "temperature",
		Tags:        map[string]string{"device": "ventilation", "name": "SupplyAirTemp", "unit": "°C"},
	}

	// ************************
	// SUPLA
	// ************************

	statsMap["supla/8dd35817e857a559a242905fec2dc932/devices/3416/channels/8873/state/phases/1/power_active"] = StatsPoint{
		Measurement: "power",
		Tags:        map[string]string{"device": "supla", "name": "1_PowerActive", "unit": "W"},
	}
	statsMap["supla/8dd35817e857a559a242905fec2dc932/devices/3416/channels/8873/state/phases/2/power_active"] = StatsPoint{
		Measurement: "power",
		Tags:        map[string]string{"device": "supla", "name": "2_PowerActive", "unit": "W"},
	}
	statsMap["supla/8dd35817e857a559a242905fec2dc932/devices/3416/channels/8873/state/phases/3/power_active"] = StatsPoint{
		Measurement: "power",
		Tags:        map[string]string{"device": "supla", "name": "3_PowerActive", "unit": "W"},
	}
	statsMap["supla/8dd35817e857a559a242905fec2dc932/devices/3416/channels/8873/state/phases/1/voltage"] = StatsPoint{
		Measurement: "voltage",
		Tags:        map[string]string{"device": "supla", "name": "1_Voltage", "unit": "V"},
	}
	statsMap["supla/8dd35817e857a559a242905fec2dc932/devices/3416/channels/8873/state/phases/2/voltage"] = StatsPoint{
		Measurement: "voltage",
		Tags:        map[string]string{"device": "supla", "name": "2_Voltage", "unit": "V"},
	}
	statsMap["supla/8dd35817e857a559a242905fec2dc932/devices/3416/channels/8873/state/phases/3/voltage"] = StatsPoint{
		Measurement: "voltage",
		Tags:        map[string]string{"device": "supla", "name": "3_Voltage", "unit": "V"},
	}

	// ************************
	// SOLAR
	// ************************

	statsMap["solar/ac/power"] = StatsPoint{
		Measurement: "power",
		Tags:        map[string]string{"device": "solar", "name": "AC_Power", "unit": "W"},
	}
	statsMap["solar/dc/power"] = StatsPoint{
		Measurement: "power",
		Tags:        map[string]string{"device": "solar", "name": "DC_Power", "unit": "W"},
	}

	// !!! overvoltage detection

	// ************************
	// WEATHER.COM
	// ************************

	statsMap["weather/current/temperature"] = StatsPoint{
		Measurement: "temperature",
		Tags:        map[string]string{"device": "weather", "name": "Temperature", "unit": "°C"},
	}
	statsMap["weather/current/apparent_temperature"] = StatsPoint{
		Measurement: "temperature",
		Tags:        map[string]string{"device": "weather", "name": "ApparentTemperature", "unit": "°C"},
	}
	statsMap["weather/current/temperature"] = StatsPoint{
		Measurement: "humidity",
		Tags:        map[string]string{"device": "weather", "name": "Humidity", "unit": "%"},
	}
	statsMap["weather/current/cloud_cover"] = StatsPoint{
		Measurement: "cover",
		Tags:        map[string]string{"device": "weather", "name": "CloudCover", "unit": "%"},
	}
	statsMap["weather/current/is_day"] = StatsPoint{
		Measurement: "status",
		Tags:        map[string]string{"device": "weather", "name": "IsDay", "unit": "bool"},
	}
	statsMap["weather/current/wind_speed"] = StatsPoint{
		Measurement: "wind",
		Tags:        map[string]string{"device": "weather", "name": "WindSpeed", "unit": "m/s"},
	}
	statsMap["weather/current/wind_gusts"] = StatsPoint{
		Measurement: "wind",
		Tags:        map[string]string{"device": "weather", "name": "WindGusts", "unit": "m/s"},
	}
	statsMap["weather/current/wind_direction"] = StatsPoint{
		Measurement: "wind",
		Tags:        map[string]string{"device": "weather", "name": "WindDirection", "unit": "angle"},
	}
	statsMap["weather/daily/uv_index_max"] = StatsPoint{
		Measurement: "uv_index",
		Tags:        map[string]string{"device": "weather", "name": "UVIndexMax", "unit": ""},
	}
	statsMap["weather/daily/uv_index_clear_sky_max"] = StatsPoint{
		Measurement: "uv_index",
		Tags:        map[string]string{"device": "weather", "name": "UVIndexClearSkyMax", "unit": ""},
	}

	// ************************
	// HEATER
	// ************************

	statsMap["ebusd/e7c/HeatCurve"] = StatsPoint{
		Measurement: "curve",
		Tags:        map[string]string{"device": "heater", "name": "HeatCurve", "unit": "num"},
	}
	statsMap["ebusd/e7c/RoomTemp"] = StatsPoint{
		Measurement: "temperature",
		Tags:        map[string]string{"device": "heater", "name": "RoomTemp", "unit": "°C"},
	}
	statsMap["ebusd/bai/FlowTempDesired"] = StatsPoint{
		Measurement: "temperature",
		Tags:        map[string]string{"device": "heater", "name": "FlowTempDesired", "unit": "°C"},
	}
	statsMap["ebusd/bai/FlowTemp"] = StatsPoint{
		Measurement: "temperature",
		Tags:        map[string]string{"device": "heater", "name": "FlowTemp", "unit": "°C"},
	}
	statsMap["ebusd/bai/OutdoorstempSensor"] = StatsPoint{
		Measurement: "temperature",
		Tags:        map[string]string{"device": "heater", "name": "OutdoorstempSensor", "unit": "°C"},
	}
	statsMap["ebusd/bai/ReturnTemp"] = StatsPoint{
		Measurement: "temperature",
		Tags:        map[string]string{"device": "heater", "name": "ReturnTemp", "unit": "°C"},
	}
	statsMap["ebusd/bai/StorageTemp"] = StatsPoint{
		Measurement: "temperature",
		Tags:        map[string]string{"device": "heater", "name": "StorageTemp", "unit": "°C"},
	}
	statsMap["ebusd/bai/PrEnergySumHc1"] = StatsPoint{
		Measurement: "power",
		Tags:        map[string]string{"device": "heater", "name": "PrEnergySumHc1", "unit": "num"},
	}
	statsMap["ebusd/bai/PrEnergySumHwc1"] = StatsPoint{
		Measurement: "power",
		Tags:        map[string]string{"device": "heater", "name": "PrEnergySumHwc1", "unit": "num"},
	}
	statsMap["ebusd/bai/Flame"] = StatsPoint{
		Measurement: "status",
		Tags:        map[string]string{"device": "heater", "name": "Flame", "unit": "bool"},
	}
	statsMap["ebusd/bai/StatusCirPump"] = StatsPoint{
		Measurement: "status",
		Tags:        map[string]string{"device": "heater", "name": "StatusCirPump", "unit": "bool"},
	}
	statsMap["ebusd/bai/Statenumber"] = StatsPoint{
		Measurement: "status",
		Tags:        map[string]string{"device": "heater", "name": "Statenumber", "unit": "num"},
	}
	statsMap["ebusd/bai/DisplayMode"] = StatsPoint{
		Measurement: "status",
		Tags:        map[string]string{"device": "heater", "name": "DisplayMode", "unit": "num"},
	}

}

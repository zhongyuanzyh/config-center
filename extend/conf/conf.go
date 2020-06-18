package conf

type Log struct {
	Logging Levels `json:"logging"`
}

type Levels struct {
	Level Ios `json:"level"`
}

type Ios struct {
	Io Swaggers `json:"io"`
}
type Swaggers struct {
	Swagger Modelss `json:"swagger"`
}

type Modelss struct {
	Models Parameterss `json:"models"`
}
type Parameterss struct {
	Parameters AbstractSerializableParameters `json:"parameters"`
}
type AbstractSerializableParameters struct {
	AbstractSerializableParameter string `json:"AbstractSerializableParameter"`
}

package conf

//Log相关的字段
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

//mybatis相关的字段
type MyBatis struct {
	MybatisPlus GlobalConfigs `json:"mybatis-plus"`
}
type GlobalConfigs struct {
	GlobalConfig    DbConfigs `json:"global-config"`
	MapperLocations string    `json:"mapper-locations"`
}
type DbConfigs struct {
	DbConfig DbConfigDetail `json:"db-config"`
}
type DbConfigDetail struct {
	LogicDeleteValue    int `json:"logic-delete-value"`
	LogicNotDeleteValue int `json:"logic-not-delete-value"`
}

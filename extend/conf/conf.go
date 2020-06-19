package conf

type ConfigYaml struct {
	Logging     Levels        `json:"logging"`
	MybatisPlus GlobalConfigs `json:"mybatis-plus"`
	Obs         Huaweis       `json:"obs"`
	Server      ServerDetail  `json:"server"`
	Spring      DataSources   `json:"spring"`
	User        Roles         `json:"user"`
	Xinchao     Premises      `json:"xinchao"`
	Xxl         Jobs          `json:"xxl"`
	KuYun       Dmps          `json:"KuYun"`
}

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

//obs相关的字段
type ObsStorage struct {
	Obs Huaweis `json:"obs"`
}
type Huaweis struct {
	Huawei HuaweiDetail `json:"huawei"`
}
type HuaweiDetail struct {
	AccessKeyId     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	BaseUrl         string `json:"baseUrl"`
	Bucket          string `json:"bucket"`
	EndPoint        string `json:"endpoint"`
}

//server相关的字段
type Servers struct {
	Server ServerDetail `json:"server"`
}
type ServerDetail struct {
	Port int `json:"port"`
}

//spring相关的字段
type Springs struct {
	Spring DataSources `json:"spring"`
}
type DataSources struct {
	DataSource DataSourceDetail `json:"datasource"`
	Kafka      KafkaDetail      `json:"kafka"`
	Redis      RedisDetail      `json:"redis"`
	Servlet    ServletDetail    `json:"servlet"`
}
type DataSourceDetail struct {
	DriverClassName string        `json:"driver-class-name"`
	Druid           DruidDetail   `json:"druid"`
	Dynamic         DynamicDetail `json:"dynamic"`
	Type            string        `json:"type"`
}
type DruidDetail struct {
	MaxActive int `json:"max-active"`
}
type DynamicDetail struct {
	DataSource DynamicDataSourceDetail `json:"datasource"`
	Health     bool                    `json:"health"`
	Primary    string                  `json:"primary"`
}
type DynamicDataSourceDetail struct {
	Master   MasterDetail   `json:"master"`
	Postgres PostgresDetail `json:"postgres"`
}
type MasterDetail struct {
	DriverClassName string `json:"driver-class-name"`
	Password        string `json:"password"`
	Type            string `json:"type"`
	Url             string `json:"url"`
	Username        string `json:"username"`
}
type PostgresDetail struct {
	DriverClassName string `json:"driver-class-name"`
	Password        string `json:"password"`
	Type            string `json:"type"`
	Url             string `json:"url"`
	Username        string `json:"username"`
}
type KafkaDetail struct {
	Producer ProducerDetail `json:"producer"`
	Topic    TopicDetail    `json:"topic"`
}
type ProducerDetail struct {
	BatchSize       int    `json:"batch-size"`
	Retries         int    `json:"retries"`
	BufferMemory    int    `json:"buffer-memory"`
	Servers         string `json:"servers"`
	Linger          int    `json:"linger"`
	ValueSerializer string `json:"value-serializer"`
	KeySerializer   string `json:"key-serializer"`
}
type TopicDetail struct {
	Biz BizDetail `json:"biz"`
}
type BizDetail struct {
	Msg string `json:"msg"`
}
type RedisDetail struct {
	Database int           `json:"database"`
	Host     string        `json:"host"`
	Lettuce  LettuceDetail `json:"lettuce"`
	Password string        `json:"password"`
	Port     int           `json:"port"`
	Timeout  int           `json:"timeout"`
}
type LettuceDetail struct {
	Pool PoolDetail `json:"pool"`
}
type PoolDetail struct {
	MaxActive int `json:"max-active"`
	MaxIdle   int `json:"max-idle"`
	MaxWait   int `json:"max-wait"`
	MinIdle   int `json:"min-idle"`
}
type ServletDetail struct {
	Multipart MultipartDetail `json:"multipart"`
}
type MultipartDetail struct {
	Enabled           bool `json:"enabled"`
	FileSizeThreshold int  `json:"file-size-threshold"`
	MaxFileSize       int  `json:"max-file-size"`
	MaxRequestSize    int  `json:"max-request-size"`
}

//user相关的字段
type Users struct {
	User Roles `json:"user"`
}
type Roles struct {
	Role RoleDetail `json:"role"`
}
type RoleDetail struct {
	Audit string `json:"audit"`
}

//xinchao相关的字段
type Xinchaos struct {
	Xinchao Premises `json:"xinchao"`
}
type Premises struct {
	Premise PremiseDetail `json:"premise"`
}
type PremiseDetail struct {
	Recommend RecommendDetail `json:"recommend"`
}
type RecommendDetail struct {
	MinDistance int `json:"minDistance"`
	Radius      int `json:"radius"`
}

//Xxl相关的字段
type Xxls struct {
	Xxl Jobs `json:"xxl"`
}
type Jobs struct {
	Job JobDetail `json:"job"`
}
type JobDetail struct {
	AccessToken string         `json:"accessToken"`
	Admin       AdminDetail    `json:"admin"`
	Executor    ExecutorDetail `json:"executor"`
}
type AdminDetail struct {
	Addresses string `json:"addresses"`
}
type ExecutorDetail struct {
	Appname          string `json:"appname"`
	Ip               string `json:"ip"`
	Logpath          string `json:"logpath"`
	Logretentiondays int    `json:"logretentiondays"`
	Port             int    `json:"port"`
}

//KuYun相关的字段
type KuYuns struct {
	KuYun Dmps `json:"KuYun"`
}
type Dmps struct {
	Dmp DmpDetail `json:"dmp"`
}
type DmpDetail struct {
	Interface InterfaceDetail `json:"interface"`
}
type InterfaceDetail struct {
	Hot HotDetail `json:"hot"`
}
type HotDetail struct {
	FileUrl    string `json:"fileUrl"`
	NewFileUrl string `json:"newFileUrl"`
}

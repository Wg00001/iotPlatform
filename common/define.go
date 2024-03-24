package common

var (
	//DSN   = os.Getenv("MysqlDSN")
	DSN             = "root:03719@tcp(127.0.0.1:3306)"
	EmqxAddr        = "http://192.168.1.8:18083/api/v5"
	EmqxKey         = "1f9c5b734fe27865"
	EmqxSecret      = "lV9C2iefOp9Cr9BeiB5rr3N9CBolJjKk3HruhqEpHQxsuVD"
	JwtAccessSecret = "iot-platform"
	JwtAccessExpire = int64(36000)
)

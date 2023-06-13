// @program:     rpc
// @file:        dao.go
// @author:      ugug
// @create:      2023-06-12 01:50
// @description:

package dao

// ScanDB 定期扫描数据库查看是否和本地服务相符
func ScanDB() {
	var err bool = false
	//err:=Scan()
	UpdateDB(err)
}

// UpdateDB 数据库和本地服务不相符时，根据更新日志更新本地或服务器
func UpdateDB(bool) {

}

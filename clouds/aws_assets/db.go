package aws_assets

// import (
// 	"RECON/config"
// 	"database/sql"
// 	"log"
// 	"strings"

// 	_ "github.com/go-sql-driver/mysql"
// )

// type InstanceWhitelistInfo struct {
// 	Servername     string
// 	Serveraddress  string
// 	Serverport     string
// 	Serverprotocol string
// 	Serverexplain  string
// }

// func InstanceWhitelist() []InstanceWhitelistInfo {
// 	db, err := sql.Open("mysql", config.MysqlUrl)
// 	if err != nil {
// 		log.Printf("Database connection error: %v", err)
// 		return nil
// 	}
// 	defer db.Close()

// 	// 测试数据库连接
// 	if err := db.Ping(); err != nil {
// 		log.Printf("Database ping failed: %v", err)
// 		return nil
// 	}
// 	log.Println("Database connected successfully")

// 	// Select all columns
// 	rows, err := db.Query("SELECT * FROM rulesmanage_serverlist") // Querying all columns
// 	if err != nil {
// 		log.Printf("Query failed: %v", err)
// 		return nil
// 	}
// 	defer rows.Close()

// 	var instanceList []InstanceWhitelistInfo
// 	for rows.Next() {
// 		var instanceInfo InstanceWhitelistInfo // Use the struct directly for scanning
// 		var id int                             // Add variable for the 6th column (assuming it's an int ID)

// 		// Provide pointers to ALL 6 columns in the correct order
// 		// Scan directly into struct fields where possible, plus the extra 'id'
// 		if err := rows.Scan(
// 			&id, // Scan the ID column first (assuming it's the first column)
// 			&instanceInfo.Servername,
// 			&instanceInfo.Serveraddress,
// 			&instanceInfo.Serverport,
// 			&instanceInfo.Serverprotocol,
// 			&instanceInfo.Serverexplain,
// 		); err != nil {
// 			log.Printf("Row scan error: %v", err)
// 			continue // Skip this row if scanning fails
// 		}

// 		// Trim leading/trailing whitespace and newlines from string fields
// 		instanceInfo.Servername = strings.TrimSpace(instanceInfo.Servername)
// 		instanceInfo.Serveraddress = strings.TrimSpace(instanceInfo.Serveraddress)
// 		instanceInfo.Serverport = strings.TrimSpace(instanceInfo.Serverport)
// 		instanceInfo.Serverprotocol = strings.TrimSpace(instanceInfo.Serverprotocol)
// 		instanceInfo.Serverexplain = strings.TrimSpace(instanceInfo.Serverexplain)

// 		// Append the cleaned struct to the list
// 		instanceList = append(instanceList, instanceInfo)

// 		// Example: Log scanned data
// 		// log.Printf("Scanned row: ID=%d, Name=%s, Address=%s, Port=%s, Protocol=%s, Explain=%s",
// 		// 	id, instanceInfo.Servername, instanceInfo.Serveraddress, instanceInfo.Serverport, instanceInfo.Serverprotocol, instanceInfo.Serverexplain)
// 	}

// 	if err = rows.Err(); err != nil {
// 		log.Printf("Error iterating rows: %v", err)
// 	}
// 	log.Printf("instanceList: %v", instanceList)
// 	log.Printf("Found %d instances in whitelist", len(instanceList))
// 	return instanceList // Return the list of structs
// }

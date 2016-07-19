package model

import (
	"log"
	"database/sql"
	"domain/entities"
	
)



func  NewCityDAO(dbconnin *Database) *CityDAO {
	citydao := new(CityDAO)
	citydao.dbconn  = *dbconnin
	return citydao
}	
	
	
type CityDAO struct {
	dbconn Database
}


func (dw *CityDAO) Insert(){

}


func (dw *CityDAO) Delete(cityid string){
	
}

func (dw *CityDAO) GetByCity(cityid string) ([]entities.City) {
	
	 rows, err := dw.dbconn.mydbconn.Query("SELECT name, latitude, longitude FROM city WHERE name = ?;", cityid)
	 
	 if err != nil {
		 	println( "error calling query")
			log.Fatal(err)
	 }
	 
	 defer rows.Close()
	 newrows := createEntities(rows)
	 rows.Close()
	 return newrows
}


func (dw *CityDAO) GetAllCities()([]entities.City){
	
	 rows, err := dw.dbconn.mydbconn.Query("SELECT * FROM city;")
	 
	 if err != nil {
		 	println( "error calling query")
			log.Fatal(err)
	 }
	 
	 defer rows.Close()
	 newrows := createEntities(rows)
	 rows.Close()
	 return newrows
}





func  createEntities(rows *sql.Rows) ([]entities.City){
	
	data :=  make([]entities.City,0) 
	for rows.Next(){
		
		city := entities.NewCity()
		rows.Scan(&city.Name, &city.Latitude, &city.Longitude)			
		data = append(data, *city)
	}
	
	if(len(data) < 1){
		log.Output(1, "no data returned")
		println("error reading from db")
	}
	  
	return data
}

func (dw *CityDAO) Update(){
	
}




package service

import (
	"database/sql"
	"fmt"
	"go-and-gin/models"
	"log"
)

type playerService struct {
	DB *sql.DB
}
type PlayerService interface {
	InsertPlayer(Player models.Player) (int64, error)
	GetAllPlayers() ([]models.Player, error)
	GetPlayer(id int64) (models.Player, error)
	UpdatePlayer(id int64, Player models.Player) (int64, error)
	DeletePlayer(id int64) (int64, error)
}

func NewPLayerService(db *sql.DB) PlayerService {
	return &playerService{DB: db}
}
func (p *playerService) InsertPlayer(Player models.Player) (int64, error) {

	// returning Playerid will return the id of the inserted Player
	sqlStatement := `INSERT INTO Players (name, location, age) VALUES ($1, $2, $3) RETURNING Playerid`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// Scan function will save the insert id in the id
	err := p.DB.QueryRow(sqlStatement, Player.Name, Player.Location, Player.Age).Scan(&id)

	if err != nil {
		return 0, err
	}

	fmt.Printf("Inserted a single record %v", id)

	// return the inserted id
	return id, nil
}

// get one Player from the p.DB by its Playerid
func (p *playerService) GetPlayer(id int64) (models.Player, error) {
	// create the postgres p.DB connection

	//defer p.DB.Close()

	// create a Player of models.Player type
	var Player models.Player

	// create the select sql query
	sqlStatement := `SELECT * FROM Players WHERE Playerid=$1`

	// execute the sql statement
	row := p.DB.QueryRow(sqlStatement, id)

	// unmarshal the row object to Player
	err := row.Scan(&Player.ID, &Player.Name, &Player.Age, &Player.Location)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return Player, nil
	case nil:
		return Player, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}

	// return empty Player on error
	return Player, err
}

// get one Player from the p.DB by its Playerid
func (p *playerService) GetAllPlayers() ([]models.Player, error) {

	// close the p.DB connection
	//defer p.DB.Close()

	var Players []models.Player

	// create the select sql query
	sqlStatement := `SELECT * FROM Players`

	// execute the sql statement
	rows, err := p.DB.Query(sqlStatement)

	if err != nil {
		return Players, err
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var Player models.Player

		// unmarshal the row object to Player
		err = rows.Scan(&Player.ID, &Player.Name, &Player.Location, &Player.Age)

		if err != nil {
			return Players, err
		}

		// append the Player in the Players slice
		Players = append(Players, Player)

	}

	// return empty Player on error
	return Players, err
}

// update Player in the p.DB
func (p *playerService) UpdatePlayer(id int64, Player models.Player) (int64, error) {

	// close the p.DB connection
	//defer p.DB.Close()

	// create the update sql query
	sqlStatement := `UPDATE Players SET name=$2, location=$3, age=$4 WHERE Playerid=$1`

	// execute the sql statement
	res, err := p.DB.Exec(sqlStatement, id, Player.Name, Player.Location, Player.Age)

	if err != nil {
		return 0, nil
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected, nil
}

// delete Player in the p.DB
func (p *playerService) DeletePlayer(id int64) (int64, error) {

	// create the postgres p.DB connection
	//defer p.DB.Close()

	// create the delete sql query
	sqlStatement := `DELETE FROM Players WHERE Playerid=$1`

	// execute the sql statement
	res, err := p.DB.Exec(sqlStatement, id)

	if err != nil {
		return -1, err
	}
	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected, nil
}

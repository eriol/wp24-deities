package database // import "github.com/eriol/wp24-deities/database"

import (
	"fmt"
	"math/rand"
	"time"
)

type Deity struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Description string `json:"description"`
}

type OlympianInfluence struct {
	SportId   string  `json:"sport_id"`
	Influence float32 `json:"influence"`
	Name      string  `json:"name"`
}

func GetDeities() ([]Deity, error) {
	query := `
    SELECT
        deity_id,
        name,
        gender,
        description
    FROM
        deities
    ORDER BY
        name;`

	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}

	deities := make([]Deity, 0)
	for rows.Next() {
		deity := Deity{}

		err = rows.Scan(
			&deity.Id,
			&deity.Name,
			&deity.Gender,
			&deity.Description,
		)
		if err != nil {
			return nil, err
		}

		deities = append(deities, deity)
	}

	fmt.Println("Deities", deities)

	return deities, nil
}

func GetDeity(id string) (Deity, error) {
	query := `
    SELECT
        deity_id,
        name,
        gender,
        description
    FROM
        deities
    WHERE
        deity_id = ?;`

	deity := Deity{}

	if err := database.QueryRow(query, id).Scan(
		&deity.Id,
		&deity.Name,
		&deity.Gender,
		&deity.Description,
	); err != nil {
		return deity, err
	}

	return deity, nil
}

func GetRandomDeity() (Deity, error) {
	query := `
    SELECT
        deity_id,
        name,
        gender,
        description
    FROM
        deities
    ORDER BY RANDOM()
    LIMIT 1;`

	deity := Deity{}

	if err := database.QueryRow(query).Scan(
		&deity.Id,
		&deity.Name,
		&deity.Gender,
		&deity.Description,
	); err != nil {
		return deity, err
	}

	return deity, nil
}

func GetDeityInfluence(id string) ([]OlympianInfluence, error) {
	query := `
    SELECT
        olympian_influence.sport_id,
        olympian_influence.influence,
        sports.name
    FROM
        olympian_influence
    INNER JOIN
        sports
    ON
        olympian_influence.sport_id = sports.sport_id
    WHERE
        olympian_influence.deity_id = ?;`

	rows, err := database.Query(query, id)
	if err != nil {
		return nil, err
	}

	influences := make([]OlympianInfluence, 0)

	for rows.Next() {
		olympianInfluence := OlympianInfluence{}
		err = rows.Scan(
			&olympianInfluence.SportId,
			&olympianInfluence.Influence,
			&olympianInfluence.Name,
		)
		if err != nil {
			return nil, err
		}

		influences = append(influences, olympianInfluence)
	}

	return influences, nil
}

// Return Eris influence on sports.
// As goddess of discord Eris will return always a negative influence on
// all sports but random to make people more concerned.
func GetErisInfluence() ([]OlympianInfluence, error) {
	query := `
    SELECT
        sport_id,
        name
    FROM
        sports;`

	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}

	influences := make([]OlympianInfluence, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for rows.Next() {
		olympianInfluence := OlympianInfluence{Influence: r.Float32()}
		err = rows.Scan(
			&olympianInfluence.SportId,
			&olympianInfluence.Name,
		)
		if err != nil {
			return nil, err
		}

		influences = append(influences, olympianInfluence)
	}

	return influences, nil
}

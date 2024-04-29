package database // import "github.com/eriol/wp24-deities/database"

import "fmt"

type Deity struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Gender      string `json:"gender"`
	Description string `json:"description"`
}

type OlympianInfluence struct {
	SportId   string  `json:"sport_id"`
	Influence float32 `json:"influence"`
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
        sport_id,
        influence
    FROM
        olympian_influence
    WHERE
        deity_id = ?;`

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
		)
		if err != nil {
			return nil, err
		}

		influences = append(influences, olympianInfluence)
	}

	return influences, nil
}

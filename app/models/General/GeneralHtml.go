package General

import (
	"firstRest/database"
	"log"
	"sync"
	"time"
)

type Html struct {
	Html string
}

var (
	htmlCache   Html
	cacheLoaded bool
	cacheMu     sync.Mutex
	cacheTTL    = time.Minute // Время жизни кэша
	cacheTime   time.Time
)

func GetFirst() (Html, error) {
	prices, err := database.Select(`SELECT html FROM general_html LIMIT 1;`)

	if err != nil {
		log.Println("Error scanning row: query", err)
		return Html{}, err
	}

	var pricesResult Html
	for prices.Next() {
		var price Html
		if err := prices.Scan(&price.Html); err != nil {
			log.Println("Error scanning row: parse", err)
			return Html{}, err
		}
		pricesResult = price
	}

	return pricesResult, nil
}

func GetFromMemory() (Html, error) {
	cacheMu.Lock()
	defer cacheMu.Unlock()

	if !cacheLoaded || time.Since(cacheTime) > cacheTTL {
		var err error
		htmlCache, err = GetFirst()
		if err != nil {
			return Html{}, err
		}
		cacheLoaded = true
		cacheTime = time.Now()
	}

	return htmlCache, nil
}

func (p *Html) Save() error {
	connection, err := database.GetDBInstance()

	if err != nil {
		return err
	}

	_, err = connection.Db.Exec(`DELETE FROM general_html WHERE true;`)
	if err != nil {
		return err
	}

	_, err = connection.Db.Exec(`INSERT INTO general_html (html) VALUES ($1);`, p.Html)

	if err != nil {
		return err
	}
	return nil
}

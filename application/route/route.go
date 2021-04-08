package route

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
)

/*
Detalhe importante:
  Como trabalharemos com JSON, podemos usar tags
    `json: "tag"`
que o Go irá converter-lo dessa forma
*/

// Definindo a estrutura do tipo Route
// Contém as informações para um novo pedido
type Route struct {
	ID       string `json:"routeId"`
	ClientID string `json:"clientID"`
	// Ao invés de usar string, podemos usar outra estrutura
	// Positions string
	Positions []Position `json:"position"`
}

// Definindo a estrutura do tipo Position.
// Representa a latitude e longitude do pedido.
type Position struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

// Definindo a estrutura do tipo PartialRoutePosition
// Representa a resposta em tempo real do pedido
type PartialRoutePosition struct {
	ID       string    `json:"routeID"`
	ClientID string    `json:"clientID"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}

// "LoadPositions" carrega de um ".txt" as posições e passa para o tipo Position por meio de ponteiro
func (r *Route) LoadPositions() error {
	if r.ID == "" {
		return errors.New("route ID not informed")
	}
	f, err := os.Open("destinations/" + r.ID + ".txt")
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ",")
		lat, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			return nil
		}
		long, err := strconv.ParseFloat(data[1], 64)
		if err != nil {
			return nil
		}
		r.Positions = append(r.Positions, Position{
			Lat:  lat,
			Long: long,
		})
	}
	return nil
}

// ExportJsonPositions generates a slice of string in Json using PartialRoutePosition struct
// "ExportJsonPositions" gera uma lista de strings em JSON e passa para o tipo PartialRoutePosition por meio de ponteiro
// O kafka lê JSON.
func (r *Route) ExportJsonPositions() ([]string, error) {
	var route PartialRoutePosition
	var result []string
	total := len(r.Positions)

	for k, v := range r.Positions {
		route.ID = r.ID
		route.ClientID = r.ClientID
		route.Position = []float64{v.Lat, v.Long}
		route.Finished = false
		if total-1 == k {
			route.ID = r.ID
			route.ClientID = r.ClientID
			route.Position = []float64{v.Lat, v.Long}
			if total-1 == k {
				route.Finished = true
			}
			jsonRoute, err := json.Marshal(route)
			if err != nil {
				return nil, err
			}
			result = append(result, string(jsonRoute))
		}
	}
	return result, nil
}

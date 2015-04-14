package bloomsource

import (
	"io/ioutil"
	"github.com/gocodo/bloomdb"
	"gopkg.in/yaml.v2"
)

func IndexDrop() error {
	bdb := bloomdb.CreateDB()

	file, err := ioutil.ReadFile("searchmapping.yaml")
	if err != nil {
		return err
	}

	mappings := []SearchSource{}
	err = yaml.Unmarshal(file, &mappings)
	if err != nil {
		return err
	}

	conn, err := bdb.SqlConnection()
	if err != nil {
		return err
	}

	searchConn := bdb.SearchConnection()

	for _, source := range mappings {
		typeName := source.Name
		_, err = searchConn.DeleteMapping(source.Name, "main")
		if err != nil {
			return err
		}

		_, err = conn.Exec("DELETE FROM search_types WHERE name = $1", typeName)
		if err != nil {
			return err
		}
	}

	return nil
}
package client

import (
	"log"
	"testing"

	"github.com/gobuffalo/pop/v6"
	"github.com/stretchr/testify/require"
)

var db *pop.Connection

func init() {
	pop.Debug = true
	pop.AddLookupPaths("../../")

	if err := pop.LoadConfigFile(); err != nil {
		log.Panic(err)
	}

	var err error
	db, err = pop.Connect("test")
	if err != nil {
		log.Panic(err)
	}
}

func transaction(fn func(tx *pop.Connection)) {
	err := db.Rollback(func(tx *pop.Connection) {
		fn(tx)
	})
	if err != nil {
		log.Fatal(err)
	}
}

func Test_Create_Client(t *testing.T) {
	if db == nil {
		t.Skip("skipping create test")
	}

	transaction(func(tx *pop.Connection) {
		r := require.New(t)

		client := Client{
			FirstName: "Edwin",
		}

		err := tx.Create(&client)
		r.NoError(err)

		clients := Clients{}
		err = tx.All(&clients)
		r.NoError(err)

		r.Equal(1, len(clients))
	})
}

// func Test_List_Client(t *testing.T) {
// 	if db == nil {
// 		t.Skip("skipping create test")
// 	}

// 	transaction(func(tx *pop.Connection) {
// 		r := require.New(t)
// 		r.Fail("Que muestre todas las mascotas/clientes y su información")
// 	})
// }

// func Test_Find_Client(t *testing.T) {
// 	if db == nil {
// 		t.Skip("skipping create test")
// 	}

// 	transaction(func(tx *pop.Connection) {
// 		r := require.New(t)
// 		r.Fail(`Animal. Ejemplo: mostrar las mascotas que sean 'perros' o mostrar las que sean 'gatos'
// 		Precio (Hasta X monto inclusive). Ejemplo :mostrar las mascotas que estén por debajo de $700.000.
// 		Edad (meses o años): mostrar las mascotas que tengan X meses o Y años (es responsabilidad del usuario especificar si es mes o año)
// 		`)
// 	})
// }

// func Test_Destroy_Client(t *testing.T) {
// 	if db == nil {
// 		t.Skip("skipping create test")
// 	}

// 	transaction(func(tx *pop.Connection) {
// 		r := require.New(t)
// 		r.Fail("Que se puedan remover o eliminar entidades")
// 	})
// }

// func Test_Update_Client(t *testing.T) {
// 	if db == nil {
// 		t.Skip("skipping create test")
// 	}

// 	transaction(func(tx *pop.Connection) {
// 		r := require.New(t)
// 		r.Fail("Que se puedan actualizar entidades")
// 	})
// }

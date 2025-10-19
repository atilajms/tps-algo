package diccionario_test

import (
	"testing"

	"tdas/diccionario"

	"github.com/stretchr/testify/require"
)

// Funcion aux para no repetir codigo
func crearABBEnteros() diccionario.DiccionarioOrdenado[int, int] {
	abb := diccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	abb.Guardar(10, 100)
	abb.Guardar(5, 50)
	abb.Guardar(15, 150)
	abb.Guardar(3, 30)
	abb.Guardar(7, 70)
	abb.Guardar(12, 120)
	abb.Guardar(20, 200)
	return abb
}

func TestIterarSinCorte(t *testing.T) {
	abb := crearABBEnteros()
	suma := 0
	abb.Iterar(func(clave, dato int) bool {
		suma += dato
		return true
	})
	require.Equal(t, 720, suma)
}

func TestIterarConCorte(t *testing.T) {
	abb := crearABBEnteros()
	suma := 0
	abb.Iterar(func(clave, dato int) bool {
		suma += dato
		return suma < 200
	})
	require.True(t, suma >= 200)
	require.True(t, suma < 670)
}

func TestIterarRangoSinLimites(t *testing.T) {
	abb := crearABBEnteros()
	suma := 0
	abb.IterarRango(nil, nil, func(clave, dato int) bool {
		suma += dato
		return true
	})
	require.Equal(t, 720, suma)
}

func TestIterarRangoConLimites(t *testing.T) {
	abb := crearABBEnteros()
	suma := 0
	desde := 6
	hasta := 15
	abb.IterarRango(&desde, &hasta, func(clave, dato int) bool {
		suma += dato
		return true
	})
	require.Equal(t, 440, suma)

}

func TestIterarCalcularSuma(t *testing.T) {
	abb := crearABBEnteros()
	suma := 0
	abb.Iterar(func(_, dato int) bool {
		suma += dato
		return true
	})
	require.Equal(t, 720, suma)
}

func TestIterarAbbVacio(t *testing.T) {
	abb := diccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	suma := 0
	abb.Iterar(func(_, dato int) bool {
		suma += dato
		return true
	})
	require.Equal(t, 0, suma)

	abb.IterarRango(nil, nil, func(_, dato int) bool {
		suma += dato
		return true
	})
	require.Equal(t, 0, suma)
}

func TestIterarAbbUnElemento(t *testing.T) {
	abb := diccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	abb.Guardar(42, 99)

	suma := 0
	abb.Iterar(func(_, dato int) bool {
		suma += dato
		return true
	})
	require.Equal(t, 99, suma)

	desde := 40
	hasta := 45
	suma = 0
	abb.IterarRango(&desde, &hasta, func(_, dato int) bool {
		suma += dato
		return true
	})
	require.Equal(t, 99, suma)

	desde = 1
	hasta = 10
	suma = 0
	abb.IterarRango(&desde, &hasta, func(_, dato int) bool {
		suma += dato
		return true
	})
	require.Equal(t, 0, suma)
}

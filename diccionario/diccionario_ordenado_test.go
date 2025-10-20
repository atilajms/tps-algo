package diccionario_test

import (
	"fmt"
	"math/rand"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN_ABB = []int{12500, 25000, 50000, 100000}

func compararEnteros(a, b int) int { return a - b }

func crearABBAlfabeto() TDADiccionario.DiccionarioOrdenado[string, int] {
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	letras := []string{
		"m", "g", "t", "d", "j", "p", "w", "b", "e", "h", "k",
		"o", "r", "u", "z", "a", "c", "f", "i", "l", "n", "q",
		"s", "v", "x", "y", "ñ",
	}

	for i, letra := range letras {
		abb.Guardar(letra, i+1)
	}
	return abb
}

func generarClavesAleatorias(n int) []string {
	claves := make([]string, n)
	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
	}
	rand.Shuffle(len(claves), func(i, j int) {
		claves[i], claves[j] = claves[j], claves[i]
	})
	return claves
}

func crearABBEnteros() TDADiccionario.DiccionarioOrdenado[int, int] {
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	abb.Guardar(10, 100)
	abb.Guardar(5, 50)
	abb.Guardar(15, 150)
	abb.Guardar(3, 30)
	abb.Guardar(7, 70)
	abb.Guardar(12, 120)
	abb.Guardar(20, 200)
	return abb
}

func TestABBDiccionarioVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestABBDiccionarioClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un Hash vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](compararEnteros)
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestABBUnElement(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestABBDiccionarioGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestABBReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestABBReemplazoDatoHopscotch(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	dic := TDADiccionario.CrearABB[int, int](compararEnteros)
	for i := 0; i < 500; i++ {
		dic.Guardar(i, i)
	}
	for i := 0; i < 500; i++ {
		dic.Guardar(i, 2*i)
	}
	ok := true
	for i := 0; i < 500 && ok; i++ {
		ok = dic.Obtener(i) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestABBDiccionarioBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	dic := TDADiccionario.CrearABB[string, string](strings.Compare)

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestABBConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := TDADiccionario.CrearABB[int, string](compararEnteros)
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestABBIterarInternoOrdenCorrecto(t *testing.T) {
	t.Log("Iterar internamente debe devolver las claves en orden in-order ascendente.")
	abb := crearABBEnteros()
	var claves []int
	abb.Iterar(func(clave, _ int) bool {
		claves = append(claves, clave)
		return true
	})
	require.Equal(t, []int{3, 5, 7, 10, 12, 15, 20}, claves)
}

func TestABBIterarCorte(t *testing.T) {
	t.Log("La función de iteración interna debe poder cortar la iteración correctamente.")

	abb := crearABBAlfabeto()

	var claves []string
	abb.Iterar(func(clave string, _ int) bool {
		claves = append(claves, clave)
		return clave != "m"
	})

	esperado := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"}
	require.Equal(t, esperado, claves)
}

func TestABBIterarCorteYRango(t *testing.T) {
	t.Log("La función de iteración interna debe poder cortar la iteración y respetar un rango.")

	abb := crearABBAlfabeto()

	var claves []string
	desde := "f"
	hasta := "s"
	abb.IterarRango(&desde, &hasta, func(clave string, _ int) bool {
		claves = append(claves, clave)
		return clave != "m"
	})
	esperado := []string{"f", "g", "h", "i", "j", "k", "l", "m"}
	require.Equal(t, esperado, claves)
}

func TestABBIterarInternoVacio(t *testing.T) {
	t.Log("Iterar un ABB vacío no ejecuta la función visitar.")
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	called := false
	abb.Iterar(func(_, _ int) bool {
		called = true
		return true
	})
	require.False(t, called)
}

func TestABBIteradorExternoCompleto(t *testing.T) {
	t.Log("El iterador externo sin rango debe recorrer todas las claves en orden.")
	abb := crearABBEnteros()
	iter := abb.IteradorRango(nil, nil)

	var claves []int
	var datos []int
	for iter.HaySiguiente() {
		c, v := iter.VerActual()
		claves = append(claves, c)
		datos = append(datos, v)
		iter.Siguiente()
	}
	require.Equal(t, []int{3, 5, 7, 10, 12, 15, 20}, claves)
	require.Equal(t, []int{30, 50, 70, 100, 120, 150, 200}, datos)
	require.False(t, iter.HaySiguiente())

	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestABBIteradorRangoNilNil(t *testing.T) {
	t.Log("Rango (nil, nil) debe incluir todos los elementos en orden ascendente.")
	abb := crearABBEnteros()
	iter := abb.IteradorRango(nil, nil)

	var claves []int
	for iter.HaySiguiente() {
		c, _ := iter.VerActual()
		claves = append(claves, c)
		iter.Siguiente()
	}
	require.Equal(t, []int{3, 5, 7, 10, 12, 15, 20}, claves)
}

func TestABBIteradorRangoDesdeNil(t *testing.T) {
	t.Log("Rango (desde, nil) debe incluir todos los elementos >= desde (inclusivo).")
	abb := crearABBEnteros()
	desde := 10
	iter := abb.IteradorRango(&desde, nil)

	var claves []int
	for iter.HaySiguiente() {
		c, _ := iter.VerActual()
		claves = append(claves, c)
		iter.Siguiente()
	}
	require.Equal(t, []int{10, 12, 15, 20}, claves)
}

func TestABBIteradorRangoNilHasta(t *testing.T) {
	t.Log("Rango (nil, hasta) debe incluir todos los elementos <= hasta (inclusivo).")
	abb := crearABBEnteros()
	hasta := 10
	iter := abb.IteradorRango(nil, &hasta)

	var claves []int
	for iter.HaySiguiente() {
		c, _ := iter.VerActual()
		claves = append(claves, c)
		iter.Siguiente()
	}
	require.Equal(t, []int{3, 5, 7, 10}, claves)
}

func TestABBIteradorRangoDesdeHasta(t *testing.T) {
	t.Log("Rango (desde, hasta) inclusivo debe respetar ambos límites.")
	abb := crearABBEnteros()
	desde := 5
	hasta := 15
	iter := abb.IteradorRango(&desde, &hasta)

	var claves []int
	for iter.HaySiguiente() {
		c, _ := iter.VerActual()
		claves = append(claves, c)
		iter.Siguiente()
	}
	require.Equal(t, []int{5, 7, 10, 12, 15}, claves)
}

func TestABBIteradorRangoSinResultados(t *testing.T) {
	t.Log("Si el rango no contiene elementos, el iterador debe empezar vacío.")
	abb := crearABBEnteros()
	desde := 0
	hasta := 2
	iter := abb.IteradorRango(&desde, &hasta)
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestABBIteradorUnElemento(t *testing.T) {
	t.Log("Iterar un ABB con un solo elemento debe devolver exactamente ese elemento.")
	abb := TDADiccionario.CrearABB[int, int](compararEnteros)
	abb.Guardar(42, 99)

	iter := abb.IteradorRango(nil, nil)
	require.True(t, iter.HaySiguiente())
	clave, dato := iter.VerActual()
	require.Equal(t, 42, clave)
	require.Equal(t, 99, dato)

	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
}

func TestABBIteradorStringsOrdenLexicografico(t *testing.T) {
	t.Log("Con claves string, el ABB debe respetar el orden lexicográfico según strings.Compare.")
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	claves := []string{"D", "B", "A", "C", "E"}
	for i, c := range claves {
		abb.Guardar(c, i)
	}

	iter := abb.IteradorRango(nil, nil)
	var orden []string
	for iter.HaySiguiente() {
		c, _ := iter.VerActual()
		orden = append(orden, c)
		iter.Siguiente()
	}
	require.Equal(t, []string{"A", "B", "C", "D", "E"}, orden)
}

func BenchmarkABB(b *testing.B) {
	b.Log("Benchmark de ABB: guarda, obtiene y borra muchos elementos en orden aleatorio")
	for _, n := range TAMS_VOLUMEN_ABB {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenABB(b, n)
			}
		})
	}
}

func BenchmarkIteradorABB(b *testing.B) {
	b.Log("Benchmark del iterador del ABB con grandes volúmenes de datos")
	for _, n := range TAMS_VOLUMEN_ABB {
		b.Run(fmt.Sprintf("Iterador %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIteradorABB(b, n)
			}
		})
	}
}

func ejecutarPruebaVolumenABB(b *testing.B, n int) {
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	claves := generarClavesAleatorias(n)
	valores := make([]int, n)

	for i := 0; i < n; i++ {
		valores[i] = i
		abb.Guardar(claves[i], valores[i])
	}
	require.EqualValues(b, n, abb.Cantidad())

	ok := true
	for i := 0; i < n; i++ {
		if !abb.Pertenece(claves[i]) || abb.Obtener(claves[i]) != valores[i] {
			ok = false
			break
		}
	}
	require.True(b, ok)
	require.EqualValues(b, n, abb.Cantidad())

	for i := 0; i < n; i++ {
		if abb.Borrar(claves[i]) != valores[i] || abb.Pertenece(claves[i]) {
			ok = false
			break
		}
	}
	require.True(b, ok)
	require.EqualValues(b, 0, abb.Cantidad())
}

func ejecutarPruebasVolumenIteradorABB(b *testing.B, n int) {
	abb := TDADiccionario.CrearABB[string, *int](strings.Compare)
	claves := generarClavesAleatorias(n)
	valores := make([]int, n)

	for i := 0; i < n; i++ {
		valores[i] = i
		abb.Guardar(claves[i], &valores[i])
	}

	iter := abb.Iterador()
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	for i = 0; i < n; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		_, v := iter.VerActual()
		if v == nil {
			ok = false
			break
		}
		*v = n
		iter.Siguiente()
	}
	require.True(b, ok)
	require.EqualValues(b, n, i)
	require.False(b, iter.HaySiguiente())

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok)
}

package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.Equal(t, 0, lista.Largo())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	})
}
func TestInsertarPrimeroUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	dato := 2
	lista.InsertarPrimero(dato)
	require.False(t, lista.EstaVacia())
	require.Equal(t, dato, lista.VerPrimero())
	require.Equal(t, dato, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())

}

func TestInsertarULtimoUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	dato := 123
	lista.InsertarUltimo(dato)
	require.False(t, lista.EstaVacia())
	require.Equal(t, dato, lista.VerPrimero())
	require.Equal(t, dato, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
}

func TestInsertarPrimeroBorrarUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	txt := "hola!"
	lista.InsertarPrimero(txt)
	require.Equal(t, txt, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
}

func TestInsertarUltimoBorrarUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[rune]()
	txt := 'a'
	lista.InsertarUltimo(txt)
	require.False(t, lista.EstaVacia())
	require.Equal(t, txt, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
}

func TestVolumenInsertarUltimoBorrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := range 10000 {
		require.Equal(t, i, lista.Largo())
		lista.InsertarUltimo(i)
		require.Equal(t, i, lista.VerUltimo())
		require.Equal(t, i+1, lista.Largo())
		require.Equal(t, 0, lista.VerPrimero())
	}

	for i := range 10000 {
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.VerPrimero())
		require.Equal(t, 9999, lista.VerUltimo())
		require.Equal(t, i, lista.BorrarPrimero())
		require.Equal(t, 9999-i, lista.Largo())
	}

	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	})
}

func TestVolumenInsertarPrimeroBorrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := range 10000 {
		require.Equal(t, i, lista.Largo())
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.VerPrimero())
		require.Equal(t, i+1, lista.Largo())
		require.Equal(t, 0, lista.VerUltimo())
	}

	for i := 9999; i >= 0; i-- {
		require.False(t, lista.EstaVacia())
		require.Equal(t, i, lista.VerPrimero())
		require.Equal(t, 0, lista.VerUltimo())
		require.Equal(t, i, lista.BorrarPrimero())
		require.Equal(t, i, lista.Largo())
	}

	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerPrimero()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.VerUltimo()
	})
	require.PanicsWithValue(t, "La lista esta vacia", func() {
		lista.BorrarPrimero()
	})
}

func TestIterarMultiplosDe3(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := range 10 {
		lista.InsertarUltimo(i)
	}

	contador := 0
	lista.Iterar(func(n int) bool {
		if n%3 == 0 {
			contador++
		}
		return true
	})

	require.Equal(t, contador, 4)

}

func TestIterarHasta20Impares(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := range 1000 {
		lista.InsertarUltimo(i)
	}
	contador := 0
	impares := 0

	lista.Iterar(func(n int) bool {
		contador++
		if n%2 == 1 {
			impares++
		}

		if impares == 20 {
			return false
		}
		return true
	})

	require.Equal(t, impares, 20)
	require.Equal(t, contador, 40)
}

func TestIteradorConListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()

	require.False(t, iterador.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iterador.Borrar()
	})
	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iterador.VerActual()
	})
	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iterador.Siguiente()
	})
}

func TestIteradorInsertarUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	dato := 10

	iterador.Insertar(dato)

	require.False(t, lista.EstaVacia())
	require.Equal(t, dato, lista.VerUltimo())
	require.Equal(t, dato, lista.VerPrimero())
	require.Equal(t, 1, lista.Largo())

	require.Equal(t, dato, iterador.VerActual())
	require.True(t, iterador.HaySiguiente())

	iterador.Siguiente()
	require.False(t, lista.EstaVacia())
	require.Equal(t, dato, lista.VerUltimo())
	require.Equal(t, dato, lista.VerPrimero())
	require.Equal(t, 1, lista.Largo())

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iterador.VerActual()
	})
	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iterador.Siguiente()
	})
	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iterador.Borrar()
	})
	require.False(t, iterador.HaySiguiente())
}

func TestIteradorInsertarBorrarUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	iterador := lista.Iterador()
	dato := "a"

	iterador.Insertar(dato)
	require.Equal(t, dato, iterador.Borrar())

	require.False(t, iterador.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iterador.Borrar()
	})
	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iterador.VerActual()
	})
	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iterador.Siguiente()
	})
	require.True(t, lista.EstaVacia())
}

func TestIteradorVerActual(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := range 10 {
		lista.InsertarUltimo(i)
	}

	iterador := lista.Iterador()
	for i := range 10 {
		if iterador.HaySiguiente() {
			require.Equal(t, i, iterador.VerActual())
			iterador.Siguiente()
		}
	}
}

func TestIteradorInsertarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := range 10 {
		lista.InsertarUltimo(i)
	}

	iterador := lista.Iterador()
	for iterador.HaySiguiente() {
		iterador.Siguiente()
	}

	dato := 200
	iterador.Insertar(dato)

	require.Equal(t, dato, lista.VerUltimo())
	require.Equal(t, dato, iterador.VerActual())
	require.True(t, iterador.HaySiguiente())
}

func TestIteradorInsertarMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(3)

	iterador := lista.Iterador()
	iterador.Siguiente()

	iterador.Insertar(2)

	i := 1
	for iterador2 := lista.Iterador(); iterador2.HaySiguiente(); iterador2.Siguiente() {
		require.Equal(t, i, iterador2.VerActual())
		i++
	}

}

func TestIteradorBorrarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 5; i++ {
		lista.InsertarUltimo(i)
	}

	iterador := lista.Iterador()

	for iterador.HaySiguiente() {
		if iterador.VerActual() == 5 {
			break
		}
		iterador.Siguiente()
	}

	ultimo := iterador.Borrar()
	require.Equal(t, 5, ultimo)
	require.Equal(t, 4, lista.VerUltimo())
	require.Equal(t, 4, lista.Largo())
}

func TestIteradorBorrarElementoMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 5; i++ {
		lista.InsertarUltimo(i)
	}

	iterador := lista.Iterador()

	for iterador.HaySiguiente() {
		if iterador.VerActual() == 3 {
			break
		}
		iterador.Siguiente()
	}

	borrado := iterador.Borrar()
	require.Equal(t, 3, borrado)

	expected := []int{1, 2, 4, 5}
	i := 0
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		require.Equal(t, expected[i], iter.VerActual())
		i++
	}

	require.Equal(t, len(expected), lista.Largo())
}

func TestIteradorInsertarYBorrarCombinado(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(3)

	iterador := lista.Iterador()
	iterador.Siguiente()
	iterador.Insertar(2)
	require.Equal(t, 2, iterador.VerActual())

	borrado := iterador.Borrar()
	require.Equal(t, 2, borrado)

	expected := []int{1, 3}
	i := 0
	for iter := lista.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		require.Equal(t, expected[i], iter.VerActual())
		i++
	}
	require.Equal(t, len(expected), i)
}

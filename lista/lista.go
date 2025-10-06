package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento a la lista, al inicio de la misma.
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento a la lista, al final de la misma.
	InsertarUltimo(T)

	// BorrarPrimero saca el primer elemento de la lista. Si tiene elementos, se quita el primero de la misma,
	// y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primero de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor del ultimo de la lista. Si está vacía, entra en pánico con un mensaje
	// "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos en la lista.
	Largo() int

	// Iterar recibe una funcion por parametro y la aplica a cada elemento de la lista hasta que la función
	// recibida devuelva false o que haya terminado de iterar todos los elementos.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un iterador externo para la lista a que se refiere.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual obtiene el elemento en el que el iterador está posicionado, y devuelve dicho elemento. Si ya
	// se iteraron todos los elementos, entra en pánico con el mensaje "El iterador termino de iterar".
	VerActual() T

	// Hay siguiente devuleve true si el iterador está posicionado a un elemento existente, false si no hay mas
	// elementos para iterar.
	HaySiguiente() bool

	// Siguiente actualiza la posicion del iterador para ubicarlo en el siguiente elemento. Si no hay
	// más elementos para ver, entra en pánico con el mensaje "El iterador termino de iterar".
	Siguiente()

	// Insertar adiciona el elemento recibido entre el elemento actual y el anterior, y ubica el iterador en el
	// nuevo elemento. Si el iterador apunta al primer elemento, inserta en esta posición; si ya no hay siguiente
	// (es decir, el iterador termino de iterar), inserta en la ultima posicion.
	Insertar(T)

	// BorrarPrimero saca de la lista el elemento a que apunta el iterador, y lo devuelve. El iterador pasa a
	// posicionarse sobre el elemento siguiente al que fue borrado. Si borra el ultimo elemento, significa
	// que el iterador termino de iterar. Si no hay más elementos para ver, entra en pánico con el mensaje
	// "El iterador termino de iterar".
	Borrar() T
}

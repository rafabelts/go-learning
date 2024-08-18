# Pointers
- Go copia los valores cuando se pasa una funcion / metodo, asi que si se esta escribiendo una
  funcion que necesita mutar el estado, necesitaras tomar un puntero hacia la cosa que se quiere
  cambiar

- El hecho de que Go tome una copia de los valores es util la mayoria de veces, pero en ocasiones
  no queras que el sistema haga una copia de algo, en ese caso necesitaras pasar una referencia. Ejemplos
  incluyen referenciar estructuras de datos largas o cosas donde solo una instancia es necesaria (como pools de 
  conexiones a bd)

# nil
- Los punteros pueden ser nil
- Cuando una funcion retorna un puntero hacia algo, necesitaras checar si es nil o si se lanzo una excepcion
- Util cuando quieres describir un valor que puede estar faltando

# Errores
- Son la manera de indicar un fallo al llamar a una funcion/metodo
- Al escuchar nuestros test, concluimos que checar por un string en un error resulta a una prueba defectuosa.
  Asi que se refactoriza para usar un valor significativo, lo que seria mas facil de probar y de entender


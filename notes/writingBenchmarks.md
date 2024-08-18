# Benchmarking
Es similar a escribir tests
```go
func BenchMarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}
```

El testing.B da acceso al crípticamente llamado b.N

Cuando el benchmark es ejecutado, ejecuta b.N veces y mido cuanto tiempo toma

La cantidad de veces que el codigo corre no debe importar, el framework determina cual es un buen valor para permitir obtener resultados decentes

Para correr benchmarks se debe ejecutar ```go test -bench=.```

Estos se corren secuencialmente

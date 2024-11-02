using CSV
using Distributed
using Base.Threads: @threads

# Configuraciones
archivo = "ratings.csv"
num_procesos = 10

# Función para contar las líneas del archivo
function contar_lineas(archivo)
    count = 0
    open(archivo, "r") do f
        for _ in eachline(f)
            count += 1
        end
    end
    return count
end

# Leer el encabezado del archivo
function leer_header(archivo)
    open(archivo, "r") do f
        return readline(f)  # Leer solo la primera línea como encabezado
    end
end

# Parámetros de partición
size_of_file = contar_lineas(archivo) - 1  # Excluyendo el encabezado
number_of_chunks = ceil(Int, size_of_file / num_procesos)
header = leer_header(archivo)

# Función para generar archivos pequeños con encabezado
function generate_small_file(archivo, i, header, num_procesos, number_of_chunks)
    start_line = (i - 1) * number_of_chunks + 2  # Saltamos el encabezado
    end_line = min(i * number_of_chunks + 1, size_of_file + 1)  # Incluir límites
    output_file = "ratings_$i.csv"
    
    open(archivo, "r") do f
        open(output_file, "w") do out
            println(out, header)  # Escribir el encabezado
            for (line_num, line) in enumerate(eachline(f))
                if line_num >= start_line && line_num <= end_line
                    println(out, line)
                end
            end
        end
    end
end

# Medir el tiempo de ejecución
start_time = time()

@threads for i in 1:num_procesos
    generate_small_file(archivo, i, header, num_procesos, number_of_chunks)
end

end_time = time()

println("Tiempo transcurrido: $(end_time - start_time) segundos")

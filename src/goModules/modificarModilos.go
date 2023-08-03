/*// Inicializar un proyecto
go mod init path_del_proyecto

// Verificar que el código externo no esté corrupto
go mod verify

// Reemplazar fuente del código
go mod edit -replace path_del_repo_online=path_del_repo_en_local

// Quitar el replace
go mod edit -dropreplace path_del_repo_online

// Empaquetar todo el código de terceros que usa nuestro código
go mod vendor

// Eliminar todos los paquetes externos que no estemos usando
go mod tidy

// Aprender más de go modules
go help mod

// replace en go modules para redireccionar una libreria hacia un repo especifico

//go mod edit -replace {{modulo origen}}={{ruta modulo local a redireccionar - destino}}

// reversar el replace
//go mod edit -dropreplace {{modulo a devolver}}

// conjunto de codigo adicional que va a agregar para que el codigo que se tenga en el main funciones de forma normal
//go mod vendor

// limpir librerias
// go mod tidy

// verificar el correcto funcionamiento de las librerias
// go mod verify
*/

/*
-- tour.golang.org
-- golangweekly.com
-- play-with-go.dev
-- gobyexample.com

*/
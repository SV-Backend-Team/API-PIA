# API-PIA
## Equipo
 - Sebastian Terrazas Santillana 1847317
 - Vicente Garza Reyna           1847176
## Requisitos 
### Backend
* GO 
* Microsoft SQL Server Management Studio 18 
*	SQL Server 2019 

### Frontend
*	Angular

## Frontend
Para ver como instalar el Frontend y sus requisitos ir al siguiente [enlace](https://github.com/SV-Backend-Team/Frontend-PIA).

## Instalaciones de Requisitos de Backend

### Go
Para instalar el GO en su machina requiere ir y instalar el GO también se recomienda instalar el Visual Studio Code. 
*	Puede instalar GO al descargar el [MS](https://golang.org/doc/install) de su página oficial. 
*	Puede también puede instalar el Visual Studio Code del siguiente [enlace](https://code.visualstudio.com/) si se instala tendrá que ir a la parte de Extensiones e instalar el GO extensión de la VS Code Market (para ver documentación de la extensión ir al siguiente [enlace](https://code.visualstudio.com/docs/languages/go)) para descargar el VSC ir a el siguiente [enlace](https://marketplace.visualstudio.com/items?itemName=golang.go) y para ver como instalarlo ir al siguiente [enlace](https://www.youtube.com/watch?v=MlIzFUI1QGA). 

### Base de Datos 
1.	Se uso Microsoft SQL Server Management Studio 18 para esto se requiere descargar de [Download SQL Server Management Studio (SSMS)](https://docs.microsoft.com/en-us/sql/ssms/download-sql-server-management-studio-ssms?view=sql-server-ver15). 
2.	Pero en adición de instalar el SQL Server 2019 donde se puede encontrar en [Download SQL Server](https://www.microsoft.com/es-mx/sql-server/sql-server-downloads). 
Nota: Para ver como instalar el SSMS y el Server se pueden ver los siguientes videos: 
*	Microsoft SQL Server Management Studio 18 
    *	https://www.youtube.com/watch?v=bhC4cORjxrA  
*	SQL Server 2019 
    *	https://www.youtube.com/watch?v=YOaC_TyOrdk 
3.	Una vez instalado estos dos recursos hay que abrir el Microsoft SQL Server Management Studio 18 y correr una vez el archivo sql que se encuentra en los archivos en especifico la carpeta Base de Datos. \
 [![image](https://user-images.githubusercontent.com/54513488/119368259-6914cb80-bc78-11eb-8466-9e0331a3ef34.png)](https://github.com/SV-Backend-Team/API-PIA/tree/master/Base%20de%20Datos)

## Instalación y Run del Programa
### Backend
1.	Para descargar el Backend ir a el siguiente [enlace](https://github.com/SV-Backend-Team/API-PIA/).
   -	Si se tiene la Consola Git hacer Git Bash en un lugar donde se desea descargar y escribir el siguiente comando: 
```
git clone https://github.com/SV-Backend-Team/API-PIA.git
```
   -	Si no se tiene ambos Visual Studio y/o la Consola Git se puede descargar el repositorio como un Zip y instalar manualmente
 
2.	Una vez descargado ejecutar el siguiente comando parar inicializar el programa: 
```
go run .
```
<p align="center">
  <img src="https://user-images.githubusercontent.com/54513488/119378761-3d97de00-bc84-11eb-905b-07c4c993ed6c.png">
</p>

## Endpoints
Todo endpoint tiene como URL inicial `http://localhost:5000` seguido de esto va dependiendo de a que registro valla dirigido o si es del JWT y dependen del método HTTP de Request que tenga. A continuación, se presenta todos los endpoints posibles en API.

### Employees
#### GetEmployees
* **Descripcion:** Devuelve en formato JSON todos los empleados registrados en la tabla
* **Método HTTP del Request:** `GET`
* **URL:** `http://localhost:5000/api/employee/getemployees`
* **Parámetros del URL:** Ninguno
* **Parámetros del Body:** Ninguno
* **Pruebas en Postman:** 
![image](https://user-images.githubusercontent.com/54513488/119393483-9f614380-bc96-11eb-8380-a99d6d8483ac.png)

#### GetEmployeeByID
* **Descripcion:** Devuelve en formato JSON un empleado especifico solicitado de los registros
* **Método HTTP del Request:** `GET`
* **URL:** `http://localhost:5000/api/employee/getemployee/:id`
* **Parámetros del URL:** `id` el cual representa siempre un valor numérico positivo
* **Parámetros del Body:** Ninguno
* **Pruebas en Postman:** 
![image](https://user-images.githubusercontent.com/54513488/119393669-d9cae080-bc96-11eb-88d0-3565ba868616.png)

#### CreateEmployee
* **Descripcion:** Crea un nuevo registro en la tabla Employees con información proveída en el Body. Al finalizar la creación devuelve un JSON CON el Empleado con su ID único.
* **Método HTTP del Request:** `POST`
* **URL:** `http://localhost:5000/api/employee/createemployee`
* **Parámetros del URL:** Ninguno
* **Parámetros del Body:
```
{
    "EmployeeID": integer,
    "LastName": string,
    "FirstName": string,
    "Title": string,
    "TitleOfCourtesy": string,
    "BirthDate": string,
    "HireDate": string,
    "Address": string,
    "City": string,
    "Region": string,
    "PostalCode": string,
    "Country": string,
    "HomePhone": string,
    "Extension": string,
    "Notes": string,
    "ReportsTo": string,
    "PhotoPath": string
}
```
**Nota 1:** Esto al ejecutar lo regresara imprimido con el Status 200 para ejemplificar que se creó con éxito el Employee.
**Nota 2:** Los únicos datos requeridos a llenar siempre son el LastName y FirstName.

* **Pruebas en Postman:**
![image](https://user-images.githubusercontent.com/54513488/119393786-041c9e00-bc97-11eb-847c-5d513ab22bf6.png)

#### UpdateEmployee
* **Descripcion:** Modifica un registro solicitado en cual se hace en el Body. En esta modificación no se puede cambiar el ID esto es inmutable.
* **Método HTTP del Request:** `PUT`
* **URL:** `http://localhost:5000/api/employee/updateemployee`
* **Parámetros del URL:** Ninguno
* **Parámetros del Body:**
```
{
    "EmployeeID": integer,
    "LastName": string,
    "FirstName": string,
    "Title": string,
    "TitleOfCourtesy": string,
    "BirthDate": string,
    "HireDate": string,
    "Address": string,
    "City": string,
    "Region": string,
    "PostalCode": string,
    "Country": string,
    "HomePhone": string,
    "Extension": string,
    "Notes": string,
    "ReportsTo": string,
    "PhotoPath": string
}
```
**Nota 1:** Este Formato solo imprimirá lo de Parámetros del Body en el Status 200 para ejemplificar que se cambió con éxito.
**Nota 2:** Los únicos datos requeridos a llenar siempre o que no pueden quedar vacíos son el LastName y FirstName.

* **Pruebas en Postman:**
![image](https://user-images.githubusercontent.com/54513488/119393902-2e6e5b80-bc97-11eb-9d04-c5b06510d561.png)

#### DeleteEmployee
* **Descripcion:** Elimina un registro de la Tabla Employees esto mediante una solicitud en el parámetro del URL para saber cuál eliminar.
* **Método HTTP del Request:** `DELETE`
* **URL:** `http://localhost:5000/api/employee/deleteemployee/:id`
* **Parámetros del URL:** `id` el cual representa siempre un valor numérico positivo
* **Parámetros del Body:** Ninguno
* **Pruebas en Postman:**
![image](https://user-images.githubusercontent.com/54513488/119394048-5e1d6380-bc97-11eb-8177-dc45ad86e299.png)
**Nota:** Cuando se pida el DELETE solo se mostrará el EmployeeID, LastName, FirstName cuando se ejecute.

### Customer
#### GetCustomers
* **Descripcion:** Devuelve en formato JSON todos los customers registrados en la tabla.
* **Método HTTP del Request:** `GET`
* **URL:** `http://localhost:5000/api/customer/getcustomers`
* **Parámetros del URL:** Ninguno
* **Parámetros del Body:** Ninguno
* **Pruebas en Postman:**
![image](https://user-images.githubusercontent.com/54513488/119394234-91f88900-bc97-11eb-87af-e2209f20686a.png)

#### GetCustomerByID
* **Descripcion:** Devuelve en formato JSON un customer especifico solicitado de los registros.
* **Método HTTP del Request:** `GET`
* **URL:** `http://localhost:5000/api/customer/getcustomer/:id`
* **Parámetros del URL:** `id` el cual representa siempre un valor Alfabético de 5 letras mayúsculas.
* **Parámetros del Body:** Ninguno
* **Pruebas en Postman:**
![image](https://user-images.githubusercontent.com/54513488/119394341-accafd80-bc97-11eb-83c5-3b0fde1b2924.png)

#### CreateCustomer
* **Descripcion:** Crea un nuevo registro en la tabla Customers con información proveída en el Body. En donde los ID son creados y modificables pero los cuales no pueden ser idénticos a otros ya existentes.
* **Método HTTP del Request:** `POST`
* **URL:** `http://localhost:5000/api/customer/createcustomer`
* **Parámetros del URL:** Ninguno
* **Parámetros del Body:**
```
{
    "CustomerID": string,
    "CompanyName": string,
    "ContactName": string,
    "ContactTitle": string,
    "Address": string,
    "City": string,
    "Region": string,
    "PostalCode": string,
    "Country": string,
    "Phone": string,
    "Fax": string
}
```
**Nota 1:** Esto al ejecutar lo regresara imprimido con el Status 200 para ejemplificar que se creó con éxito el Customer.
**Nota 2:** Los únicos datos requeridos a llenar siempre son el CustomerID y CompanyName.

* **Pruebas en Postman:**
![image](https://user-images.githubusercontent.com/54513488/119394458-ccfabc80-bc97-11eb-8edb-17a0838751b6.png)

#### UpdateCustomer
* **Descripcion:** Modifica un registro solicitado en cual se hace en el Body. 
* **Método HTTP del Request:** `PUT`
* **URL:** `http://localhost:5000/api/customer/updatecustomer`
* **Parámetros del URL:** Ninguno
* **Parámetros del Body:**
```
{
    "CustomerID": string,
    "CompanyName": string,
    "ContactName": string,
    "ContactTitle": string,
    "Address": string,
    "City": string,
    "Region": string,
    "PostalCode": string,
    "Country": string,
    "Phone": string,
    "Fax": string
}
```
**Nota 1:** Esto al ejecutar lo regresara imprimido con el Status 200 para ejemplificar que se creó con éxito el Customer.
**Nota 2:** Los únicos datos requeridos a llenar siempre son el CustomerID y CompanyName.
* **Pruebas en Postman:**
![image](https://user-images.githubusercontent.com/54513488/119394595-f9163d80-bc97-11eb-8aca-6f6fdda3ea16.png)


#### DeleteCustomer
* **Descripcion:** Elimina un registro de la Tabla Customers esto mediante una solicitud en el parámetro del URL para saber cuál eliminar.
* **Método HTTP del Request:** `DELETE`
* **URL:** `http://localhost:5000/api/employee/deleteemployee/:id` 
* **Parámetros del URL:** `id` el cual representa siempre un valor Alfabético de 5 letras mayúsculas.
* **Parámetros del Body:** Ninguno
* **Pruebas en Postman:**
![image](https://user-images.githubusercontent.com/54513488/119394725-2367fb00-bc98-11eb-803e-fce947452481.png)

**Nota:** A comparación con DeleteEmployee este si mostrara todos los valores del registro borrado.

### JWT
#### CreateToken
* **Descripcion:** Crea el token de autorización para poder usar en el resto de los Controladores.
* **Método HTTP del Request:** `POST`
* **URL:** `http://localhost:5000/misc/get-token?app=example`
* **Parámetros del URL:** Se conforma del “?app=” lo que le sigue es el token.
* **Parámetros del Body:** Ninguno
* **Pruebas en Postman:**
![image](https://user-images.githubusercontent.com/54513488/119394833-498d9b00-bc98-11eb-860f-efcd8ba51ed8.png)

**Nota:** Este Token para poder ser usado tiene que ser redimido en Parámetros de Authorization de Tipo Bearer Token.

###Postman
Si se quiere probar los endpoints se puede descargar el Postman Collection encontrado en el siguiente [link](https://github.com/SV-Backend-Team/API-PIA/tree/master/Postman)



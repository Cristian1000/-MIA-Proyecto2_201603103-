# Universidad De San Carlos de Guatemala
# Facultad de Ingeniería

# Escuela de Ciencias y Sistemas

# Manejo e Implemantacion de Archivos

## Manual Tecnico

**Nombre: Cristian Daniel Raguay Vicente**

**Carne: 201603103**

**Sección: “A-”**

[Enunciado del Proyecto](https://1drv.ms/b/s!AkHFM7g7YKzJgb0MlBdBV1FFGCQK3w?e=dtLy5X)

## Información técnica

Editor de Código usado: Visual Studio Code

Lenguajes  utilizados para el Desarrollo del proyecto
Java Scrip 6.14.4
Golan version go1.13.8 linux/amd64
Oracle 18C
Docker para la base de datos
Framework
React Js 17.0.2

bootstrap 4.6.0

![Diagrama ER](/imagen/[MIA]Proyecto_2.png "This is a sample image.")

## Para comenzar los servidores
~~~
Para inicial el servidor de React debe estar ubicado en la carpeta front-end seguido del comando
npm start

Para iniciar el servidor de Go debe estar ubicado en la carpeta back-end seguido del comando
go run main

Para obtener la imagen de oracle se usa el comando
sudo docker pull dockerhelp/docker-oracle-ee-18c
Para iniciar el contenedor en Docker se usa los comandos
sudo docker run -p 1521:1521 -it dockerhelp/docker-oracle-ee-18c bash
Dentro del contenedor se ejecunatan los comandos
sh post_install.sh
sqlplus / as sysdba
alter session set "_ORACLE_SCRIPT"=true;

Desues de esto se uede crear el usuario con el que se va a trabajar
Ejemplo
create user TEST identified by 1234;
GRANT ALL PRIVILEGES TO TEST;
~~~

[Cliente sqlplus](https://www.oracle.com/es/database/technologies/instant-client/linux-x86-64-downloads.html)

Se utilizo Visual Studio Code para conectarse a Oracle
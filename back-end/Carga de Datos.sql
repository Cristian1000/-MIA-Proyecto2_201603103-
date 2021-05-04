CREATE OR REPLACE PROCEDURE Ingresar_Cliente (nombreE in VARCHAR2, apellidoE in VARCHAR2, passE in VARCHAR2, usernameE in VARCHAR2)
AS
BEGIN
            INSERT INTO CLIENTE (NOMBRE,APELLIDO, PASSWORD, USUARIO)
            VALUES (nombreE, apellidoE,passE,usernameE);
END;

CREATE OR REPLACE PROCEDURE Ingresar_Deporte (nombreE in VARCHAR2)
AS
BEGIN
        INSERT INTO DEPORTE (NOMBRE)
        VALUES (nombreE);
END;

CREATE OR REPLACE PROCEDURE Ingresar_temporada (nombre in VARCHAR2)
AS
BEGIN
        INSERT INTO TEMPORADA (NOMBRE)
        VALUES (nombre);
END;

CREATE OR REPLACE PROCEDURE Ingresar_Membresia (nombreE in VARCHAR2)
AS
BEGIN
        INSERT INTO MEMBRESIA (NOMBRE)
        VALUES (nombreE);
END;

CREATE OR REPLACE PROCEDURE Ingresar_Evento (nombreL in VARCHAR2, nombreV in VARCHAR2, resultadoL in INTEGER, resultadoV in INTEGER, fecha in VARCHAR2, deporte in VARCHAR2, jornada in VARCHAR2, temporadaE in INTEGER)
AS
BEGIN
        INSERT INTO EVENTO (NOMBRE_LOCAL, NOMBRE_VISITANTE, R_LOCAL, R_VISITANTE, FECHA, ID_DEPORTE, ID_JORNADA)
        VALUES (nombreL, 
        nombreV,
        resultadoL,
        resultadoV,
        TO_DATE(fecha, 'DD-MM-YYYY HH:MI'),
        (select DEPORTE.ID from DEPORTE where DEPORTE.NOMBRE = deporte),
        (SELECT JORNADA.ID from JORNADA where JORNADA.NOMBRE = jornada AND JORNADA.ID_TEMPORADA = temporadaE)
        );
END;

CREATE OR REPLACE PROCEDURE Ingresar_Jornada (nombreE in VARCHAR2, fecha_i in VARCHAR2, fecha_f in VARCHAR2, temporadaE in VARCHAR2, faseE in VARCHAR2)
AS
BEGIN
        INSERT INTO JORNADA (NOMBRE, FECHA_INICIO, FECHA_FIN, ID_TEMPORADA, ID_FASE)
        VALUES (nombreE, 
        TO_DATE(fecha_i, 'DD-MM-YYYY HH:MI'),
        TO_DATE(fecha_f, 'DD-MM-YYYY HH:MI'),
        (SELECT TEMPORADA.ID from TEMPORADA where TEMPORADA.NOMBRE = temporadaE),
        (select FASE.ID from FASE WHERE FASE.NOMBRE = faseE)
        );
END;

CREATE OR REPLACE PROCEDURE Ingresar_Membresia_temp (cli in VARCHAR2, mem in VARCHAR2, tempo in VARCHAR2)
AS
BEGIN
        INSERT INTO MEMBRESIA_TEMPORADA (ID_CLIENTE, ID_MEMBRESIA, ID_TEMPORADA)
        VALUES ((SELECT cliente.ID from CLIENTE where CLIENTE.USUARIO = cli), 
        (SELECT MEMBRESIA.ID from MEMBRESIA where MEMBRESIA.NOMBRE = mem),
        (SELECT TEMPORADA.ID from TEMPORADA where TEMPORADA.NOMBRE = tempo)
        );
END;

CREATE OR REPLACE PROCEDURE Ingresar_Prediccion (pre_local in INTEGER, pre_visitante in INTEGER, cli in VARCHAR2, even in INTEGER)
AS
BEGIN
        INSERT INTO PREDICCION (PUNTOD_LOCAL, PUNTOS_VISITANTE,ID_CLIENTE, ID_EVENTO)
        VALUES (pre_local, 
        pre_visitante,
        (SELECT cliente.ID from CLIENTE where CLIENTE.NOMBRE = cli),
        even
        );
END;

CREATE OR REPLACE PROCEDURE RetornarEvento (nombreL in VARCHAR2, nombreV in )
AS
BEGIN
        INSERT INTO PREDICCION (PUNTOD_LOCAL, PUNTOS_VISITANTE,ID_CLIENTE, ID_EVENTO)
        VALUES (pre_local, 
        pre_visitante,
        (SELECT cliente.ID from CLIENTE where CLIENTE.NOMBRE = cli),
        even
        );
END;

CALL Ingresar_Cliente('jose', 'ra', '1234', 'car');

select usuario from CLIENTE WHERE USUARIO = ;

SELECT * from TEMPORADA
SELECT * FROM JORNADA 

select object_type,count(*) from user_objects where status = 'INVALID' 
group by object_type;

drop PROCEDURE Ingresar_Cliente;

SELECT * from DEPORTE;
SELECT * FROM EVENTO




DELETE FROM EVENTO;
DELETE FROM JORNADA;
DELETE FROM MEMBRESIA_TEMPORADA;
DELETE FROM PREDICCION;
DELETE FROM TEMPORADA;
DELETE FROM CLIENTE;
DELETE FROM MEMBRESIA;
DELETE FROM DEPORTE;

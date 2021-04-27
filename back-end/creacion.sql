-- Create a new relational table with 3 columns

CREATE TABLE Cliente 
(
  id INTEGER NOT NULL,
  nombre VARCHAR2(25) NOT NULL,
  apellido VARCHAR2(25),
  password VARCHAR2(25),
  USUARIO VARCHAR2(25),
);

ALTER TABLE Cliente ADD CONSTRAINT cliente_pk PRIMARY KEY ( id );

-- Create a new relational table with 3 columns

CREATE TABLE Temporada 
(
  id INTEGER NOT NULL,
  nombre VARCHAR2(65),
  fecha_inicio DATE,
  fecha_fin DATE
);

ALTER TABLE Temporada ADD CONSTRAINT temporada_pk PRIMARY KEY ( id );

-- Create a new relational table with 3 columns

CREATE TABLE Fase 
(
  id INTEGER NOT NULL,
  nombre VARCHAR2(300)
);

ALTER TABLE Fase ADD CONSTRAINT fase_pk PRIMARY KEY ( id );

-- Create a new relational table with 3 columns

CREATE TABLE Jornada 
(
  id INTEGER NOT NULL,
  nombre VARCHAR2(300),
  fecha_inicio DATE,
  fecha_fin DATE,
  id_temporada INTEGER NOT NULL,
  id_fase INTEGER NOT NULL
);

ALTER TABLE Jornada ADD CONSTRAINT jornada_pk PRIMARY KEY ( id );

-- Create a new relational table with 3 columns

CREATE TABLE Deporte 
(
  id VARCHAR2(255) NOT NULL,
  nombre VARCHAR2(1024),
  imagen VARCHAR2(1024),
  color VARCHAR2(1024)
);

ALTER TABLE Deporte ADD CONSTRAINT deporte_pk PRIMARY KEY ( id );

-- Create a new relational table with 3 columns

CREATE TABLE Membresia 
(
  id INTEGER NOT NULL,
  nombre VARCHAR2(1024)
);

ALTER TABLE Membresia ADD CONSTRAINT membresia_pk PRIMARY KEY ( id );

-- Create a new relational table with 3 columns

CREATE TABLE Evento 
(
  id VARCHAR2(255) NOT NULL,
  nombre_local VARCHAR2(1024),
  nnombre_visitante VARCHAR2(1024),
  p_local INTEGER, 
  p_p_visitante INTEGER,
  fecha DATE,
  id_deporte INTEGER NOT NULL
);

ALTER TABLE Evento ADD CONSTRAINT evento_pk PRIMARY KEY ( id );

-- Create a new relational table with 3 columns

CREATE TABLE Prediccion 
(
  id VARCHAR2(255) NOT NULL,
  puntod_local INTEGER,
  puntos_visitante INTEGER,
  puntos_obtenidos INTEGER,
  id_cliente INTEGER NOT NULL ,
  id_evento INTEGER NOT NULL
);

ALTER TABLE Prediccion ADD CONSTRAINT prediccion_pk PRIMARY KEY ( id );

-- Create a new relational table with 3 columns

CREATE TABLE Membresia_Temporada 
(
  id INTEGER NOT NULL,
  id_cliente INTEGER NOT NULL,
  id_membresia INTEGER NOT NULL,
  id_temporada INTEGER NOT NULL
);

ALTER TABLE Membresia_Temporada ADD CONSTRAINT membresia_t_pk PRIMARY KEY ( id );

-- Create a new relational table with 3 columns

CREATE TABLE Bitacora 
(
  id INTEGER NOT NULL,
  usuario VARCHAR2(100),
  accion VARCHAR2(300)
);

alter TABLE Bitacora ADD CONSTRAINT bitacora_pk PRIMARY key (id);


CREATE SEQUENCE Autoincremento
START WITH 1
INCREMENT BY 1;

CREATE TRIGGER TRIG_FAB
BEFORE INSERT ON CLIENTE
FOR EACH ROW
BEGIN
SELECT Autoincremento.NEXTVAL INTO :NEW.id FROM DUAL;
END;



SELECT * FROM cliente

-- Insert rows in a Table

INSERT INTO cliente 
(
  ID,
  NOMBRE
)
VALUES
(
  1,
  'Cristian'
);
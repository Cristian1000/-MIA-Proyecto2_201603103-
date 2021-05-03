-- Create a new relational table with 3 columns

CREATE TABLE Cliente 
(
  id INTEGER NOT NULL,
  nombre VARCHAR2(25) NOT NULL,
  apellido VARCHAR2(25),
  password VARCHAR2(25),
  USUARIO VARCHAR2(25)
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
  id INTEGER NOT NULL,
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
  id INTEGER NOT NULL,
  nombre_local VARCHAR2(1024),
  nombre_visitante VARCHAR2(1024),
  r_local INTEGER, 
  r_visitante INTEGER,
  fecha DATE,
  id_deporte INTEGER NOT NULL,
  id_jornada INTEGER NOT NULL
);

ALTER TABLE Evento ADD CONSTRAINT evento_pk PRIMARY KEY ( id );

-- Create a new relational table with 3 columns

CREATE TABLE Prediccion 
(
  id INTEGER NOT NULL,
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

ALTER TABLE JORNADA
    ADD CONSTRAINT fase_jornada_fk FOREIGN KEY ( ID_FASE )
        REFERENCES FASE ( id );

ALTER TABLE EVENTO
    ADD CONSTRAINT jornada_evento_fk FOREIGN KEY ( id_jornada )
        REFERENCES JORNADA ( ID );

ALTER TABLE EVENTO
    ADD CONSTRAINT deporte_evento_fk FOREIGN KEY ( ID_DEPORTE )
        REFERENCES DEPORTE ( ID );

ALTER TABLE PREDICCION
    ADD CONSTRAINT evento_prediccion_fk FOREIGN KEY ( ID_EVENTO )
        REFERENCES EVENTO ( ID );

ALTER TABLE PREDICCION
    ADD CONSTRAINT cliente_prediccion_fk FOREIGN KEY ( ID_CLIENTE )
        REFERENCES CLIENTE ( id )

ALTER TABLE JORNADA
    ADD CONSTRAINT temporada_jornada_fk FOREIGN KEY ( ID_TEMPORADA )
        REFERENCES TEMPORADA ( ID );

ALTER TABLE MEMBRESIA_TEMPORADA
    ADD CONSTRAINT cliente_temp_fk FOREIGN KEY ( ID_CLIENTE )
        REFERENCES CLIENTE ( id );

ALTER TABLE MEMBRESIA_TEMPORADA
    ADD CONSTRAINT temporada_temp_fk FOREIGN KEY ( ID_TEMPORADA )
        REFERENCES TEMPORADA ( id );

ALTER TABLE MEMBRESIA_TEMPORADA
    ADD CONSTRAINT membresia_temp_fk FOREIGN KEY ( ID_MEMBRESIA )
        REFERENCES MEMBRESIA ( ID );


CREATE SEQUENCE Autoincremento
START WITH 1
INCREMENT BY 1;

CREATE TRIGGER AUTOINCREMENTO_Cliente
BEFORE INSERT ON CLIENTE
FOR EACH ROW
BEGIN
SELECT Autoincremento.NEXTVAL INTO :NEW.id FROM DUAL;
END;

-- Insert rows in a Table

INSERT INTO FASE 
(
  ID,
  NOMBRE
)
VALUES
(
  3,
  'Finalizada'
);

SELECT ID, NOMBRE, CONTRASEÑA FROM cliente

-- Insert rows in a Table

INSERT INTO cliente 
(
  NOMBRE,
  APELLIDO,
  CONTRASEÑA
)
VALUES
(
  'Cris',
  'Vicente',
  '123'
);

-- Delete rows from a Table

DELETE FROM  Cliente
WHERE id = 1;

-- Drop a table

DROP TABLE MEMBRESIA_TEMPORADA;
-- Drop a table

DROP TABLE PREDICCION;
-- Drop a table

DROP TABLE EVENTO;
-- Drop a table

DROP TABLE DEPORTE;    

--Drop a Procedure

DROP TABLE JORNADA;
-- Drop a table

DROP TABLE FASE;
-- Drop a table

DROP TABLE CLIENTE;
-- Drop a table

DROP TABLE MEMBRESIA;
-- Drop a table

DROP TABLE TEMPORADA;    

-- Insert rows in a Table

INSERT INTO Deporte 
(
  NOMBRE,
  IMAGEN,
  COLOR
)
VALUES
(
  'Golf',
  'FirstName.LastName',
  'black'
);
CREATE  SEQUENCE Autoincremento
START WITH 1
INCREMENT BY 1;

CREATE  TRIGGER AUTOINCREMENTO_Cliente
BEFORE INSERT ON CLIENTE
FOR EACH ROW
BEGIN
SELECT Autoincremento.NEXTVAL INTO :NEW.id FROM DUAL;
END;

CREATE  SEQUENCE Autoincremento_Deporte
START WITH 1
INCREMENT BY 1;

CREATE  TRIGGER AUTOINCREMENTO_Deporte
BEFORE INSERT ON DEPORTE
FOR EACH ROW
BEGIN
SELECT Autoincremento_Deporte.NEXTVAL INTO :NEW.id FROM DUAL;
END;

CREATE  SEQUENCE Autoincremento_Evento
START WITH 1
INCREMENT BY 1;

CREATE  TRIGGER AUTOINCREMENTO_Evento
BEFORE INSERT ON EVENTO
FOR EACH ROW
BEGIN
SELECT Autoincremento_Evento.NEXTVAL INTO :NEW.id FROM DUAL;
END;

CREATE  SEQUENCE Autoincremento_fase
START WITH 1
INCREMENT BY 1;

CREATE  TRIGGER AUTOINCREMENTO_Fase
BEFORE INSERT ON FASE
FOR EACH ROW
BEGIN
SELECT Autoincremento_fase.NEXTVAL INTO :NEW.id FROM DUAL;
END;

CREATE  SEQUENCE Autoincremento_Jornada
START WITH 1
INCREMENT BY 1;

CREATE  TRIGGER AUTOINCREMENTO_Jornada
BEFORE INSERT ON JORNADA
FOR EACH ROW
BEGIN
SELECT Autoincremento_Jornada.NEXTVAL INTO :NEW.id FROM DUAL;
END;

CREATE  SEQUENCE Autoincremento_Membresia
START WITH 1
INCREMENT BY 1;

CREATE  TRIGGER AUTOINCREMENTO_Membresia
BEFORE INSERT ON MEMBRESIA
FOR EACH ROW
BEGIN
SELECT Autoincremento_Membresia.NEXTVAL INTO :NEW.id FROM DUAL;
END;

CREATE  SEQUENCE Autoincremento_Membresia_t
START WITH 1
INCREMENT BY 1;

CREATE  TRIGGER AUTOINCREMENTO_Membresia_t
BEFORE INSERT ON MEMBRESIA_TEMPORADA
FOR EACH ROW
BEGIN
SELECT Autoincremento_Membresia_t.NEXTVAL INTO :NEW.id FROM DUAL;
END;

CREATE  SEQUENCE Autoincremento_Prediccion
START WITH 1
INCREMENT BY 1;

CREATE  TRIGGER AUTOINCREMENTO_Prediccion
BEFORE INSERT ON PREDICCION
FOR EACH ROW
BEGIN
SELECT Autoincremento_Prediccion.NEXTVAL INTO :NEW.id FROM DUAL;
END;

CREATE  SEQUENCE Autoincremento_Temporada
START WITH 1
INCREMENT BY 1;

CREATE  TRIGGER AUTOINCREMENTO__temporada
BEFORE INSERT ON TEMPORADA
FOR EACH ROW
BEGIN
SELECT Autoincremento_Temporada.NEXTVAL INTO :NEW.id FROM DUAL;
END;

CREATE  SEQUENCE Autoincremento_Bitacora
START WITH 1
INCREMENT BY 1;

CREATE  TRIGGER AUTOINCREMENTO_Bitacora
BEFORE INSERT ON BITACORA
FOR EACH ROW
BEGIN
SELECT Autoincremento_Bitacora.NEXTVAL INTO :NEW.id FROM DUAL;
END;


drop TRIGGER Validar_Usuario;

drop SEQUENCE Autoincremento;
drop SEQUENCE AUTOINCREMENTO_DEPORTE;
drop SEQUENCE AUTOINCREMENTO_EVENTO;
drop SEQUENCE AUTOINCREMENTO_FASE;
drop SEQUENCE AUTOINCREMENTO_JORNADA;
drop SEQUENCE AUTOINCREMENTO_MEMBRESIA;
drop SEQUENCE AUTOINCREMENTO_MEMBRESIA_T;
drop SEQUENCE AUTOINCREMENTO_PREDICCION;
drop SEQUENCE AUTOINCREMENTO_TEMPORADA;

SELECT nombre FROM cliente


------trigger----
--Create a new Table Trigger

COMMIT

create or replace  TRIGGER Validar_Usuario
AFTER INSERT ON CLIENTE
FOR EACH ROW
    DECLARE
    correoI VARCHAR2;
    nombreI VARCHAR2;
    edad INTEGER;
    actual DATE;
BEGIN
        SELECT (CLIENTE.USUARIO) into nombreI FROM CLIENTE WHERE CLIENTE.USUARIO = :new.USUARIO;
        edad := TO_NUMBER(to_char(:new.fecha_nacimiento, 'YYYY'));
        SELECT regexp_substr(:new.correo,'[a-zA-Z0-9._%-]+@[a-zA-Z0-9._%-]+\.[a-zA-Z]{2,4}') into correoI FROM DUAL;
        SELECT SYSDATE INTO actual from DUAL;
        edad := TO_NUMBER(to_char(actual, 'YYYY')) - edad;
        
        if (correoI = '' or edad < 18 or nombreI <> '') THEN
            DELETE FROM CLIENTE WHERE CLIENTE.ID = :new.id ;
            DBMS_OUTPUT.PUT_LINE('Cliente no valido');
        end if;
END;

SELECT * FROM CLIENTE

create or replace trigger Cargar_Puntos
after update on Evento
for each row
DECLARE
BEGIN

END;


-- Insert rows in a Table

INSERT INTO CLIENTE 
(
  NOMBRE,
  APELLIDO,
  PASSWORD,
  USUARIO,
  FECHA_NACIMIENTO,
  CORREO
)
VALUES
(
  'Cris',
  'Raguay',
  '1235',
  'pablo',
  '25/may/2010',
  'dfghjkl'
);



SELECT     regexp_substr('cristian@gmail.com','[a-zA-Z0-9._%-]+@[a-zA-Z0-9._%-]+\.[a-zA-Z]{2,4}') FROM  DUAL
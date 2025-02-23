CREATE TABLE participantes (
    id int  PRIMARY KEY auto_increment NOT NULL,
    nome VARCHAR(30) NOT NULL,
    residencia VARCHAR(50),
    ocupacao VARCHAR(30),
    status BOOLEAN
);

CREATE TABLE historico_votos (
id int  PRIMARY KEY auto_increment NOT NULL,
id_participante int NOT NULL,
ip VARCHAR(39),
created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
FOREIGN KEY (id_participante) REFERENCES participantes(id) ON DELETE CASCADE
);
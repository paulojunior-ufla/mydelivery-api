create table cliente (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT CHECK( nome != '' ),
    email TEXT CHECK( email != '' ),
    telefone TEXT CHECK( telefone != '' )
);

create table entrega (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    cliente_id INTEGER NOT NULL,
    taxa REAL NOT NULL,
    status TEXT CHECK( status != '' ),
    data_pedido DATETIME NOT NULL,
    dest_nome TEXT CHECK( status != '' ),
    dest_endereco TEXT CHECK( status != '' ),
    
    FOREIGN KEY(cliente_id) REFERENCES cliente(id)
);

create table ocorrencia (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    entrega_id INTEGER NOT NULL,
    descricao TEXT CHECK( descricao != '' ),
    data_registro DATETIME NOT NULL,
    
    FOREIGN KEY(entrega_id) REFERENCES entrega(id)
);
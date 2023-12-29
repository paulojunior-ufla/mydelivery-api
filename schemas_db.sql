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
    data_pedido TEXT CHECK( data_pedido != '' ),
    data_finalizacao TEXT,
    dest_nome TEXT CHECK( dest_nome != '' ),
    dest_logradouro TEXT CHECK( dest_logradouro != '' ),
    dest_numero TEXT CHECK( dest_numero != '' ),
    dest_complemento TEXT CHECK( dest_complemento != '' ),
    dest_bairro TEXT CHECK( dest_bairro != '' ),

    FOREIGN KEY(cliente_id) REFERENCES cliente(id)
);
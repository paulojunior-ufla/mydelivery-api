create table cliente (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    nome TEXT CHECK( nome != '' ),
    email TEXT CHECK( email != '' ),
    telefone TEXT CHECK( telefone != '' )
);
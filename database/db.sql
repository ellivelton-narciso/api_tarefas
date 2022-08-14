create table if not exists despesas
(
    id        int auto_increment
        primary key,
    descricao varchar(2000) default 'Sem descrição.' not null,
    valor     float                                  not null,
    data      date                                   not null
);

create table if not exists receitas
(
    id        int auto_increment
        primary key,
    descricao varchar(5000) default 'Sem descrição.' not null,
    valor     float                                  not null,
    data      date                                   not null
);



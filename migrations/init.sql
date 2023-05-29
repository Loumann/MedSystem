create database medbase

create table "User"
(
    "ID"         serial primary key,
    "FirstName"  varchar,
    "LastName"   varchar,
    "Patronymic" varchar,
    "SNILS"      varchar
);


create table Analise
(
    "ID"   integer primary key,
    "Date" varchar,
    "Bld"  varchar,
    "Ubg"  varchar,
    "Bil"  varchar,
    "Pro"  varchar,
    "Nit"  varchar,
    "Ket"  varchar,
    "Glu"  varchar,
    "PH"   varchar,
    "SG"   varchar,
    "Leu"  varchar
);
create table "UserAnalise"
(
    "AnaliseId" integer   not null references "Analise",
    "UserId"    integer   not null references "User",
        constraint "UserAnalisePK"
            primary key ("AnaliseId", "UserId")
);



 create table UserLogin
(
    Username varchar,
    Password varchar
);

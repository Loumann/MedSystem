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

const getAnaliseBox = (analise) => { const box = document.createElement("div"); box.className = "analise-box";

box.append(

getAnaliseRow("Билирубин(Bil)", analise.bil, (v) => v >=7 && v <=9 || v === undefined), ); return box; };

const getAnaliseRow = (name, value, valid) => { const row = document.createElement("div"); const nameElement = document.createElement("div"); const valueElement = document.createElement("div");

nameElement.className = "analise-name"; nameElement.textContent = ${name}:;

valueElement.className = "analise-value"; valueElement.textContent = value;

if (valid) { const isValid = valid(value); if (isValid === undefined) { row.classList.add("undefined"); } else if (isValid) { row.classList.add("valid"); } else { row.classList.add("invalid"); } } else { row.classList.add("valid"); }

row.append(nameElement, valueElement);

return row; };
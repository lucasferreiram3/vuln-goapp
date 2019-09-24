create database vulnapp;
create table vulnapp.user (id int not null auto_increment primary key, name varchar(255) not null,mail varchar(255),age int not null,passwd varchar(255) not null, created_at timestamp not null default current_timestamp, updated_at timestamp not null default current_timestamp on update current_timestamp);
insert into vulnapp.user (name,mail,age,passwd) values ("Amuro Ray","RX-78-2@EFSF.com",15,"Mieru,Mieruzo!"),("Char Aznable","MS-06-S@Zeon.com",20,"BouyaDakarasa..."),("Kamille Bidan","MSZ-006@AEUG.com",17,"Kikoeru...Koega..."),("Judau Ashta","MSZ-010@AEUG.com",14,"Hamaaaan!!"),("Banagher Links","RX-0@londo.bell",16,"HitoNoMiraiHa...HitoGaTsukuruMonoDa!!!");
create table vulnapp.sessions (uid int,sessionid varchar(128));


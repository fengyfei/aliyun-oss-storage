create database if not exists oss;
set database=core;

create table if not exists files(
	id 			int	  		serial primary_key,
	url			string(128) not null,
	hash 		string(128) not null,
	filename	string(64) 	not null,
	countall 	int			not null,
	daycount	int			not null,
	hourcount	int 		not null,
	created		date 		not null default current_timestamp(),
);

create user oss_acc with password 'oqv0qvk0J5jmv';
grant select, update, insert on table oss.files to oss_acc;
